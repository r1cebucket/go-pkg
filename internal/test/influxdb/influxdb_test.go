package influxdb_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/r1cebucket/gopkg/config"
	"github.com/r1cebucket/gopkg/influxdb"
	"github.com/r1cebucket/gopkg/log"
)

func init() {
	log.Setup("debug")
	config.Parse("../../configs/conf.json")
}

func TestWrite(t *testing.T) {
	writeAPI := influxdb.GetWriteAPI()
	// for value := 0; value < 5; value++ {
	// }

	for i := 0; i < 3; i++ {
		tags := map[string]string{
			"project_name": "test_project_1",
		}
		fields := map[string]interface{}{
			"device_num": i,
		}

		err := influxdb.Write(writeAPI, "project_devices", tags, fields)
		if err != nil {
			t.Errorf("%s\n", err.Error())
		}
		time.Sleep(time.Second)
	}

}

func TestQuery(t *testing.T) {
	queryAPI := influxdb.GetQueryAPI()

	r, err := influxdb.Query(queryAPI, `from(bucket: "dev")
	|> range(start: -1d)
	|> filter(fn: (r) => r["_measurement"] == "project_devices")
	|> yield()`)

	if err != nil {
		t.Error()
	}

	log.Info().Msg(fmt.Sprintln(r))
}
