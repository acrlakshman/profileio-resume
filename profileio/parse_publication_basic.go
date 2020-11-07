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
	case PublicationFieldNameMap[ArticleField]:
		t, _ := template.New(PublicationFieldNameMap[ArticleField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{ SanitizeText .Title }}. {\\it {{ SanitizeText .Journal -}} }, {{ .Volume }}{{if .Number}}({{ .Number }}){{end}}:{{ .Pages }}, {{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case PublicationFieldNameMap[BookField]:
		t, _ := template.New(PublicationFieldNameMap[BookField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {\\it {{ SanitizeAndMakeTitle .Title -}} }{{if .Series}}, volume {{ .Volume }} of {\\it {{ SanitizeText .Series -}} }{{end}}. {{ SanitizeText .Publisher }}, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{if .Edition}}{{ SanitizeText .Edition }} edition, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case PublicationFieldNameMap[ThesisField]:
		t, _ := template.New(PublicationFieldNameMap[ThesisField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {\\it {{ SanitizeAndMakeTitle .Title -}} }. {{ MakeTitle .Category }} thesis, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case PublicationFieldNameMap[TechReportField]:
		t, _ := template.New(PublicationFieldNameMap[TechReportField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{ SanitizeText .Title }}. {{ SanitizeText .Series }}, {{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case PublicationFieldNameMap[InCollectionField]:
		t, _ := template.New(PublicationFieldNameMap[InCollectionField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{ SanitizeText .Title }}. In {{if .Editor}}{{ SanitizeText .Editor }}, editor, {{end}}{\\it {{ SanitizeText .Booktitle -}} }{{if .Pages}}, pages {{ SanitizeText .Pages }}{{end}}. {{if .Publisher}}{{ SanitizeText .Publisher }}, {{end}}{{if .Address}}{{ SanitizeText .Address }}, {{end}}{{ .Year }}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case PublicationFieldNameMap[MiscField]:
		t, _ := template.New(PublicationFieldNameMap[MiscField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{if .Title}}{{ SanitizeText .Title }}. {{end}}{{if IsValidURL .Howpublished}}\\url{ {{- .Howpublished -}} }{{else}}{{ .Howpublished }}{{end}}{{if .Month}}, {{ .Month }}{{end}}{{if .Year}} {{ .Year }}{{end}}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	case PublicationFieldNameMap[UnPublishedField]:
		t, _ := template.New(PublicationFieldNameMap[UnPublishedField]).Funcs(fMap).Parse("{{ SanitizeText .Author }}. {{if .Title}}{{ SanitizeText .Title }}. {{end}}{{if .Note}}{{ .Note }}{{end}}{{if .Year}}, {{ .Year }}{{end}}.")
		var parsedStr bytes.Buffer
		t.Execute(&parsedStr, item.Value)

		return parsedStr.String()

	default:
		return ""
	}
}
