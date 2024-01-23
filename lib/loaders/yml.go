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
		Type          string
		Title         string
		Bgimg         string
		Bgcolor       string
		Content       string
		Intro         string
		Image         string
		Imagealt      string
		Imageposition string
		Textcolor     string
		Features      []struct {
			Title    string
			Image    string
			Imagealt string
			Intro    string
			Link     string
		}
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
	files, err := os.ReadDir("./yml/blog")
	if err != nil {
		fmt.Printf("%v", err)
	}
	for i, file := range files {
		if i == n {
			// limit reached break
			break
		}
		filename := fmt.Sprint(file)
		filename = filename[2 : len(filename)-4]
		fmt.Printf("Loading %s\n", filename)
		p, err := Yaml(filename, "blog")
		if err != nil {
			fmt.Println(err)
		}
		posts = append(posts, p)
	}

	return posts
}
