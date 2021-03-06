package cli

import (
	"fmt"

	"io/ioutil"
	"path"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

// read lookup directory by parsing its contents
func parseLookup(directory string) (regexMap map[string]string, err error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		regexMap, err = generateRegexes(path.Join(directory, file.Name()))
	}
	return
}

// parse lookup directory and unmarshal its contents
func generateRegexes(fileName string) (regexMap map[string]string, err error) {
	if filepath.Ext(fileName) != ".yml" {
		return nil, fmt.Errorf("Cannot parse non-yaml file: %s", fileName)
	}
	pairs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to read yaml regex lookup: %s. Error: %v", fileName, err)
	}
	if err := yaml.Unmarshal(pairs, &regexMap); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal yaml lookup: %s. Error: %v", fileName, err)
	}
	return
}

// read specs from directory by parsing its contents
func parseSpec(directory string, timeout time.Duration) ([]*TestSpec, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	tests := []*TestSpec{}
	for _, file := range files {
		test, err := generateTestSpecs(path.Join(directory, file.Name()), timeout)
		if err != nil {
			return nil, err
		}
		if test != nil {
			tests = append(tests, test)
		}
	}
	return tests, nil
}

// parse samples directory and unmarshal its contents
func generateTestSpecs(fileName string, timeout time.Duration) (*TestSpec, error) {
	if filepath.Ext(fileName) != ".yml" {
		return nil, fmt.Errorf("Cannot parse non-yaml file: %s", fileName)
	}
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to read yaml test spec: %s. Error: %v", fileName, err)
	}
	testSpec := &TestSpec{
		Name:    fileName,
		Timeout: timeout,
	}
	var commandMap []CommandSpec
	if err = yaml.Unmarshal(contents, &commandMap); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal yaml test spec: %s. Error: %v", fileName, err)
	}
	for _, command := range commandMap {
		if command.Timeout == "" {
			// default command spec timeout
			command.Timeout = "30s"
		}
		testSpec.Commands = append(testSpec.Commands, command)
	}
	return testSpec, nil
}
