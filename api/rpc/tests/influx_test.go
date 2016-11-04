package tests

import (
	"testing"
	"time"

	"github.com/appcelerator/amp/api/server"
	. "github.com/appcelerator/amp/data/influx"
)

var (
	influx Influx
)

func TestInfluxQuery(t *testing.T) {
	res, err := influx.Query("SHOW MEASUREMENTS")
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Results) == 0 {
		t.Errorf("Expected results")
	}
	if len(res.Results[0].Series) == 0 {
		t.Errorf("Expected series")
	}
	if res.Results[0].Series[0].Name != "measurements" {
		t.Errorf("Expected name to be %s, actual=%s \n", "measurement", res.Results[0].Series[0].Name)
	}
}

func influxInit(config server.Config) {
	cstr := config.InfluxURL
	influx = New(cstr, "_internal", "admin", "changme")
	influx.Connect(60 * time.Second)
}

func influxEnd() {
	influx.Close()
}
