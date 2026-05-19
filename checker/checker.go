package checker

import (
	"fmt"
	"net/http"
	"time"
)

func CheckUrls(urls []string) []string {
	result := []string{}
	ch := make(chan string, len(urls))

	for _, url := range urls {
		go func(u string) {
			myTime := time.Now()
			response, err := http.Get(u)
			if err != nil {
				ch <- fmt.Sprintf("[DOWN] URL: %s, Error: %v\n", u, err)
			} else {
				ch <- fmt.Sprintf("[UP] URL: %s, Status Code: %d, Time: %s\n", u, response.StatusCode, myTime.Format(time.RFC3339))
			}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result = append(result, <-ch)
	}

	return result
}
