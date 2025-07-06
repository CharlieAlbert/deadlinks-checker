package checker

import (
	"net/http"
	"time"
)

type LinkStatus struct {
	URL string
	Alive bool
	Status string
}

func CheckLink(url string) LinkStatus {
	client := &http.Client {
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
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