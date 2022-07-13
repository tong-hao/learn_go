package main

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


type Config struct {
	Tasks  map[string]Task //`yaml:"tasks"`
}

type Task struct {
	Ssh string
	Input map[string]string //`yaml:"input"`
}

func main() {
	// resultMap := make(map[string]interface{})
	//inputMap:= make(map[string]interface{})
	yamlFile, err := ioutil.ReadFile("yaml2.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err. %v ", err)
	}

	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println(c)
	// log.Println("conf", resultMap)

}
