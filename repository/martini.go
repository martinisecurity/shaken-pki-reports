package repository

import (
	"fmt"
	"net/http"
	"net/url"
)

// MartiniSecurityProvider struct
type MartiniSecurityProvider struct {
}

// GetURLs implements Provider interface
func (t *MartiniSecurityProvider) GetURLs() ([]*url.URL, error) {
	mtSecUrl := "https://p.mtsec.me"
	resp, err := http.Get(mtSecUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s, %w", mtSecUrl, err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get %s, incorrect Status code %d", mtSecUrl, resp.StatusCode)
	}
	defer resp.Body.Close()
	list, err := ListBucketResultFromXML(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read content of %s, %w", mtSecUrl, err)
	}

	var repos []*url.URL
	for _, v := range list.Contents {
		pathStr := fmt.Sprintf("%s/%s", mtSecUrl, v.Key)
		keyUrl, err := url.Parse(pathStr)
		if err != nil {
			continue
		}

		repos = append(repos, keyUrl)
	}

	return repos, nil
}
