package templates

import (
	"embed"
	"strings"

	. "github.com/bitsnap/bitsnap/packages/go-util"
)

//go:embed list.go.tpl
var listTemplateFS embed.FS

var listTemplate = MustParseTemplate(listTemplateFS, "list.go.tpl")

func GenerateList(name string, nameSingular string, url string) string {
	return GenerateListFor(name, nameSingular, url, "accounting/account/", "accountId")
}

func GenerateListFor(name string, nameSingular string, url string, urlSuffix string, idName string) string {
	type ListEndpoint struct {
		Name         string
		NameSingular string
		Url          string
		MaxRedirects uint
		Timeout      uint

		URLSuffix string
		IDName    string
	}

	output := strings.Builder{}
	err := listTemplate.Execute(&output, ListEndpoint{
		Name:         name,
		NameSingular: nameSingular,
		Url:          url,
		MaxRedirects: 5,
		Timeout:      10,

		URLSuffix: urlSuffix,
		IDName:    idName,
	})
	if err != nil {
		panic(err)
	}

	return output.String()
}
