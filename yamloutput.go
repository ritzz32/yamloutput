package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"regexp"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
	Meta
}

type Meta struct {
	Color string `yaml:"color"`
	Size  string `yaml:"size"`
}

func main() {
	p := Person{"John", 30, Meta{"Yellow", "Large"}}
	y, err := yaml.Marshal(p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(string(y))

	var node1 string = "prod-boi-service-group1-app1"
	var node2 string = "puppetagent-adskafjlk-adsf"
	var node3 string = "woot1"

	fmt.Println(evalName(node1))
	fmt.Println(evalName(node2))
	fmt.Println(evalName(node3))
}

func evalName(nodeName string) (string, []string) {
	reLinuxServer := regexp.MustCompile(`^([a-z0-9]+)-([a-z0-9]+)-([a-z0-9]+)-([a-z0-9]+)-([a-z]+)([0-9]+)`)
	reKubernetes := regexp.MustCompile(`^([a-z0-9]+)-.*`)

	switch {
	case reLinuxServer.MatchString(nodeName):
		return "LinuxServer", reLinuxServer.FindStringSubmatch(nodeName)
	case reKubernetes.MatchString(nodeName):
		return "Kubernetes", reKubernetes.FindStringSubmatch(nodeName)
	default:
		return "No Match", make([]string, 0)
	}
}
