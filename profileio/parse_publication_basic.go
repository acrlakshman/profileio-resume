package profileio

import (
	"bytes"
	"text/template"
)

// ParseBasic Publication
func (item *PublicationDetail) ParseBasic(config Config) string {
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
	}

	switch item.Value.Type {
	case publicationFieldNameMap[ArticleField]:
		t, _ := template.New(publicationFieldNameMap[ArticleField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{ SanitizeText .Title }}. {\\it {{ SanitizeText .Journal -}} }, {{ .Volume }}{{if .Number}}({{ .Number }}){{end}}:{{ .Pages }}, {{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[BookField]:
		t, _ := template.New(publicationFieldNameMap[BookField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {\\it {{ SanitizeAndMakeTitle .Title -}} }{{if .Series}}, volume {{ .Volume }} of {\\it {{ SanitizeText .Series -}} }{{end}}. {{ SanitizeText .Publisher }}, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{if .Edition}}{{ SanitizeText .Edition }} edition, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[ThesisField]:
		t, _ := template.New(publicationFieldNameMap[ThesisField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {\\it {{ SanitizeAndMakeTitle .Title -}} }. {{ MakeTitle .Category }} thesis, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[TechReportField]:
		t, _ := template.New(publicationFieldNameMap[TechReportField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{ SanitizeText .Title }}. {{ SanitizeText .Series }}, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[InCollectionField]:
		t, _ := template.New(publicationFieldNameMap[InCollectionField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{ SanitizeText .Title }}. In {{if .Editor}}{{ SanitizeText .Editor }}, editor, {{end}}{\\it {{ SanitizeText .Booktitle -}} }{{if .Pages}}, pages {{ SanitizeText .Pages }}{{end}}. {{if .Publisher}}{{ SanitizeText .Publisher }}, {{end}}{{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[MiscField]:
		t, _ := template.New(publicationFieldNameMap[MiscField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{if .Title}}{{ SanitizeText .Title }}. {{end}}{{if IsValidURL .Howpublished}}\\url{ {{- .Howpublished -}} }{{else}}{{ .Howpublished }}{{end}}{{if .Month}}, {{ .Month }}{{end}}{{if .Year}} {{ .Year }}{{end}}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case publicationFieldNameMap[UnPublishedField]:
		t, _ := template.New(publicationFieldNameMap[UnPublishedField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{if .Title}}{{ SanitizeText .Title }}. {{end}}{{if .Note}}{{ .Note }}{{end}}{{if .Year}}, {{ .Year }}{{end}}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	default:
		return ""
	}
}
