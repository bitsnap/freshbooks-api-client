package scraper

import (
	"fmt"
	"slices"
	"strings"

	. "github.com/bitsnap/bitsnap/packages/go-util"
)

func modelName(url string, objName string) string {
	return formatModelName(url, objName, false)
}

func enumName(url string, objName string) string {
	return formatModelName(url, objName, true)
}

func formatModelName(url string, objName string, isEnum bool) string {
	comp := strings.Split(url, "/")
	modelName := SnakeToCamel(strings.Replace(comp[len(comp)-1], "-", "_", -1), true)

	if isEnum {
		return "enums." + modelName + objName
	}

	return "models." + modelName + objName
}

func formatTag(tag string, required bool) string {
	r := ""
	if required {
		r = ` validate:"required"`
	}

	return fmt.Sprintf("`json:\"%v\"%v`", tag, r)
}

func isEnum(name string) bool {
	return slices.Contains([]string{
		//"type",
	}, name)
}

func formatType(t, d, url, name string) string {
	if t == "int" || t == "integer" {
		if strings.Contains(d, "0 for ") && strings.Contains(d, "1 for ") {
			return "uint8"
		}

		return "int64"
	}

	if t == "datetime" || t == "DateTime" {
		return "*time.Time"
	}

	if t == "date" || t == "Date" {
		return "*time.Time"
	}

	if t == "boolean" {
		return "bool"
	}

	if t == "decimal" {
		return "*decimal.Decimal"
	}

	if t == "string array" {
		return "[]string"
	}

	if t == "array" {
		return "[]" + modelName(url, name)
	}

	if t == "object array" {
		return "[]" + modelName(url, name)
	}

	if t == "enum" {
		return enumName(url, name)
	}

	if isEnum(t) {
		return enumName(url, name)
	}

	if t == "object" {
		return modelName(url, name)
	}

	return t
}
