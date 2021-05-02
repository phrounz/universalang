package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//------------------------------------------------------------------------------

type TranslationAPIInput struct {
	Q      string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
	Format string `json:"format"`
}

type TranslationAPIOutput struct {
	Data TranslationAPIOutputData `json:"data"`
}

type TranslationAPIOutputData struct {
	Translations []TranslationAPIOutputTranslations `json:"translations"`
}

type TranslationAPIOutputTranslations struct {
	TranslatedText string `json:"translatedText"`
}

//------------------------------------------------------------------------------

func translateQuery2(token string, queryText string, sourceLang string, targetLang string) string {

	var jsonData, err = json.Marshal(TranslationAPIInput{
		Q:      queryText,
		Source: sourceLang,
		Target: targetLang,
		Format: "text",
	})
	if err != nil {
		log.Fatal(err)
	}

	var outputStatus int
	var outputBody string
	outputStatus, outputBody, err = runRequest(requestInput{
		URL:      "https://translation.googleapis.com/language/translate/v2",
		Method:   "POST",
		SendData: true,
		Data:     jsonData,
		FnSigningCallback: func(req *http.Request, body *strings.Reader) error {
			req.Header.Set("Authorization", "Bearer \""+token+"\"")
			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d %s\n", outputStatus, outputBody)

	var dec = json.NewDecoder(strings.NewReader(outputBody))
	var m TranslationAPIOutput
	if err := dec.Decode(&m); err != nil && err != io.EOF {
		log.Fatal(err)
	}

	return m.Data.Translations[0].TranslatedText
}

//------------------------------------------------------------------------------
