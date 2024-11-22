{{- range $method := .Methods }}

func List{{ $method.Name }}(id string) ([]{{ $method.NameSingular }}, error) {
	return ListAll(id, DefaultSortBy(), List{{ $method.Name }}PageSortBy)
}

func List{{ $method.Name }}CreatedSince(id string, createdSince *time.Time) ([]{{ $method.NameSingular }}, error) {
	return ListAll(id, &SortBy{
		Name:      "date",
		Asc:       true,
		AfterDate: createdSince,
	}, List{{ $method.Name }}PageSortBy)
}

func List{{ $method.Name }}UpdatedSince(accountId string, updatedSince *time.Time) ([]{{ $method.NameSingular }}, error) {
	return ListAll(accountId, &SortBy{
		Name:         "updated",
		Asc:          true,
		AfterDate:    updatedSince,
		UpdatedAfter: true,
	}, List{{ $method.Name }}PageSortBy)
}
{{end}}
