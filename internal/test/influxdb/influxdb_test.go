package influxdb_test

import (
	"fmt"
	"testing"

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

		err := influxdb.WriteData(writeAPI, "project_devices", tags, fields)
		if err != nil {
			t.Errorf("%s\n", err.Error())
		}
	}

}

func TestQuery(t *testing.T) {
	queryAPI := influxdb.GetQueryAPI()

	r, err := influxdb.Query(queryAPI, `from(bucket: "dev")
	|> range(start: v.timeRangeStart, stop: v.timeRangeStop)
	|> filter(fn: (r) => r["_measurement"] == "test_measurement")
	|> aggregateWindow(every: v.windowPeriod, fn: mean, createEmpty: false)
	|> yield(name: "mean")
  `)
	if err != nil {
		log.Err(err).Msg("")
		t.Error()
		return
	}
	for r.Next() {
		log.Info().Msg((fmt.Sprintln(r.Record())))
	}
	if err := r.Err(); err != nil {
		log.Err(err)
	}
}
