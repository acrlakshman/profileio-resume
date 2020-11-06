package profileio

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// ProfileIO processes the given json file and generates tex and pdf resume files.
func ProfileIO(jsonData []byte) {
	profile := Profile{}

	// Map json data to profile object
	PopulateProfile(&jsonData, &profile)

	// Populate sorted section list
	sortedSectionList := GetSortedSectionArray(&profile)

	// preProcessProfile data based on themes
	preProcessProfile(&profile)

	var templateName string

	if profile.Config.Theme.Value == profileThemes[BasicTheme] {
		templateName = profileThemes[BasicTheme] + ".tmpl"
	} else if profile.Config.Theme.Value == profileThemes[PantherTheme] {
		templateName = profileThemes[PantherTheme] + ".tmpl"
	} else {
		templateName = profileThemes[BasicTheme] + ".tmpl"
	}

	templateFile := "./" + templateName

	// Generate template file
	GenerateTemplate(&profile, &sortedSectionList, templateFile)

	t, err := template.New(templateName).Funcs(template.FuncMap{
		"HasRender": func(r Renderable) bool {
			return r.HasRender()
		},
		"ParseName": func(name string, theme string) string {
			return parseName(name, theme)
		},
		"ParseGrade": func(value float32) string {
			return fmt.Sprintf("%.2f", value)
		},
		"ParsePublication": func(item PublicationDetail, config Config) string {
			return item.Parse(config)
		},
		"SanitizeLabel": func(text string) string {
			return SanitizeLabel(text)
		},
		"SanitizeText": func(text string) string {
			return SanitizeText(text)
		},
		"Join": func(strArray []string, sep string) string {
			return strings.Join(strArray[:], sep)
		},
		"Inc": func(num int, inc int) int {
			return num + inc
		},
		"MapLanguageFluency": func(l []LanguageDetail) map[string]string {
			return mapLanguageFluency(l)
		},
	}).ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	resumeTeXFile := resumeRenderPath + "resume.tex"

	if _, err := os.Stat(resumeRenderPath); os.IsNotExist(err) {
		err := os.MkdirAll(resumeRenderPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	// Write class file.
	WriteResumeClass(resumeRenderPath + "res.cls")

	f, err := os.Create(resumeTeXFile)
	err = t.Execute(f, profile)
	f.Close()
	if err != nil {
		panic(err)
	}
}
