package linter

import (
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	_ "github.com/martinisecurity/shaken-pki-reports/lints"
)

func LintUrl(url *url.URL) *lint.LintResultSet {
	result := &lint.LintResultSet{
		Timestamp: time.Now(),
		Time:      0,
		Url:       url,
		Results:   map[string]*lint.LintResult{},
	}
	urlString := url.String()

	// send GET request
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{}
	reqStart := time.Now()
	response, err := http.Get(urlString)
	reqEnd := time.Now()
	if err != nil {
		// Disable SSL verification
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		var err2 error
		reqStart = time.Now()
		response, err2 = http.Get(urlString)
		reqEnd = time.Now()
		if err2 == nil {
			result.Append("e_tls_transport", &lint.LintResult{
				Status:  lint.Error,
				Details: err.Error(),
			})
		} else {
			err = err2
		}
	}
	result.Time = int(reqEnd.Sub(reqStart).Milliseconds())
	testData := &lint.LintData{
		Url:      url,
		Response: response,
		Error:    err,
	}

	if response != nil {
		// test redirect
		if response.StatusCode == 200 {
			req, _ := http.NewRequest("GET", urlString, nil)
			if req != nil {
				client := new(http.Client)
				client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
					result.Append("e_atis_redirect", &lint.LintResult{
						Status:  lint.Error,
						Details: "The STI-VS shall not follow HTTP redirections",
					})

					return errors.New("the STI-VS shall not follow HTTP redirections")
				}
				client.Do(req)
			}
		}

		// read response body
		result.StatusCode = response.StatusCode
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			testData.Error = err
		} else {
			testData.Body = body
			result.Body = body
		}
	} else {
		result.Append("e_bad_url", &lint.LintResult{
			Status:  lint.Error,
			Details: err.Error(),
		})

		return result
	}

	registry := lint.GetGlobalRegistry()
	for code, v := range registry.Rules {
		if v.Source == lint.SystemSource {
			// skip System tests
			continue
		}

		rule := v.Rule()
		if rule.CheckApplies(testData) {
			result.Append(code, v.Rule().Execute(testData))
		}
	}

	return result
}
