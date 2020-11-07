package profileio

import (
	"os"
	"strconv"
)

func generateTemplateBasic(profile *Profile, sortedSectionList []SectionIndexRank, templateFile string) {
	f, err := os.Create(templateFile)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	f.WriteString(`
	\documentclass{res}
	\usepackage{fancyhdr}
	\usepackage{multicol}
	\usepackage{setspace}
	\usepackage[hidelinks]{hyperref}
	\usepackage{enumitem}
	\setlength{\textwidth}{6.5in}
	
	\newcommand{\HRule}{\rule{\linewidth}{0.25pt}}
	\newsectionwidth{0.0in}
	\sectionskip=0.15in

	{{if .Config.Theme.Meta.ShowPageNumbers}}\pagestyle{fancy}
	\pagenumbering{arabic}{{end}}
	
	\begin{document}
	\renewcommand{\headrulewidth}{0.0pt}
	
	% Center the name over the entire width of resume:
	\moveleft.5\hoffset\centerline{ {{- SanitizeLabel .Basics.Name.Value -}} }
	\moveleft.5\hoffset\centerline{\url{ {{.Config.Homepage}} }}
	{{if or .Basics.Email.Render .Basics.Phone.Render }}\moveleft.5\hoffset\centerline{ {{if .Basics.Email.Render}}\href{mailto: {{- .Basics.Email.Value -}}}{ {{- .Basics.Email.Value -}} }{{end}}{{if and .Basics.Email.Render .Basics.Phone.Render}} $|$ {{end}}{{if .Basics.Phone.Render}}{{.Basics.Phone.Value}}{{end}} }{{end}}
	
	{{ $config := .Config }}\begin{resume}
	
	{{if .Basics.Summary.Render}}
	% Summary
	\section{\sc {{ .Basics.Summary.Label }}} \smallskip
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}
	{{ .Basics.Summary.Value }}
	{{end}}
	`)

	for _, data := range sortedSectionList {
		switch data.name {
		case ProfileFieldNameMap[WorkField]:
			profile.Work.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[EducationField]:
			profile.Education.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[PublicationsField]:
			profile.Publications.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[ProjectsField]:
			profile.Projects.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[AwardsField]:
			profile.Awards.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[SkillsField]:
			profile.Skills.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[LanguagesField]:
			profile.Languages.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[InterestsField]:
			profile.Interests.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[ReferencesField]:
			profile.References.List.writeTemplateBasic(f)

		case ProfileFieldNameMap[CustomField]:
			profile.Custom.writeTemplateBasic(f, data.index)

		default:
			info := data.name + " is NOT AVAILABLE"
			panic(info)
		}
	}

	f.WriteString(`
	\end{resume}
	\end{document}
	`)
}

func (s *WorkDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Work}}{{if HasRender .Work.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Experience %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Work.Label -}} } \smallskip
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}
	{{range $index, $item := .Work.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Experience {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{$item.Value.Position -}} } \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end}} \\
	{\it {{$item.Value.Name}}{{if $item.Value.Location}}, {{$item.Value.Location}}{{end -}} }
	\begin{itemize}[leftmargin=*]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}{{if $subItem}}
		\item {{ $subItem.Detail }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	{{end}}{{end}}
	`)
}

func (s *EducationDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Education}}{{if HasRender .Education.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Education %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Education.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	{{range $index, $item := .Education.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Education {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0mm} \smallskip
		\item[] { {{- $item.Value.Institution}}, \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end -}} }
		\item[] {\it {{$item.Value.Degree}} {{$item.Value.Major}}{{if $item.Value.Minor}}, {{$item.Value.Minor}} (minor){{end}} \hfill {{if or $item.Value.Grade $item.Value.GradeTotal}}GPA: {{ParseGrade $item.Value.Grade}}{{if and $item.Value.Grade $item.Value.GradeTotal}}/{{end}}{{$item.Value.GradeTotal}}{{end -}} }
	\end{itemize}
	{{end}}{{end}}
	{{end}}{{end}}
	`)
}

