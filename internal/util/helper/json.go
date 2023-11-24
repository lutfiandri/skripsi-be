package helper

import (
	"encoding/json"
)

func UnmarshalJson[T any](data []byte) (T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}

	return result, nil
}
