package profileio

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/buger/jsonparser"
)

// PopulateProfile maps json data to profile struct.
func PopulateProfile(jsonData *[]byte, profile *Profile) {
	// Config
	config, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[ConfigField])
	json.Unmarshal(config, &profile.Config)

	// Basics
	basics, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[BasicsField])
	json.Unmarshal(basics, &profile.Basics)

	// Work
	work, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[WorkField])
	json.Unmarshal(work, &profile.Work)

	// Education
	education, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[EducationField])
	json.Unmarshal(education, &profile.Education)

	// Projects
	projects, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[ProjectsField])
	json.Unmarshal(projects, &profile.Projects)

	// Awards
	awards, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[AwardsField])
	json.Unmarshal(awards, &profile.Awards)

	// Publications
	publications, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[PublicationsField])
	json.Unmarshal(publications, &profile.Publications)

	// Skills
	skills, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[SkillsField])
	json.Unmarshal(skills, &profile.Skills)

	// Languages
	languages, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[LanguagesField])
	json.Unmarshal(languages, &profile.Languages)

	// Interests
	interests, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[InterestsField])
	json.Unmarshal(interests, &profile.Interests)

	// References
	references, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[ReferencesField])
	json.Unmarshal(references, &profile.References)

	// Custom
	custom, _, _, _ := jsonparser.Get(*jsonData, profileFieldNameMap[CustomField])
	json.Unmarshal(custom, &profile.Custom)
}

// preProcessProfile data
func preProcessProfile(profile *Profile) {
	if profile.Config.Theme.Value == profileThemes[BasicTheme] {
		profile.Basics.Summary.Label = SanitizeLabel(profile.Basics.Summary.Label)
		profile.Basics.Summary.Value = SanitizeText(profile.Basics.Summary.Value)
	} else if profile.Config.Theme.Value == profileThemes[PantherTheme] {
		profile.Basics.Summary.Label = SanitizeText(profile.Basics.Summary.Label)
		profile.Basics.Summary.Value = SanitizeText(profile.Basics.Summary.Value)
		profile.Education.Label = SanitizeText(profile.Education.Label)
		profile.Work.Label = SanitizeText(profile.Work.Label)
		profile.Publications.Label = SanitizeText(profile.Publications.Label)
		profile.Projects.Label = SanitizeText(profile.Projects.Label)
		profile.Awards.Label = SanitizeText(profile.Awards.Label)
		profile.Skills.Label = SanitizeText(profile.Skills.Label)
		profile.Languages.Label = SanitizeText(profile.Languages.Label)
		profile.Interests.Label = SanitizeText(profile.Interests.Label)
		profile.References.Label = SanitizeText(profile.References.Label)
	}
}

// parseName func
func parseName(name string, theme string) string {
	switch theme {
	case profileThemes[BasicTheme]:
		return SanitizeLabel(name)

	case profileThemes[PantherTheme]:
		if len(strings.SplitN(name, " ", -1)) == 2 {
			strs := strings.SplitN(name, " ", -1)
			return strs[0] + "\\textbf{" + strs[1] + "}"
		}

		return name

	default:
		return SanitizeLabel(name)
	}
}

// mapLanguageFluency test
func mapLanguageFluency(languages []LanguageDetail) map[string]string {
	m := make(map[string]string)

	for _, language := range languages {
		if language.Render {
			if m[language.Value.Fluency] == "" {
				m[language.Value.Fluency] = language.Value.Language
			} else {
				m[language.Value.Fluency] += ", " + language.Value.Language
			}
		}
	}

	return m
}

// isValidURL test url
func isValidURL(text string) bool {
	_, err := url.ParseRequestURI(text)
	if err != nil {
		return false
	}

	url, err := url.Parse(text)
	if err != nil || url.Scheme == "" || url.Host == "" {
		return false
	}

	return true
}

// makeTitle string
func makeTitle(text string) string {
	return strings.Title(strings.ToLower(text))
}

// SanitizeText updates string by escaping required characters such as '\', '&'.
func SanitizeText(text string) string {
	// Backslash
	text = strings.Replace(text, "\\", "\\textbackslash ", -1)
	// Replace & with \\&
	text = strings.Replace(text, "&", "\\&", -1)

	return text
}

// SanitizeAndMakeTitle updates a string using strings.Title(...) and escapes required characters.
func SanitizeAndMakeTitle(text string) string {
	return SanitizeText(makeTitle(text))
}

// SanitizeLabel returns escaped string for required characters, for a basic theme.
func SanitizeLabel(label string) string {
	if label == "" {
		return label
	}

	label = strings.ToUpper(label)

	strs := strings.SplitN(label, " ", -1)
	label = ""
	for _, str := range strs {
		str = "\\large " + string(str[0]) + "\\small {" + string(str[1:]) + "}"
		label += str + " "
	}

	label = strings.Replace(label, "&", "\\&", -1)

	return label
}

// Parse returnes parsed Publication field based on the theme.
func (item *PublicationDetail) Parse(config Config) string {
	switch config.Theme.Value {
	case profileThemes[BasicTheme]:
		return item.ParseBasic(config)

	case profileThemes[PantherTheme]:
		return item.ParsePanther(config)

	default:
		return item.ParseBasic(config)
	}
}

// Sanitize escapes required characters for ListDetail type.
func (s *ListDetail) Sanitize() {
	s.Value.Brief = SanitizeText(s.Value.Brief)
	s.Value.Detail = SanitizeText(s.Value.Detail)
}

// HasRender returns true if any item in the Work list needs to be rendered.
func (s WorkDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Publications list needs to be rendered.
func (s PublicationDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Projects list needs to be rendered.
func (s ProjectDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Education list needs to be rendered.
func (s EducationDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Awards list needs to be rendered.
func (s AwardDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Skills list needs to be rendered.
func (s SkillDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Languages list needs to be rendered.
func (s LanguageDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Interests list needs to be rendered.
func (s InterestDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the Reference list needs to be rendered.
func (s ReferenceDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}

// HasRender returns true if any item in the list needs to be rendered.
func (s ListDetailSlice) HasRender() bool {
	for _, sDetail := range s {
		if sDetail.Render {
			return true
		}
	}

	return false
}
