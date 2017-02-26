package utils

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(body []byte) ([]byte, error) {
	pretty := bytes.Buffer{}
	err := json.Indent(&pretty, body, "", "  ")
	if err != nil {
		return nil, err
	}

	return pretty.Bytes(), nil
}
