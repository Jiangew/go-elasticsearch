// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// PutScript - the scripting module enables you to use scripts to evaluate custom expressions. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-scripting.html for more info.
//
// lang: script language.
//
// body: the document.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) PutScript(lang string, body map[string]interface{}, options ...*Option) (*PutScriptResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "PUT",
	}
	methodOptions := supportedOptions["PutScript"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &PutScriptResponse{resp}, err
}

// PutScriptResponse is the response for PutScript
type PutScriptResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
