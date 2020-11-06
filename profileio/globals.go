package profileio

var resumeRenderPath string = `./resume/`

var profileThemes map[ProfileTheme]string = map[ProfileTheme]string{
	BasicTheme:   "basic",
	PantherTheme: "panther"}

var profileFieldNameMap map[ProfileField]string = map[ProfileField]string{
	ConfigField:       "config",
	BasicsField:       "basics",
	WorkField:         "work",
	EducationField:    "education",
	PublicationsField: "publications",
	ProjectsField:     "projects",
	AwardsField:       "awards",
	SkillsField:       "skills",
	LanguagesField:    "languages",
	InterestsField:    "interests",
	ReferencesField:   "references",
	ListField:         "list",
	CustomField:       "custom"}

var publicationFieldNameMap map[PublicationField]string = map[PublicationField]string{
	ArticleField:      "article",
	BookField:         "book",
	ThesisField:       "thesis",
	TechReportField:   "techreport",
	InCollectionField: "incollection",
	MiscField:         "misc",
	UnPublishedField:  "unpublished"}
