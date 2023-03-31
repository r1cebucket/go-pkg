package utils

import (
	"database/sql"

	"github.com/r1cebucket/gopkg/log"
)

func RowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		log.Err(err).Msg("Faile to read field name")
		return nil, err
	}
	// write values to map
	data := []map[string]interface{}{}
	row := make([]interface{}, len(columns))
	for i := range row {
		// convert to *interface
		var tmp interface{}
		row[i] = &tmp
	}
	for rows.Next() {
		err = rows.Scan(row...)
		if err != nil {
			log.Err(err).Msg("Faile to read row")
			return nil, err
		}

		item := make(map[string]interface{})
		for i, val := range row {
			// val is interface{} which store *interface{} value
			// convert val to *interface{} and then get the value
			item[columns[i]] = *(val.(*interface{}))
		}
		data = append(data, item)
	}
	return data, nil
}
