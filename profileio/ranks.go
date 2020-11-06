package profileio

import "math"

// GetDefaultRanks returns default ranks of each section (work, education, publications, etc.)
// with the exception of custom field in the profile object.
func GetDefaultRanks(theme string) map[string]int {
	m := make(map[string]int)

	switch theme {
	case profileThemes[BasicTheme]:
		m[profileFieldNameMap[WorkField]] = 1
		m[profileFieldNameMap[EducationField]] = 2
		m[profileFieldNameMap[PublicationsField]] = 3
		m[profileFieldNameMap[ProjectsField]] = 4
		m[profileFieldNameMap[AwardsField]] = 5
		m[profileFieldNameMap[SkillsField]] = 6
		m[profileFieldNameMap[LanguagesField]] = math.MaxInt32
		m[profileFieldNameMap[InterestsField]] = math.MaxInt32
		m[profileFieldNameMap[ReferencesField]] = math.MaxInt32

		return m

	case profileThemes[PantherTheme]:
		m[profileFieldNameMap[WorkField]] = 1
		m[profileFieldNameMap[EducationField]] = 2
		m[profileFieldNameMap[PublicationsField]] = 3
		m[profileFieldNameMap[ProjectsField]] = 4
		m[profileFieldNameMap[AwardsField]] = 5
		m[profileFieldNameMap[SkillsField]] = 6
		m[profileFieldNameMap[LanguagesField]] = math.MaxInt32
		m[profileFieldNameMap[InterestsField]] = math.MaxInt32
		m[profileFieldNameMap[ReferencesField]] = math.MaxInt32

		return m

	default:
		m[profileFieldNameMap[WorkField]] = 2
		m[profileFieldNameMap[EducationField]] = 1
		m[profileFieldNameMap[PublicationsField]] = 3
		m[profileFieldNameMap[ProjectsField]] = 4
		m[profileFieldNameMap[AwardsField]] = 5
		m[profileFieldNameMap[SkillsField]] = 6
		m[profileFieldNameMap[LanguagesField]] = math.MaxInt32
		m[profileFieldNameMap[InterestsField]] = math.MaxInt32
		m[profileFieldNameMap[ReferencesField]] = math.MaxInt32

		return m
	}
}