func (s *PublicationDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Publications}}{{if HasRender .Publications.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Publications %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Publications.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
    \begin{itemize}[leftmargin=0pt]
    \setlength{\itemsep}{2pt} \smallskip
	{{ $length := len .Publications.List }}{{range $index, $item := .Publications.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Publication {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item[] {{ $parsedPublication := ParsePublication $item $config }}{{if $parsedPublication}}{{ $parsedPublication }}{{end}}{{end}}{{end}}
    \end{itemize}
	{{end}}{{end}}
	`)
}

func (s *ProjectDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Projects}}{{if HasRender .Projects.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Projects %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Projects.Label -}} } \smallskip
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}{{range $index, $item := .Projects.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Project {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{ $item.Value.Name -}} } \\
	{{if $item.Value.Team}}{{ $item.Value.Team }}{{end}}{{if $item.Value.Note}} {\it {{ $item.Value.Note -}} }{{end}} \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end}}
	{{if $item.Value.Highlights}}{{ $length := len $item.Value.Highlights }}{{if gt $length 0}}\begin{itemize}[leftmargin=*]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}
	\item {{ $subItem.Detail }}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}
	`)
}

func (s *AwardDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Awards}}{{if HasRender .Awards.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Awards %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Awards.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=*]
	\setlength{\itemsep}{2pt} \smallskip
	{{range $index, $item := .Awards.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Award {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item {{ $item.Value.Title }}.{{if $item.Value.Summary}} {\it {{ SanitizeText $item.Value.Summary -}} }.{{end}} {{if $item.Value.Awarder}}{{ $item.Value.Awarder }}, {{end}}{{ $item.Value.Date }}{{end}}{{end}}
		\end{itemize}
	{{end}}{{end}}
	`)
}

func (s *SkillDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Skills}}{{if HasRender .Skills.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Skills %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Skills.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0pt} \smallskip
	{{ $length := len .Skills.List }}{{range $index, $item := .Skills.List}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Skill {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{if $item.Render}}{\bf {{ $item.Value.Name -}} :} {{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	`)
}

func (s *LanguageDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Languages}}{{if HasRender .Languages.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Languages %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Languages.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0pt} \smallskip
	{{ $languagesMap := MapLanguageFluency .Languages.List }}{{ $length := len $languagesMap }}{{ $count := 0 }}{{range $key, $value := $languagesMap}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Language {{ $count }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{ $key }}: {{ $value }}{{end}}
	\end{itemize}{{end}}{{end}}
	`)
}

func (s *InterestDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .Interests}}{{if HasRender .Interests.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Interests %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .Interests.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0pt} \smallskip
	{{ $length := len .Interests.List }}{{range $index, $item := .Interests.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Interest {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{ $item.Value.Name -}} : {{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	`)
}

func (s *ReferenceDetailSlice) writeTemplateBasic(f *os.File) {
	f.WriteString(`
	{{if .References}}{{if HasRender .References.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% References %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel .References.Label -}} } \smallskip
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}{{range $index, $item := .References.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Reference {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{{ $item.Value.Title }} {{ $item.Value.Name }} \\
	{{if $item.Value.Affiliation}}{{ $item.Value.Affiliation }} \\{{end}}
	{{if $item.Value.Address}}{{ $item.Value.Address }} \\{{end}}
	{{if $item.Value.City}}{{ $item.Value.City }}, {{end}}{{if $item.Value.PostalCode}}{{ $item.Value.PostalCode }}{{end}}{{if or $item.Value.City $item.Value.PostalCode}} \\{{end}}
	{{if $item.Value.Email}}\href{mailto: {{- $item.Value.Email -}} }{ {{- $item.Value.Email -}} }{{end}}
	{{end}}{{end}}
	{{end}}{{end}}
	`)
}

func (s *CustomSlice) writeTemplateBasic(f *os.File, index int) {
	switch (*s)[index].Type {
	case ProfileFieldNameMap[WorkField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Work}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Work/Experience: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label -}} } \smallskip
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}
	{{range $index, $item := $customSection.Work}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Experience {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{$item.Value.Position -}} } \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end}} \\
	{\it {{$item.Value.Name}}{{if $item.Value.Location}}, {{$item.Value.Location}}{{end -}} }
	\begin{itemize}[leftmargin=*]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}{{if $subItem}}
		\item {{ $subItem.Detail }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[EducationField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Education}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Education: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	{{range $index, $item := $customSection.Education}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Education {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0mm} \smallskip
		\item[] { {{- $item.Value.Institution}}, \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end -}} }
		\item[] {\it {{$item.Value.Degree}} {{$item.Value.Major}}{{if $item.Value.Minor}}, {{$item.Value.Minor}} (minor){{end}} \hfill {{if or $item.Value.Grade $item.Value.GradeTotal}}GPA: {{ParseGrade $item.Value.Grade}}{{if and $item.Value.Grade $item.Value.GradeTotal}}/{{end}}{{$item.Value.GradeTotal}}{{end -}} }
	\end{itemize}
	{{end}}{{end}}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[PublicationsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Publications}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Publications: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
    \begin{itemize}[leftmargin=0pt]
    \setlength{\itemsep}{2pt} \smallskip
	{{ $length := len $customSection.Publications }}{{range $index, $item := $customSection.Publications}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Publication {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item[] {{ $parsedPublication := ParsePublication $item $config }}{{if $parsedPublication}}{{ $parsedPublication }}{{end}}{{end}}{{end}}
    \end{itemize}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[ProjectsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Projects}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Projects: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label -}} } \smallskip
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}{{range $index, $item := $customSection.Projects}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Project {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{ $item.Value.Name -}} } \\
	{{if $item.Value.Team}}{{ $item.Value.Team }}{{end}}{{if $item.Value.Note}} {\it {{ $item.Value.Note -}} }{{end}} \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end}}
	{{if $item.Value.Highlights}}{{ $length := len $item.Value.Highlights }}{{if gt $length 0}}\begin{itemize}[leftmargin=*]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}
	\item {{ $subItem.Detail }}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[AwardsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Awards}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Awards: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=*]
	\setlength{\itemsep}{2pt} \smallskip
	{{range $index, $item := $customSection.Awards}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Award {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item {{ $item.Value.Title }}.{{if $item.Value.Summary}} {\it {{ SanitizeText $item.Value.Summary -}} }.{{end}} {{if $item.Value.Awarder}}{{ $item.Value.Awarder }}, {{end}}{{ $item.Value.Date }}{{end}}{{end}}
		\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[SkillsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Skills}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Skills: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label }}} {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0pt} \smallskip
	{{ $length := len $customSection.Skills }}{{range $index, $item := $customSection.Skills}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Skill {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{if $item.Render}}{\bf {{ $item.Value.Name -}} :} {{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[LanguagesField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Languages}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Languages: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label -}} } {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0pt} \smallskip
	{{ $languagesMap := MapLanguageFluency $customSection.Languages }}{{ $length := len $languagesMap }}{{ $count := 0 }}{{range $key, $value := $languagesMap}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Language {{ $count }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{ $key }}: {{ $value }}{{end}}
	\end{itemize}{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[InterestsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Interests}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Interests: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label }}} {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.15in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.05in{{end}}
	\begin{itemize}[leftmargin=0pt]
	\setlength{\itemsep}{0pt} \smallskip
	{{ $length := len $customSection.Interests }}{{range $index, $item := $customSection.Interests}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Interest {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{ $item.Value.Name -}} : {{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[ReferencesField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.References}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% References: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label }}} \smallskip
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}{{range $index, $item := $customSection.References}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Reference {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{{ $item.Value.Title }} {{ $item.Value.Name }} \\
	{{if $item.Value.Affiliation}}{{ $item.Value.Affiliation }} \\{{end}}
	{{if $item.Value.Address}}{{ $item.Value.Address }} \\{{end}}
	{{if $item.Value.City}}{{ $item.Value.City }}, {{end}}{{if $item.Value.PostalCode}}{{ $item.Value.PostalCode }}{{end}}{{if or $item.Value.City $item.Value.PostalCode}} \\{{end}}
	{{if $item.Value.Email}}\href{mailto: {{- $item.Value.Email -}} }{ {{- $item.Value.Email -}} }{{end}}
	{{end}}{{end}}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[ListField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% List: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\sc {{ SanitizeLabel $customSection.Label }}} {{if not $config.Theme.Meta.HideSectionLines}}\smallskip{{else}}\vskip 0.2in{{end}}
	{{if not $config.Theme.Meta.HideSectionLines}}\moveleft\hoffset\vbox{\hrule width\resumewidth height 0.25pt} \vskip -0.15in{{end}}
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{6pt}
	{{range $index, $item := $customSection.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item ~~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{if $item.Value.Brief}}{\bf {{ $item.Value.Brief -}} }. {{end}}{{ $item.Value.Detail }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	default:
		info := "Type: \"" + (*s)[index].Type + "\" in \"custom\" section is NOT AVAILABLE"
		panic(info)
	}
}
