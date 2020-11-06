package profileio

// ProfileTheme type identifier for themes, i.e. basic, panther, etc.
type ProfileTheme int

// Theme identifier consts
const (
	BasicTheme   ProfileTheme = iota
	PantherTheme ProfileTheme = iota
)

// ProfileField type identifier for each section in a profile.
type ProfileField int

// ProfileField identifier consts.
const (
	ConfigField       ProfileField = iota
	BasicsField       ProfileField = iota
	WorkField         ProfileField = iota
	EducationField    ProfileField = iota
	ProjectsField     ProfileField = iota
	AwardsField       ProfileField = iota
	PublicationsField ProfileField = iota
	SkillsField       ProfileField = iota
	LanguagesField    ProfileField = iota
	InterestsField    ProfileField = iota
	ReferencesField   ProfileField = iota
	ListField         ProfileField = iota
	CustomField       ProfileField = iota
)

// PublicationField type identifier for publication fields like article, book, thesis, etc.
type PublicationField int

// article field types
const (
	ArticleField      PublicationField = iota
	BookField         PublicationField = iota
	ThesisField       PublicationField = iota
	TechReportField   PublicationField = iota
	InCollectionField PublicationField = iota
	MiscField         PublicationField = iota
	UnPublishedField  PublicationField = iota
)

// ValueType one of the basic types to store label, value, render info.
type ValueType struct {
	Label  string
	Value  string
	Render bool
}

// ConfigMeta meta data for Config type.
type ConfigMeta struct {
	HideFooterCredit bool
}

// ThemeType defines the type for a theme object.
type ThemeType struct {
	Label  string
	Value  string
	Render bool
	Meta   struct {
		ShowPageNumbers           		bool
		HideSectionLines          		bool
		PantherHeaderNameFontSize 		int
		PantherHeaderColumnOneWidth 	float32
		PantherHeaderColumnTwoWidth 	float32
	}
}

// Config type defining configuration.
type Config struct {
	Homepage string
	Theme    ThemeType
	Meta     ConfigMeta
}

// Location type
type Location struct {
	Value struct {
		Address     string
		PostalCode  string
		City        string
		CountryName string
		Region      string
	}
	Render bool
}

// SocialProfile type
type SocialProfile struct {
	Value struct {
		Network  string
		Username string
		URL      string
	}
	Render bool
}

// Basics type
type Basics struct {
	Name     ValueType
	Label    ValueType
	Email    ValueType
	Phone    ValueType
	URL      ValueType
	Summary  ValueType
	Location Location
	Profiles []SocialProfile
	Rank     int
}

// EducationDetail type
type EducationDetail struct {
	Value struct {
		Institution string
		URL         string
		Major       string
		Minor       string
		Degree      string
		StartDate   string
		EndDate     string
		Grade       float32
		GradeTotal  float32
		Courses     []string
	}
	Render bool
}

// EducationDetailSlice type
type EducationDetailSlice []EducationDetail

// Education type
type Education struct {
	Label string
	List  EducationDetailSlice
	Rank  int
}

// BriefAndDetail type
type BriefAndDetail struct {
	Brief  string
	Detail string
}

// WorkDetail type
type WorkDetail struct {
	Value struct {
		Name       string
		Location   string
		Brief      string
		Position   string
		URL        string
		StartDate  string
		EndDate    string
		Active     bool
		Highlights []BriefAndDetail
	}
	Render bool
}

// WorkDetailSlice type
type WorkDetailSlice []WorkDetail

// Work type
type Work struct {
	Label string
	List  WorkDetailSlice
	Rank  int
}

// ProjectDetail type
type ProjectDetail struct {
	Value struct {
		Name        string
		Description string
		Team        string
		Note        string
		Highlights  []BriefAndDetail
		Keywords    []string
		StartDate   string
		EndDate     string
		URL         string
		Roles       []string
		Type        string
	}
	Render bool
}

// ProjectDetailSlice type
type ProjectDetailSlice []ProjectDetail

