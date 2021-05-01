package main

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// https://cloud.google.com/translate/docs/samples/translate-text-with-model#translate_text_with_model-go

// TODO : alternatively: https://github.com/googleapis/google-api-go-client/blob/master/translate/v2/translate-gen.go

//------------------------------------------------------------------------------

func translateQuery(queryText string, sourceLang string, targetLang string) string {

	lang, err := language.Parse(targetLang)
	if err != nil {
		log.Fatal(err)
	}

	var model string // TODO

	ctx := context.Background()

	client, err := translate.NewClient(ctx)
	//client, err := translate.NewService(ctx, option.WithCredentialsFile("path/to/keyfile.json"))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Translate(ctx, []string{queryText}, lang, &translate.Options{
		Model: model, // Either "nmt" or "base".
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(resp) == 0 {
		return ""
	}
	return resp[0].Text
}

//------------------------------------------------------------------------------
