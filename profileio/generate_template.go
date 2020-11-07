package profileio

//GenerateTemplate writes template file for a given theme.
func GenerateTemplate(profile *Profile, sortedSectionList *[]SectionIndexRank, templateFile string) {
	switch profile.Config.Theme.Value {
	case ProfileThemes[BasicTheme]:
		generateTemplateBasic(profile, *sortedSectionList, templateFile)

	case ProfileThemes[PantherTheme]:
		generateTemplatePanther(profile, *sortedSectionList, templateFile)
	}
}
