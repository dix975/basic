package utils

import (
	"encoding/json"
)

func ToJsonRepresatation(o interface{}) string {

	out, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	return string(out)

}

