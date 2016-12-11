package main

import (
	"bytes"
	"github.com/appcelerator/amp/api/rpc/function"
	"github.com/appcelerator/amp/config"
	"github.com/appcelerator/amp/pkg/nats-streaming"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stringid"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

// ## `amp-function-worker`
// This service role is to:
// - listen to function calls on the "function call" NATS topic
// - create the corresponding function container
// - attach to the created container (for stream management)
// - pass the function call input parameter through standard input
// - start the container
// - read the standard output of the container and wait for it to close
// - post the response back to NATS on the "returnTo" topic specified in the call

// build vars
var (
	// Version is set with a linker flag (see Makefile)
	Version string

	// Build is set with a linker flag (see Makefile)
	Build string

	// NATS Streaming
	natsStreaming ns.NatsStreaming

	// Docker is the Docker client
	docker *client.Client
)

func main() {
	log.Printf("%s (version: %s, build: %s)\n", os.Args[0], Version, Build)

	// NATS Connect
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Unable to get hostname: %s", err)
	}
	if natsStreaming.Connect(amp.NatsDefaultURL, amp.NatsClusterID, os.Args[0]+"-"+hostname, amp.DefaultTimeout) != nil {
		log.Fatal(err)
	}

	// NATS, subscribe to function topic
	_, err = natsStreaming.GetClient().Subscribe(amp.NatsFunctionTopic, messageHandler, stan.DeliverAllAvailable())
	if err != nil {
		natsStreaming.Close()
		log.Fatalf("Unable to subscribe to %s topic: %s", amp.NatsFunctionTopic, err)
	}
	log.Println("Subscribed to topic:", amp.NatsFunctionTopic)

	// Docker
	log.Printf("connecting to Docker API at %s version API: %s\n", amp.DockerDefaultURL, amp.DockerDefaultVersion)
	docker, err = client.NewClient(amp.DockerDefaultURL, amp.DockerDefaultVersion, nil, nil)
	if err != nil {
		log.Fatalln("Unable to connect to docker", err)
	}
	log.Printf("Connected to Docker at %s\n", amp.DockerDefaultURL)

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			log.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			natsStreaming.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func messageHandler(msg *stan.Msg) {
	go processMessage(msg)
}

func processMessage(msg *stan.Msg) {
	// Parse function call message
	functionCall := function.FunctionCall{}
	err := proto.Unmarshal(msg.Data, &functionCall)
	if err != nil {
		log.Println("error unmarshalling function call:", err)
		return
	}
	log.Println("Received a function call", functionCall.String())

	ctx := context.Background()

	// Create container
	container, err := containerCreate(ctx, "appcelerator/amp-demo-function")
	if err != nil {
		log.Println("error creating container:", err)
		return
	}
	log.Println("Created container:", container.ID)

	// Attach container streams
	attachment, err := containerAttach(ctx, container.ID)
	if err != nil {
		log.Println("error attaching container:", err)
		return
	}
	log.Println("Attached to container:", container.ID)

	// Write function call input to container standard input
	inputStream := bytes.NewReader(functionCall.Input)
	written, err := io.Copy(attachment.Conn, inputStream)
	if err != nil {
		log.Println("Unable to write to input stream:", err)
		return
	}
	log.Println("bytes written:", written)

	// Start
	if err = containerStart(ctx, container.ID); err != nil {
		log.Println("error starting container:", err)
		return
	}
	log.Println("Function call executed")

	// Close the attachment
	if err := attachment.CloseWrite(); err != nil {
		log.Println("Unable to close attachment:", err)
		return
	}
	log.Println("attachment successfully closed")

	// Eat 8 bytes of header on the standard output of the process
	for i := 0; i < 8; i++ {
		if _, err := attachment.Reader.ReadByte(); err != nil {
			log.Println("Unable to read from output stream:", err)
			return
		}
	}

	// Read standard output until it ends
	output, err := ioutil.ReadAll(attachment.Reader)
	if err != nil {
		log.Println("Unable to read from output stream:", err)
		return
	}
	read := len(output)
	log.Println("bytes read:", read)
	attachment.Close()

	// Post response to NATS
	functionReturn := function.FunctionReturn{
		CallID:     functionCall.CallID,
		FunctionID: functionCall.FunctionID,
		Output:     output,
	}

	// Encode the proto object
	encoded, err := proto.Marshal(&functionReturn)
	if err != nil {
		log.Println("error marshalling function return:", err)
		return
	}

	// Publish the return to NATS
	_, err = natsStreaming.GetClient().PublishAsync(functionCall.ReturnTo, encoded, nil)
	if err != nil {
		log.Println("error publishing function return:", err)
		return
	}
	log.Println("function return successfuly submitted:", functionReturn.String())
}

func containerCreate(ctx context.Context, image string) (container.ContainerCreateCreatedBody, error) {
	containerConfig := &container.Config{
		Image:        image,
		OpenStdin:    true,
		AttachStdin:  true,
		AttachStdout: true,
		StdinOnce:    true,
	}
	hostConfig := &container.HostConfig{}
	networkingConfig := &network.NetworkingConfig{}
	return docker.ContainerCreate(ctx, containerConfig, hostConfig, networkingConfig, stringid.GenerateNonCryptoID())
}

func containerAttach(ctx context.Context, containerID string) (types.HijackedResponse, error) {
	attachOptions := types.ContainerAttachOptions{
		Stdin:  true,
		Stdout: true,
		Stream: true,
	}
	return docker.ContainerAttach(ctx, containerID, attachOptions)
}

func containerStart(ctx context.Context, containerID string) error {
	startOptions := types.ContainerStartOptions{}
	return docker.ContainerStart(ctx, containerID, startOptions)
}
