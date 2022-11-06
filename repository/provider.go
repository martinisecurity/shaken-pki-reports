package repository

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/url"

	"golang.org/x/net/html/charset"
)

// Provider interface
type Provider interface {
	// GetURLs returns list of SHAKEN certificate repositories
	GetURLs() ([]*url.URL, error)
}

// Providers is list of Provider
type Providers []Provider

// ContentResult struct
type ContentResult struct {
	Key string
}

// ListBucketResult struct
type ListBucketResult struct {
	Contents []ContentResult
}

// ListBucketResultFromXML creates ListBucketResult from XML
func ListBucketResultFromXML(reader io.Reader) (*ListBucketResult, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	var res ListBucketResult
	err := decoder.Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ListBucketResult from XML")
	}

	return &res, nil
}
