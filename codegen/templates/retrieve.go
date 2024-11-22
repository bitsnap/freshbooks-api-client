package templates

import (
	"embed"
	"strings"

	. "github.com/bitsnap/bitsnap/packages/go-util"
)

//go:embed retrieve.go.tpl
var retrieveTemplateFS embed.FS

var retrieveTemplate = MustParseTemplate(retrieveTemplateFS, "retrieve.go.tpl")

func GenerateRetrieve(nameSingular string, url string) string {
	return GenerateRetrieveForWithEnding(nameSingular, url, "accounting/account/", "", "accountId")
}

func GenerateRetrieveFor(nameSingular string, url string, urlSuffix string, idName string) string {
	return GenerateRetrieveForWithEnding(nameSingular, url, urlSuffix, "", idName)
}

func GenerateRetrieveForWithEnding(nameSingular string, url string, urlSuffix string, urlEnding string, idName string) string {
	type RetrieveEndpoint struct {
		NameSingular string
		Url          string
		MaxRedirects uint
		Timeout      uint

		URLSuffix string
		URLEnding string
		IDName    string
	}

	output := strings.Builder{}
	err := retrieveTemplate.Execute(&output, RetrieveEndpoint{
		NameSingular: nameSingular,
		Url:          url,
		MaxRedirects: 5,
		Timeout:      10,

		URLSuffix: urlSuffix,
		URLEnding: urlEnding,
		IDName:    idName,
	})
	if err != nil {
		panic(err)
	}

	return output.String()
}
