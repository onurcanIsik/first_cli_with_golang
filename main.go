package main

import (
	"encoding/json"
	"health_checker/first_cli/checker"
	"os"
	"strings"

	"health_checker/first_cli/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type Config struct {
	Urls []string `json:"urls"`
}

func main() {

	var urls []string

	isFile := strings.HasSuffix(os.Args[1], ".json")
	if isFile {
		var config Config
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		data := make([]byte, 1024)
		count, err := file.Read(data)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data[:count], &config)
		if err != nil {
			panic(err)
		}
		urls = config.Urls
	} else {
		urls = os.Args[1:]
	}

	results := checker.CheckUrls(urls)
	p := tea.NewProgram(ui.NewModel(urls, results))
	if _, err := p.Run(); err != nil {
		panic(err)
	}

}
