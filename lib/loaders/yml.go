package loaders

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Page struct {
	Title    string
	Url      string
	Slug     string
	Sections []struct {
		Type    string
		Title   string
		Bgimg   string
		Content string
		Intro   string
	}
}

func Yaml(f string) (Page, error) {

	p := Page{}

	yamlFile, err := os.ReadFile("./yml/pages/" + f + ".yml")

	if err != nil {
		log.Printf("Load err #%v ", err)
	} else {
		err = yaml.Unmarshal(yamlFile, &p)

		if err != nil {
			log.Printf("Unmarshal: %v", err)
		}
	}

	return p, err
}
