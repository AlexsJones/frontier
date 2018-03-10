package config

import (
	"log"
	"testing"
)

func TestLoadExampleConfig(t *testing.T) {

	_, err := LoadResource("test.yaml", "example_config.yaml")
	if err != nil {
		t.Fail()
	}
}

func TestLoadNonExistingConfig(t *testing.T) {

	_, err := LoadResource("test.yaml", "null.yaml")
	if err != nil {
		log.Println(err.Error())
	} else {
		t.Fail()
	}
}
