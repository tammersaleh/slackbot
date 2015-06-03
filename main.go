package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/nobueno", func() string {
		return "https://dl.dropbox.com/s/cr9rl6r6pqsbe1n/maxresdefault.jpg_1280720_2015-06-03_1_pm-48-12.png?dl=0"
	})
	m.Run()
}
