package helper

import (
	"bytes"
	"encoding/json"
	"errors"
)

func MarshalToString(value interface{}) string {
	byt, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	str := bytes.NewBuffer(byt).String()
	return str
}

func Marshal(value interface{}) []byte {
	byt, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	return byt
}

func UnmarshalFromString(str string, value interface{}) error {
	byt := []byte(str)
	if len(byt) == 0 {
		return errors.New("str is empty")
	}
	err := json.Unmarshal(byt, value)
	if err != nil {
		return err
	}
	return nil
}

func Unmarshal(data []byte, value interface{}) error {
	if len(data) == 0 {
		return errors.New("data is empty")
	}
	err := json.Unmarshal(data, value)
	if err != nil {
		return err
	}
	return nil
}
