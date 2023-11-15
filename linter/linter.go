package linter

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
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

	// create http client
	client := &http.Client{
		Timeout: time.Second * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{},
		},
	}

	// send request
	reqStart := time.Now()
	response, err := client.Get(urlString)
	reqEnd := time.Now()

	if err != nil {
		// check timeout
		if timeoutErr, ok := err.(net.Error); ok && timeoutErr.Timeout() {
			result.Append("e_request_timeout", &lint.LintResult{
				Status:  lint.Error,
				Details: "Request timed out (3s)",
			})
		} else {
			// try again with insecure TLS
			client.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
			reqStart = time.Now()
			var err2 error
			response, err2 = client.Get(urlString)
			reqEnd = time.Now()
			if err2 != nil {
				err = err2
			} else {
				result.Append("e_tls_transport", &lint.LintResult{
					Status:  lint.Error,
					Details: err.Error(),
				})
			}
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
	} else if timeoutErr, ok := err.(net.Error); ok && timeoutErr.Timeout() {
		return result
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
