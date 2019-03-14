package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
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
}
