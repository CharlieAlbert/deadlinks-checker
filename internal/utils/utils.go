package utils

import (
	"net/url"
	"strings"
)

func NormalizeLinks(rawLinks []string, baseURL string) ([]string, error) {
	seen := make(map[string]bool)
	var cleaned []string

	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	for _, href := range rawLinks {
		href = strings.TrimSpace(href)

		if href == "" || strings.HasPrefix(href, "#") ||
		strings.HasPrefix(href, "mailto:") ||
		strings.HasPrefix(href, "tel:") ||
		strings.HasPrefix(href, "javascript:") {
			continue
		}

		parsed, err := url.Parse(href)
		if err != nil {
			continue
		}

		resolved := base.ResolveReference(parsed).String()
		if !seen[resolved] {
			cleaned = append(cleaned, resolved)
			seen[resolved] = true
		}
	}

	return cleaned, nil
}