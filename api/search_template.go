// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// SearchTemplate - see https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithAllowNoIndices, WithExpandWildcards, WithExplain, WithIgnoreUnavailable, WithPreference, WithProfile, WithRouting, WithScroll, WithSearchType, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) SearchTemplate(options ...*Option) (*SearchTemplateResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["SearchTemplate"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &SearchTemplateResponse{resp}, err
}

// SearchTemplateResponse is the response for SearchTemplate
type SearchTemplateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
