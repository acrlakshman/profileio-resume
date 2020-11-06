package profileio

import (
	"bytes"
	"strings"
	"text/template"
)

// ParsePanther Publication
func (item *PublicationDetail) ParsePanther(config Config) string {
	fMap := template.FuncMap{
		"SanitizeText": func(text string) string {
			return SanitizeText(text)
		},
		"MakeTitle": func(text string) string {
			return makeTitle(text)
		},
		"SanitizeAndMakeTitle": func(text string) string {
			return SanitizeAndMakeTitle(text)
		},
		"IsValidURL": func(text string) bool {
			return isValidURL(text)
		},
		"ParseTitleWithURL": func(title string) string {
			return item.parseTitleWithURL(title, config)
		},
	}

	switch item.Value.Type {
	case publicationFieldNameMap[ArticleField]:
		t, _ := template.New(publicationFieldNameMap[ArticleField]).Funcs(fMap).Parse("{{ $title := SanitizeText .Title }}{{ ParseTitleWithURL $title }}, {{ SanitizeText .Author }}, {{ SanitizeText .Journal -}}, {{ .Volume }}{{if .Number}}({{ .Number }}){{end}}:{{ .Pages }}, {{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[BookField]:
		t, _ := template.New(publicationFieldNameMap[BookField]).Funcs(fMap).Parse("{{ $title := SanitizeText .Title }}{{ ParseTitleWithURL $title }}, {{ SanitizeText .Author }}{{if .Series}}, volume {{ .Volume }} of {{ SanitizeText .Series -}}{{end}}, {{ SanitizeText .Publisher }}, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{if .Edition}}{{ SanitizeText .Edition }} edition, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[ThesisField]:
		t, _ := template.New(publicationFieldNameMap[ThesisField]).Funcs(fMap).Parse("{{ $title := SanitizeText .Title }}{{ ParseTitleWithURL $title }}, {{ SanitizeText .Author }}, {{ MakeTitle .Category }} thesis, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[TechReportField]:
		t, _ := template.New(publicationFieldNameMap[TechReportField]).Funcs(fMap).Parse("{{ $title := SanitizeText .Title }}{{ ParseTitleWithURL $title }}, {{ SanitizeText .Author }}, {{ SanitizeText .Series }}, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[InCollectionField]:
		t, _ := template.New(publicationFieldNameMap[InCollectionField]).Funcs(fMap).Parse("{{ $title := SanitizeText .Title }}{{ ParseTitleWithURL $title }}, {{ SanitizeText .Author }}{{if .Editor}}{{ SanitizeText .Editor }}, editor, in{{end}} {{ SanitizeText .Booktitle -}}{{if .Pages}}, pages {{ SanitizeText .Pages }}{{end}}. {{if .Publisher}}{{ SanitizeText .Publisher }}, {{end}}{{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[MiscField]:
		t, _ := template.New(publicationFieldNameMap[MiscField]).Funcs(fMap).Parse("{{if .Title}}{{ $title := SanitizeText .Title }}{{ ParseTitleWithURL $title }}, {{end}}{{ SanitizeText .Author }}. {{if IsValidURL .Howpublished}}\\url{ {{- .Howpublished -}} }{{else}}{{ .Howpublished }}{{end}}{{if .Month}}, {{ .Month }}{{end}}{{if .Year}} {{ .Year }}{{end}}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[UnPublishedField]:
		t, _ := template.New(publicationFieldNameMap[UnPublishedField]).Funcs(fMap).Parse("{{if .Title}}{{ $title := SanitizeText .Title }}{{ ParseTitleWithURL $title }}, {{end}}{{ SanitizeText .Author }}. {{if .Note}}{{ .Note }}{{end}}{{if .Year}}, {{ .Year }}{{end}}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	default:
		return ""
	}
}

// parseTitleWithURL parse
func (item *PublicationDetail) parseTitleWithURL(title string, config Config) string {
	if title == "" {
		title = item.Value.Title
	}
	if title == "" {
		title = item.Value.Booktitle
	}

	if item.Value.URL != "" && item.Value.WebPage.Slug == "" {
		return "\\href{" + item.Value.URL + "}{" + title + "}"
	} else if item.Value.WebPage.Slug != "" {
		return "\\href{" + strings.TrimRight(config.Homepage, "/") + "/" + item.Value.WebPage.Slug + "}{" + title + "}"
	} else {
		return title
	}
}
