package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONMap is a map that can be used to store JSONB data in a database.
type JSONMap struct {
	M map[string]interface{}
}

func (jm JSONMap) Value() (driver.Value, error) {
	jsonBytes, err := json.Marshal(jm.M)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func (jm JSONMap) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var jsonBytes []byte
	switch src := src.(type) {
	case string:
		jsonBytes = []byte(src)
	case []byte:
		jsonBytes = src
	default:
		return fmt.Errorf("unsupported src type: %T for JSONMap", src)
	}

	return json.Unmarshal(jsonBytes, &jm.M)
}
