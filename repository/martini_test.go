package repository_test

import (
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/repository"
)

func TestMartiniSecurityProvider_GetURLs(t *testing.T) {
	prov := repository.MartiniSecurityProvider{}
	repos, err := prov.GetURLs()
	if err != nil {
		t.Errorf("MartiniSecurityProvider.GetURLs() error = %s", err.Error())
		return
	}
	if len(repos) == 0 {
		t.Errorf("MartiniSecurityProvider.GetURLs() len([]*url.URL) = 0")
		return
	}
}
