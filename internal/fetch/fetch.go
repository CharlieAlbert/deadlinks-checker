package fetch

import (
	"errors"
	"io"
	"net/http"
	"time"
)

func FetchHTML(url string) (string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Failed to fetch URL " + resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}