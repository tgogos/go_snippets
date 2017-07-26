# Go example - read yml
Golang, read configuration parameters from .yaml file

## test yaml

```yaml
app_config:
    database_name:     mane_vision_daem
    database_pass:     ubuntu
    database_port:     3306
    database_url:      "127.0.0.1"
    database_user:     root
    drop_when_no_rule: false
```

## Go code

```golang
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Settings struct {
	DBname         string `yaml:"database_name"`
	DBpass         string `yaml:"database_pass"`
	DBport         string `yaml:"database_port"`
	DBurl          string `yaml:"database_url"`
	DBuser         string `yaml:"database_user"`
	DropWhenNOrule bool   `yaml:"drop_when_no_rule"`
}

type Config struct {
	AppSettings Settings `yaml:"app_config"`
}

func main() {
	filename, _ := filepath.Abs("./app.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	//
	// print everything...
	//
	fmt.Printf("%#v\n\n", config.AppSettings)

	//
	// print one by one...
	//
	fmt.Printf("database_name: %s\n", config.AppSettings.DBname)
	fmt.Printf("database_pass: %s\n", config.AppSettings.DBpass)
	fmt.Printf("database_port: %s\n", config.AppSettings.DBport)
	fmt.Printf("database_url: %s\n", config.AppSettings.DBurl)
	fmt.Printf("database_user: %s\n", config.AppSettings.DBuser)
	fmt.Printf("drop_when_no_rule: %t\n", config.AppSettings.DropWhenNOrule)
}
```
