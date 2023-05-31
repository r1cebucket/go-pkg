package influxdb

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/r1cebucket/gopkg/config"
)

func GetWriteAPI() api.WriteAPIBlocking {
	token := config.Influxdb.Token
	url := config.Influxdb.URL
	client := influxdb2.NewClient(url, token)

	org := config.Influxdb.Org
	bucket := config.Influxdb.Bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)

	return writeAPI
}

func WriteData(writeAPI api.WriteAPIBlocking, measurement string, tags map[string]string,
	fields map[string]interface{}) error {
	point := write.NewPoint(measurement, tags, fields, time.Now())

	err := writeAPI.WritePoint(context.Background(), point)

	return err
}

func GetQueryAPI() api.QueryAPI {
	client := influxdb2.NewClient(config.Influxdb.URL, config.Influxdb.Token)
	queryAPI := client.QueryAPI(config.Influxdb.Org)

	return queryAPI
}

func Query(queryAPI api.QueryAPI, query string) ([]map[string]interface{}, error) {
	rows, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	results := []map[string]interface{}{}
	for rows.Next() {
		results = append(results, rows.Record().Values())
	}

	return results, nil
}
