package reporter

import (
	"fmt"
	"health_checker/first_cli/checker"
)

func ReportUrls(urls []string) {
	result := checker.CheckUrls(urls)
	fmt.Printf("URL Check Report: %s\n", result)
}
