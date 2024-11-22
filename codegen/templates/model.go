package templates

import (
	"embed"
	"strings"

	"github.com/bitsnap/bitsnap/packages/data-pipeline-general-accounting/cmd/bitsnap-accounting-freshbooks-codegen/codegen/scraper"
	. "github.com/bitsnap/bitsnap/packages/go-util"
)

//go:embed model.go.tpl
var modelTemplateFS embed.FS

var modelTemplate = MustParseTemplate(modelTemplateFS, "model.go.tpl")

type Model struct {
	Name       string
	Attributes []*scraper.Attribute
}

func GenerateModel(name, page string) string {
	attributes, err := scraper.ScrapeAttributes(page)
	if err != nil {
		Logger().Fatalln(err)
	}

	output := strings.Builder{}
	err = modelTemplate.Execute(&output, Model{name, attributes})
	if err != nil {
		panic(err)
	}

	return output.String()
}
