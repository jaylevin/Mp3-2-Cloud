package config

import (
	"os"
	"encoding/json"
	"bytes"
)

func New() error {
	f, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer f.Close()

	uglyJson, err := json.Marshal(DefaultConfig())
	if err != nil {
		return err
	}

	var formattedJson bytes.Buffer
	err = json.Indent(&formattedJson, uglyJson, "", "    ")
	if err != nil {
		return err
	}

	_, err = f.Write(formattedJson.Bytes())
	if err != nil {
		return err
	}

	return nil
}
