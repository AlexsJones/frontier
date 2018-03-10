package config

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	yaml "gopkg.in/yaml.v2"
)

//Server Configuration related structs
type Server struct {
	Port string `yaml:"Port"`
}

//Config is the struct for configuration management
type Config struct {
	APIVersion string `yaml:"APIVersion"`
	Server     Server `yaml:"Server"`
}

func (*Config) loadConfigurationFromFile(filePath string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	d := &Config{}
	err = yaml.Unmarshal(yamlFile, d)
	if err != nil {
		return nil, err
	}
	log.Println("Successfully loaded configuration file")
	return d, nil
}

//LoadResource either from URL or local path
func (c *Config) LoadResource(remoteConfigSavePath string, resourceURI string) (*Config, error) {
	u, err := url.ParseRequestURI(resourceURI)
	if err != nil {
		if _, err = os.Stat(resourceURI); os.IsNotExist(err) {
			return nil, err
		} else {
			//Load from file...
			return c.loadConfigurationFromFile(resourceURI)
		}
	}
	//Load from URL ...
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Delete local config if exists
	if _, err = os.Stat(remoteConfigSavePath); os.IsExist(err) {
		os.Remove(remoteConfigSavePath)
		log.Printf("Removed old version of local configuration: %s\n", remoteConfigSavePath)
	}

	log.Printf("Creating local file: %s\n", remoteConfigSavePath)
	out, err := os.Create(remoteConfigSavePath)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}

	return c.loadConfigurationFromFile(remoteConfigSavePath)
}
