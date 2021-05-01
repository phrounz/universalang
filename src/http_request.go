package main

import (
	"bytes"
	"net/http"
	"strings"
)

//------------------------------------------------------------------------------

type RequestInput struct {
	URLStr            string // (compulsory)
	Method            string // (compulsory) HTTP method, e.g. GET
	SendData          bool   // (optional, default false)
	ContentType       string // (needed only if SendData==true) Content-Type of body
	Data              []byte // (needed only if SendData==true) body
	FnSigningCallback func(req *http.Request, body *strings.Reader) error
}

//------------------------------------------------------------------------------

func runRequest(input RequestInput) (outputStatus int, outputBody string, err error) {

	// if true {
	// 	outputBody = `{
	// "data": {
	// "translations": [
	// {
	// "translatedText": "La Gran Pirámide de Giza (también conocida como la Pirámide de Khufu o la Pirámide de Keops) es la más antigua y más grande de las tres pirámides en el complejo de la pirámide de Giza."
	// }
	// ]
	// }
	// }`
	// 	return
	// }

	var client = &http.Client{Transport: &http.Transport{}}

	var req *http.Request
	if input.SendData {
		req, err = http.NewRequest(input.Method, input.URLStr, bytes.NewBuffer(input.Data))
		if err != nil {
			return
		}
		if input.ContentType != "" {
			req.Header.Set("Content-Type", input.ContentType)
		}
	} else {
		var body = bytes.NewReader([]byte{})
		req, err = http.NewRequest(input.Method, input.URLStr, body)
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
