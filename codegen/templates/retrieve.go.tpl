func Retrieve{{.NameSingular}}ById({{ .IDName }}, id string) (*{{.NameSingular}}, error) {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	{{ if .URLEnding }}parsedUrl, err := url.Parse("https://api.freshbooks.com/{{ .URLSuffix }}" + {{ .IDName }} + "{{.Url}}/" + id + "{{ .URLEnding }}")
	{{ else }}parsedUrl, err := url.Parse("https://api.freshbooks.com/{{ .URLSuffix }}" + {{ .IDName }} + "{{.Url}}/" + id)
	{{ end }}	
	if err != nil {
		return nil, err
	}
	
    content, err := getQueryPage(client, parsedUrl, -1, nil)
    if err != nil {
        return nil, err
    }
    	
    type {{.NameSingular}}Result struct {
         {{.NameSingular}} {{.NameSingular}} `json:"{{CamelToSnake .NameSingular}}"`
    }
    
    type {{.NameSingular}}ResponseContent struct {
        Result {{.NameSingular}}Result `json:"result"`
    }
    
    type {{.NameSingular}}Response struct {
        Response {{.NameSingular}}ResponseContent `json:"response"`
    }

	entry := {{.NameSingular}}Response{}

	err = json.Unmarshal(content, &entry)
	if err != nil {
		return nil, err
	}

	return &entry.Response.Result.{{.NameSingular}}, nil
}
