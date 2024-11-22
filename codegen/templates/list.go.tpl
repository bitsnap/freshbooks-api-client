func List{{.Name}}PageSortBy({{ .IDName }} string, page uint64, sort *SortBy) ([]{{.NameSingular}}, *Pagination, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	parsedUrl, err := url.Parse("https://api.freshbooks.com/{{ .URLSuffix }}" + {{ .IDName }} + "{{.Url}}")
	if err != nil {
		return nil, nil, err
	}
	
    content, err := getQueryPage(client, parsedUrl, page, sort)
    if err != nil {
        return nil, nil, err
    }
    
    type {{.NameSingular}}Result struct {
         {{.Name}} []{{.NameSingular}} `json:"{{CamelToSnake .Name}}"`
    }
    
    type {{.NameSingular}}ResponseContent struct {
        Result {{.NameSingular}}Result `json:"result"`
        Pagination
    }
    
    type {{.NameSingular}}Response struct {
        Response {{.NameSingular}}ResponseContent `json:"response"`
    }

	entries := {{.NameSingular}}Response{
		Response: {{.NameSingular}}ResponseContent{
		    Result: {{.NameSingular}}Result{
		        {{.Name}}: make([]{{.NameSingular}}, 0, 10),
		    },
		},
	}

	err = json.Unmarshal(content, &entries)
	if err != nil {
		return nil, nil, err
	}
	
	if len(entries.Response.Result.{{.Name}}) == 0 {
        return []{{.NameSingular}}{}, nil, nil
    }
	
	return entries.Response.Result.{{.Name}}, &entries.Response.Pagination, nil
}

func List{{.Name}}Page(site string, page uint64) ([]{{.NameSingular}}, *Pagination, error) {
	return List{{.Name}}PageSortBy(site, page, nil)
}
