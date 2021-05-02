package main

import (
	"bytes"
	"net/http"
	"strings"
)

//------------------------------------------------------------------------------

type requestInput struct {
	URL               string // (compulsory)
	Method            string // (compulsory) HTTP method, e.g. GET
	SendData          bool   // (optional, default false)
	ContentType       string // (needed only if SendData==true) Content-Type of body
	Data              []byte // (needed only if SendData==true) body
	FnSigningCallback func(req *http.Request, body *strings.Reader) error
}

//------------------------------------------------------------------------------

func runRequest(input requestInput) (outputStatus int, outputBody string, err error) {

	var client = &http.Client{Transport: &http.Transport{}}

	var req *http.Request
	if input.SendData {
		req, err = http.NewRequest(input.Method, input.URL, bytes.NewBuffer(input.Data))
		if err != nil {
			return
		}
		if input.ContentType != "" {
			req.Header.Set("Content-Type", input.ContentType)
		}
	} else {
		var body = bytes.NewReader([]byte{})
		req, err = http.NewRequest(input.Method, input.URL, body)
		if err != nil {
			return
		}
	}

	var response *http.Response
	if input.FnSigningCallback != nil {
		err = input.FnSigningCallback(req, strings.NewReader(string(input.Data)))
		if err != nil {
			return
		}
	}

	req.Close = true

	response, err = client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	outputStatus = response.StatusCode
	var buf = new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	outputBody = buf.String()

	return
}

//------------------------------------------------------------------------------
