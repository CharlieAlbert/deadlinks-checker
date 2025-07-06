package checker

import (
	"net/http"
	"sync"
	"time"
)

type LinkStatus struct {
	URL string
	Alive bool
	Status string
}

func CheckLink(url string) LinkStatus {
	client := &http.Client {
		Timeout: 5 * time.Second,
	}

	resp, err := client.Head(url)
	if err != nil || resp.StatusCode >= 400 {
		resp, err = client.Get(url)
		if err != nil {
			return LinkStatus{URL: url, Alive: false, Status: "Unreachable"}
		}
		defer resp.Body.Close()
	} else {
		defer resp.Body.Close()
	}

	alive := resp.StatusCode >= 200 && resp.StatusCode < 400
	return LinkStatus{
		URL: url,
		Alive: alive,
		Status: resp.Status,
	}
}

func CheckLinksConcurrently(urls []string) []LinkStatus {
	var wg sync.WaitGroup
	results := make(chan LinkStatus, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			result := CheckLink(link)
			results <- result
		}(url)
	}

	wg.Wait()
	close(results)

	var statuses []LinkStatus
	for res := range results {
		statuses = append(statuses, res)
	}

	return statuses
}