// Projects type
type Projects struct {
	Label string
	List  ProjectDetailSlice
	Rank  int
}

// AwardDetail type
type AwardDetail struct {
	Value struct {
		Title   string
		Date    string
		Awarder string
		Summary string
	}
	Render bool
}

// AwardDetailSlice type
type AwardDetailSlice []AwardDetail

// Awards type
type Awards struct {
	Label string
	List  AwardDetailSlice
	Rank  int
}

// Publication article format extends: http://web.mit.edu/rsi/www/pdfs/bibtex-format.pdf

// WebPage type
type WebPage struct {
	Slug string
}

// PublicationDetail type
type PublicationDetail struct {
	Value struct {
		Type         string
		Category     string
		Author       string
		Title        string
		Editor       string
		Edition      string // either edition or series and volume
		Series       string
		Journal      string
		Volume       int
		Booktitle    string
		Publisher    string
		Address      string
		Note         string
		Howpublished string
		Year         int
		Number       int
		Pages        string
		URL          string
		WebPage      WebPage
	}
	Render bool
}

// PublicationDetailSlice type
type PublicationDetailSlice []PublicationDetail

// Publications type
type Publications struct {
	Label string
	List  PublicationDetailSlice
	Rank  int
}

// SkillDetail type
type SkillDetail struct {
	Value struct {
		Name     string
		Level    string
		Keywords []string
	}
	Render bool
}

// SkillDetailSlice type
type SkillDetailSlice []SkillDetail

// Skills type
type Skills struct {
	Label string
	List  SkillDetailSlice
	Rank  int
}

// LanguageDetail type
type LanguageDetail struct {
	Value struct {
		Language string
		Fluency  string
	}
	Render bool
}

// LanguageDetailSlice type
type LanguageDetailSlice []LanguageDetail

// Languages type
type Languages struct {
	Label string
	List  LanguageDetailSlice
	Rank  int
}

// InterestDetail type
type InterestDetail struct {
	Value struct {
		Name     string
		Keywords []string
	}
	Render bool
}

// InterestDetailSlice type
type InterestDetailSlice []InterestDetail

// Interests type
type Interests struct {
	Label string
	List  InterestDetailSlice
	Rank  int
}

// ReferenceDetail type
type ReferenceDetail struct {
	Value struct {
		Name        string
		Title       string
		Affiliation string
		Address     string
		PostalCode  string
		City        string
		CountryName string
		Region      string
		PhoneNumber string
		Email       string
		URL         string
	}
	Render bool
}

// ReferenceDetailSlice type
type ReferenceDetailSlice []ReferenceDetail

// References type
type References struct {
	Label string
	List  ReferenceDetailSlice
	Rank  int
}

// ListDetail type
type ListDetail struct {
	Value struct {
		Brief  string
		Detail string
	}
	Render bool
}

// ListDetailSlice type
type ListDetailSlice []ListDetail

// CustomMeta type
type CustomMeta struct {
	ListStyleType string
}

// Custom type
type Custom struct {
	Label        string
	Type         string
	Work         WorkDetailSlice
	Publications PublicationDetailSlice
	Projects     ProjectDetailSlice
	Education    EducationDetailSlice
	Awards       AwardDetailSlice
	Skills       SkillDetailSlice
	Languages    LanguageDetailSlice
	Interests    InterestDetailSlice
	List         ListDetailSlice
	Meta         CustomMeta
	Rank         int
}

// CustomSlice type
type CustomSlice []Custom

// Profile type
type Profile struct {
	Config       Config
	Basics       Basics
	Work         Work
	Education    Education
	Publications Publications
	Projects     Projects
	Awards       Awards
	Skills       Skills
	Languages    Languages
	Interests    Interests
	References   References
	Custom       CustomSlice
}

// SectionIndexRank type
type SectionIndexRank struct {
	name  string
	index int
	rank  int
}

// Renderable interface
type Renderable interface {
	HasRender() bool
}
