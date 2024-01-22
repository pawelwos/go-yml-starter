package loaders

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Page struct {
	Title     string
	Url       string
	Slug      string
	Excerpt   string
	Thumbnail string
	Date      string
	Sections  []struct {
		Type    string
		Title   string
		Bgimg   string
		Content string
		Intro   string
	}
}

func Yaml(f string, t string) (Page, error) {

	p := Page{}

	yamlFile, err := os.ReadFile("./yml/" + t + "/" + f + ".yml")

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

func Posts(n int) []Page {
	posts := make([]Page, 0)
	p, err := Yaml("first-post", "blog")
	if err != nil {
		fmt.Println(err)
	}
	posts = append(posts, p)
	return posts
}
