package json

import (
	"github.com/goccy/go-json"
)

func Json(data interface{}) string {
	jsonBytes, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonBytes)
}
