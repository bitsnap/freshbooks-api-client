package scraper

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/bitsnap/bitsnap/packages/go-util"
	"github.com/gocolly/colly/v2"
)

type Attribute struct {
	Name        string
	Tag         string
	Type        string
	Description string
	Required    bool
}

func (a *Attribute) String() string {
	if a.Required {
		return fmt.Sprintf("%v required %v %v", a.Name, a.Tag, a.Type)
	}

	return fmt.Sprintf("%v %v %v", a.Name, a.Tag, a.Type)
}

func NewAttribute(name, t, description string, required bool) *Attribute {
	normalizedName := name
	if strings.HasSuffix(name, "id") && !strings.HasSuffix(name, "paid") && name != "id" && name != "uuid" {
		normalizedName = strings.TrimSuffix(normalizedName, "id")
		normalizedName = normalizedName + "_id"
	}

	if strings.HasPrefix(normalizedName, "s_") {
		normalizedName = strings.TrimPrefix(normalizedName, "s_")
		normalizedName = "shipping_" + normalizedName
	}

	if strings.HasPrefix(normalizedName, "p_") {
		normalizedName = strings.TrimPrefix(normalizedName, "p_")
		normalizedName = "billing_" + normalizedName
	}

	if strings.HasPrefix(normalizedName, "pref_") {
		normalizedName = strings.TrimPrefix(normalizedName, "pref_")
		normalizedName = "prefers_" + normalizedName
	}

	if strings.HasSuffix(normalizedName, "_gmail") {
		normalizedName = strings.TrimSuffix(normalizedName, "_gmail")
		normalizedName = normalizedName + "_ground_mail"
	}

	if normalizedName == "fname" {
		normalizedName = "first_name"
	}

	if normalizedName == "lname" {
		normalizedName = "last_name"
	}

	if !strings.HasPrefix(name, "-") && !strings.Contains(normalizedName, "_") {
		normalizedName = Capitalize(normalizedName)
	} else {
		normalizedName = SnakeToCamel(normalizedName, true)
	}

	if name == "id" {
		normalizedName = "Id"
		t = "int64" // NOTE: there are some typos, like https://www.freshbooks.com/api/services
	}

	return &Attribute{
		Name:        normalizedName,
		Tag:         name,
		Type:        t,
		Description: description,
		Required:    required,
	}
}

func formatName(s string) string {
	l := strings.Split(s, "\n")
	sb := strings.Builder{}
	for _, comp := range l {
		sb.WriteString(strings.TrimSpace(comp))
	}

	name := strings.TrimSpace(sb.String())
	name = strings.Replace(name, " ", "", -1)
	name = strings.Replace(name, "¶", "", -1)

	return name
}

func ScrapeAttributes(page string) ([]*Attribute, error) {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.freshbooks.com", "www.freshbooks.com"),
		colly.CacheDir("./.cache"),
	)

	attributes := make([]*Attribute, 0, 10)

	collector.OnHTML("div.col-content", func(content *colly.HTMLElement) {
		fieldsTable := content.DOM.Find("figure.wp-block-table").FilterFunction(func(_ int, table *goquery.Selection) bool {
			filterFunc := func(_ int, header *goquery.Selection) bool {
				return strings.ToLower(header.Text()) == "field"
			}

			exclusionFilterFunc := func(_ int, header *goquery.Selection) bool {
				return strings.Contains(strings.ToLower(header.Text()), "filter type")
			}

			inclusion := table.Find("thead tr th").FilterFunction(filterFunc).Length() > 0 || table.Find("tbody tr td").FilterFunction(filterFunc).Length() > 0
			exclusion := table.Find("thead tr th").FilterFunction(exclusionFilterFunc).Length() == 0 && table.Find("tbody tr td").FilterFunction(exclusionFilterFunc).Length() == 0

			return inclusion && exclusion
		}).First()

		if fieldsTable.Length() > 0 {
			fieldsTable.Find("tbody tr").Each(func(_ int, row *goquery.Selection) {
				name := ""
				t := ""
				description := ""
				row.Find("td").Each(func(i int, column *goquery.Selection) {
					if i == 0 {
						name = formatName(column.Text())
					}

					if i == 1 {
						t = strings.Replace(column.Text(), "¶", "", -1)
						t = strings.TrimSpace(t)
					}

					if i == 2 {
						description = strings.Replace(column.Text(), "¶", "", -1)
						description = strings.Replace(column.Text(), "\n", " ", -1)
						description = strings.TrimSpace(description)
					}
				})

				if name != "" && t != "" && (name != "Field" && t != "Type") {
					attribute := NewAttribute(name, t, description, false)
					attribute.Tag = formatTag(attribute.Tag, attribute.Required)
					attribute.Type = formatType(attribute.Type, attribute.Description, page, attribute.Name)

					attributes = append(attributes, attribute)
				}
			})

			fieldsTable.Find("tbody tr td u").Each(func(_ int, underLineNameColumn *goquery.Selection) {
				requiredFieldTag := formatName(underLineNameColumn.Text())

				for i := range attributes {
					if attributes[i].Tag == "`json:\""+requiredFieldTag+"\"`" {
						attributes[i].Tag = "`json:\"" + requiredFieldTag + "\" validate:\"required\"`"

						attributes[i].Required = true
					}
				}
			})

		} else {
			Logger().Infoln(page, "page field descriptions not found")
		}
	})

	collector.OnError(func(response *colly.Response, err error) {
		Logger().Errorw(response.Request.URL.String(), err)
	})

	err := collector.Visit(page)
	if err != nil {
		return nil, err
	}

	filtered := make([]*Attribute, 0, len(attributes))
	for _, attribute := range attributes {
		if !strings.HasPrefix(attribute.Tag, "`json:\"–") && !strings.HasPrefix(attribute.Tag, "`json:\"-") { // NOTE: BOTH HYPHENS and EM Dash !
			found := false
			for _, existing := range filtered {
				if existing.Name == attribute.Name {
					found = true

					if existing.Type != attribute.Type {
						Logger().Infoln("multiple types for", page, existing.Name, "field has both:", existing.Type, "and", attribute.Type)
					}

					break
				}
			}

			if !found {
				filtered = append(filtered, attribute)
			}
		}
	}

	return filtered, nil
}
