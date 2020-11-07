package profileio

var resumeRenderPath string = `./resume/`

// ProfileThemes is a map of list of available themes.
var ProfileThemes map[ProfileTheme]string = map[ProfileTheme]string{
	BasicTheme:   "basic",
	PantherTheme: "panther"}

// ProfileFieldNameMap stores the field names of different sections in a profile, e.g. "config", "basics", etc.
var ProfileFieldNameMap map[ProfileField]string = map[ProfileField]string{
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

// PublicationFieldNameMap stores the names of different publication types, "article", "book", etc.
var PublicationFieldNameMap map[PublicationField]string = map[PublicationField]string{
	ArticleField:      "article",
	BookField:         "book",
	ThesisField:       "thesis",
	TechReportField:   "techreport",
	InCollectionField: "incollection",
	MiscField:         "misc",
	UnPublishedField:  "unpublished"}
