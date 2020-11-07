package profileio

import "math"

// GetDefaultRanks returns default ranks of each section (work, education, publications, etc.)
// with the exception of custom field in the profile object.
func GetDefaultRanks(theme string) map[string]int {
	m := make(map[string]int)

	switch theme {
	case ProfileThemes[BasicTheme]:
		m[ProfileFieldNameMap[WorkField]] = 1
		m[ProfileFieldNameMap[EducationField]] = 2
		m[ProfileFieldNameMap[PublicationsField]] = 3
		m[ProfileFieldNameMap[ProjectsField]] = 4
		m[ProfileFieldNameMap[AwardsField]] = 5
		m[ProfileFieldNameMap[SkillsField]] = 6
		m[ProfileFieldNameMap[LanguagesField]] = math.MaxInt32
		m[ProfileFieldNameMap[InterestsField]] = math.MaxInt32
		m[ProfileFieldNameMap[ReferencesField]] = math.MaxInt32

		return m

	case ProfileThemes[PantherTheme]:
		m[ProfileFieldNameMap[WorkField]] = 1
		m[ProfileFieldNameMap[EducationField]] = 2
		m[ProfileFieldNameMap[PublicationsField]] = 3
		m[ProfileFieldNameMap[ProjectsField]] = 4
		m[ProfileFieldNameMap[AwardsField]] = 5
		m[ProfileFieldNameMap[SkillsField]] = 6
		m[ProfileFieldNameMap[LanguagesField]] = math.MaxInt32
		m[ProfileFieldNameMap[InterestsField]] = math.MaxInt32
		m[ProfileFieldNameMap[ReferencesField]] = math.MaxInt32

		return m

	default:
		m[ProfileFieldNameMap[WorkField]] = 2
		m[ProfileFieldNameMap[EducationField]] = 1
		m[ProfileFieldNameMap[PublicationsField]] = 3
		m[ProfileFieldNameMap[ProjectsField]] = 4
		m[ProfileFieldNameMap[AwardsField]] = 5
		m[ProfileFieldNameMap[SkillsField]] = 6
		m[ProfileFieldNameMap[LanguagesField]] = math.MaxInt32
		m[ProfileFieldNameMap[InterestsField]] = math.MaxInt32
		m[ProfileFieldNameMap[ReferencesField]] = math.MaxInt32

		return m
	}
}
