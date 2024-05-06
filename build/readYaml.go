package build

import (
	"fmt"
	"log"
	"os"

	"ezfaas/models"

	"gopkg.in/yaml.v1"
)

func ReadYaml(path string) (models.Config, error) {
	// obj := make(map[string]interface{})
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var conf models.Config
	fmt.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	var name string = conf.Provider.Name
	// fmt.Println(conf.provider.name)
	log.Println("Provider Name:" + name)
	log.Println("Provider Gateway:" + conf.Provider.Gateway)
	for _, function := range conf.Functions {
		fmt.Println(function.Name)
	}
	return conf, nil
}
