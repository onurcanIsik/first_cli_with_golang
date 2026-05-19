# healthcheck

A fast URL health checker CLI written in Go.

## Usage

Check URLs from a JSON file:
healthcheck urls.json

Check URLs directly:
healthcheck https://google.com https://twitter.com

## urls.json format

```json
{
  "urls": [
    "https://google.com",
    "https://twitter.com"
  ]
}
```

## Features

- Parallel URL checking with goroutines
- Response time measurement
- Terminal UI with Bubbletea
- Press `q` to quit

## Built with

- Go
- [Bubbletea](https://github.com/charmbracelet/bubbletea)
