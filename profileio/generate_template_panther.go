package profileio

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

func generateTemplatePanther(profile *Profile, sortedSectionList []SectionIndexRank, templateFile string) {
	f, err := os.Create(templateFile)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	setMainFontLiteral := ""
	if runtime.GOOS == "darwin" {
		setMainFontLiteral = `\setmainfont{Helvetica Neue Light}[ItalicFont=Helvetica Neue Light Italic]`
	}

	headerColumnOneWidth := float32(0.7)
	headerColumnTwoWidth := float32(0.3)

	if profile.Config.Theme.Meta.PantherHeaderColumnOneWidth > 0 {
		headerColumnOneWidth = profile.Config.Theme.Meta.PantherHeaderColumnOneWidth
		headerColumnTwoWidth = profile.Config.Theme.Meta.PantherHeaderColumnTwoWidth
	}

	if headerColumnOneWidth+headerColumnTwoWidth > 1.0 {
		headerColumnOneWidth /= (headerColumnOneWidth + headerColumnTwoWidth)
		headerColumnTwoWidth /= (headerColumnOneWidth + headerColumnTwoWidth)
	}

	headerColumnOneWidthStr := fmt.Sprintf("%.1f", headerColumnOneWidth)
	headerColumnTwoWidthStr := fmt.Sprintf("%.2f", headerColumnTwoWidth+0.04)

	f.WriteString(`
	\documentclass{res}
	\usepackage{fancyhdr}
	\usepackage{fontspec}
	\usepackage{color}
	\usepackage{multicol}
	\usepackage{setspace}
	\usepackage{xcolor}
	\usepackage{hyperref}
	\usepackage{enumitem}` +
		setMainFontLiteral + `
	\setlength{\textwidth}{6.5in}
	
	\newcommand{\HRule}{\rule{\linewidth}{0.25pt}}
	\newsectionwidth{0.25in}

	% format URLs
	\colorlet{ucolor}{yellow!14!black!14!red!72!}
	\urlstyle{same}
	\hypersetup{
		colorlinks=true,
		linkcolor=ucolor,
		urlcolor=ucolor
	}

	{{if .Config.Theme.Meta.ShowPageNumbers}}\pagestyle{fancy}
	\pagenumbering{arabic}{{end}}

	\newcommand{\spacedsc}{\scshape\addfontfeatures{LetterSpace=15}}
	\renewcommand{\sectionfont}{\normalfont\Large\spacedsc\color{gray}\raggedright}
	\renewcommand{\headrulewidth}{0.0pt}
	
	\begin{document}

	\hspace*{-0.05\textwidth}
	\begin{minipage}[t]{` + headerColumnOneWidthStr + `\textwidth}
	\begin{flushleft}
	\vspace*{\fill}
	{\fontsize{ {{- if .Config.Theme.Meta.PantherHeaderNameFontSize}}{{- .Config.Theme.Meta.PantherHeaderNameFontSize -}}{{else -}} 35 {{end -}} }{ {{- if .Config.Theme.Meta.PantherHeaderNameFontSize}}{{- .Config.Theme.Meta.PantherHeaderNameFontSize -}}{{else -}} 35 {{end -}} }\selectfont {{ ParseName .Basics.Name.Value .Config.Theme.Value -}} }
	\end{flushleft}
	\end{minipage}
	\begin{minipage}[t]{` + headerColumnTwoWidthStr + `\textwidth}
	\begin{flushright}
	{{if .Basics.Email.Render}}\href{mailto: {{- .Basics.Email.Value -}} }{ {{- .Basics.Email.Value -}} } \\ {{end}}{{if .Basics.URL.Render}}\url{ {{- .Basics.URL.Value -}} } \\ {{end}}{{if .Basics.Phone.Render}}\textcolor{ucolor}{ {{- .Basics.Phone.Value -}} }{{end}}
	\end{flushright}
	\end{minipage}

	{{ $config := .Config }}\begin{resume}
	
	%%%%%%%%%%%%%%%%%%%%%%%%%% Introduction %%%%%%%%%%%%%%%%%%%%%%%%%%
	{{if .Basics.Summary.Render}}
	\section{\MakeUppercase{ {{- .Basics.Summary.Label -}} }} \vskip 0.15in
	{{ .Basics.Summary.Value }}
	{{end}}
	`)

	for _, data := range sortedSectionList {
		switch data.name {
		case ProfileFieldNameMap[WorkField]:
			profile.Work.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[EducationField]:
			profile.Education.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[PublicationsField]:
			profile.Publications.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[ProjectsField]:
			profile.Projects.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[AwardsField]:
			profile.Awards.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[SkillsField]:
			profile.Skills.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[LanguagesField]:
			profile.Languages.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[InterestsField]:
			profile.Interests.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[ReferencesField]:
			profile.References.List.writeTemplatePanther(f)

		case ProfileFieldNameMap[CustomField]:
			profile.Custom.writeTemplatePanther(f, data.index)

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

func (s *WorkDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Work}}{{if HasRender .Work.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Work/Experience %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Work.Label -}} }} \vskip 0.15in
	{{range $index, $item := .Work.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\hspace*{-0.25in}{\bf {{$item.Value.Position -}} } \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}}-{{end}}{{$item.Value.EndDate -}}{{end}} \\
	\hspace*{-0.25in}{\it {{$item.Value.Name}}{{if $item.Value.Location}}, {{$item.Value.Location}}{{end -}} }
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}{{if $subItem}}
		\item {{if $subItem.Brief}}{\bf {{ SanitizeText $subItem.Brief -}} } {{end}}{{ SanitizeText $subItem.Detail }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	{{end}}{{end}}
	`)
}

func (s *EducationDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Education}}{{if HasRender .Education.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Education %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Education.Label -}} }} \vskip 0.15in
	{{ $length := len .Education.List }}{{range $index, $item := .Education.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{$item.Value.Degree}} {{$item.Value.Major -}} }{{if $item.Value.Minor}}, {{$item.Value.Minor}} (minor){{end}} \hfill {{if or $item.Value.Grade $item.Value.GradeTotal}}GPA: {{ParseGrade $item.Value.Grade}}{{if and $item.Value.Grade $item.Value.GradeTotal}}/{{end}}{{$item.Value.GradeTotal}}{{end}} \\
	{ {{- $item.Value.Institution}} \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end -}} } {{ $indexP1 := Inc $index 1 }}{{if eq $length $indexP1}}{{else}}\\ \\{{end}}{{end}}{{end}}
	{{end}}{{end}}
	`)
}

func (s *PublicationDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Publications}}{{if HasRender .Publications.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Publications %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Publications.Label -}} }} \vskip 0.35in
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{4pt}
	{{ $length := len .Publications.List }}{{range $index, $item := .Publications.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item[] {{ $parsedPublication := ParsePublication $item $config }}{{if $parsedPublication}}{{ $parsedPublication }}{{end}}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	`)
}

func (s *ProjectDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Projects}}{{if HasRender .Projects.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Projects %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Projects.Label -}} }} \vskip 0.15in
	{{range $index, $item := .Projects.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\hspace*{-0.25in}{\bf {{ $item.Value.Name -}} } \\
	\hspace*{-0.25in}{{if $item.Value.Team}}{{ SanitizeText $item.Value.Team }}{{end}}{{if $item.Value.Note}} ( {{- SanitizeText $item.Value.Note -}} ){{end}} \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate -}} - {{- end}}{{$item.Value.EndDate }}{{end}}
	{{if $item.Value.Highlights}}{{ $length := len $item.Value.Highlights }}{{if gt $length 0}}\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}
	\item {{ SanitizeText $subItem.Detail }}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}
	`)
}

func (s *AwardDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Awards}}{{if HasRender .Awards.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Awards %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Awards.Label -}} }} \vskip 0.35in
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}
	{{range $index, $item := .Awards.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{if $item.Value.Title}}{\bf {{ $item.Value.Title -}} }: {{end}}{{if $item.Value.Summary}} {{ SanitizeText $item.Value.Summary }}.{{end}} {{if $item.Value.Awarder}}{{ $item.Value.Awarder }}, {{end}}{{ $item.Value.Date }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	`)
}

func (s *SkillDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Skills}}{{if HasRender .Skills.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Skills %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Skills.Label -}} }} \vskip 0.35in
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}
	{{range $index, $item := .Skills.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item[] {{if $item.Value.Name}}{\bf {{ $item.Value.Name -}} }: {{end}}{{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	`)
}

func (s *LanguageDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Languages}}{{if HasRender .Languages.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Languages %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Languages.Label -}} }} \vskip 0.35in
	{{ $languagesMap := MapLanguageFluency .Languages.List }}{{ $length := len $languagesMap }}{{ $count := 0 }}
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}{{range $key, $value := $languagesMap}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $count }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item[] {{ $key }}: {{ $value }}{{end}}
	\end{itemize}{{end}}
	{{end}}
	`)
}

func (s *InterestDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .Interests}}{{if HasRender .Interests.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Interests %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .Interests.Label -}} }} \vskip 0.35in
	{{ $length := len .Interests.List }}
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}{{range $index, $item := .Interests.List}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item[] {{if $item.Render}}{\bf {{ $item.Value.Name -}} :} {{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}{{end}}{{end}}
	`)
}

func (s *ReferenceDetailSlice) writeTemplatePanther(f *os.File) {
	f.WriteString(`
	{{if .References}}{{if HasRender .References.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% References %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- .References.Label -}} }} \vskip 0.15in
	{{range $index, $item := .References.List}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{if $item.Value.URL}}\href{ {{- $item.Value.URL -}} }{ {{- $item.Value.Title }} {{ $item.Value.Name -}} }{{else}}{{ $item.Value.Title }} {{ $item.Value.Name }}{{end -}} } \\
	{{if $item.Value.Affiliation}}{{ $item.Value.Affiliation }} \\{{end}}
	{{if $item.Value.Address}}{{ $item.Value.Address }} \\{{end}}
	{{if $item.Value.City}}{{ $item.Value.City }}, {{end}}{{if $item.Value.PostalCode}}{{ $item.Value.PostalCode }}{{end}}{{if or $item.Value.City $item.Value.PostalCode}} \\{{end}}
	{{if $item.Value.Email}}\href{mailto: {{- $item.Value.Email -}} }{ {{- $item.Value.Email -}} }{{end}}
	{{end}}
	{{end}}{{end}}
	`)
}

func (s *CustomSlice) writeTemplatePanther(f *os.File, index int) {
	switch (*s)[index].Type {
	case ProfileFieldNameMap[WorkField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Work}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Work/Experience: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.15in
	{{range $index, $item := $customSection.Work}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\hspace*{-0.25in}{\bf {{$item.Value.Position -}} } \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}}-{{end}}{{$item.Value.EndDate -}}{{end}} \\
	\hspace*{-0.25in}{\it {{$item.Value.Name}}{{if $item.Value.Location}}, {{$item.Value.Location}}{{end -}} }
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}{{if $subItem}}
		\item {{if $subItem.Brief}}{\bf {{ SanitizeText $subItem.Brief -}} } {{end}}{{ SanitizeText $subItem.Detail }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[EducationField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Education}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Education: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.15in
	{{ $length := len $customSection.Education }}{{range $index, $item := $customSection.Education}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{$item.Value.Degree}} {{$item.Value.Major -}} }{{if $item.Value.Minor}}, {{$item.Value.Minor}} (minor){{end}} \hfill {{if or $item.Value.Grade $item.Value.GradeTotal}}GPA: {{ParseGrade $item.Value.Grade}}{{if and $item.Value.Grade $item.Value.GradeTotal}}/{{end}}{{$item.Value.GradeTotal}}{{end}} \\
	{ {{- $item.Value.Institution}} \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{\it {{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate}} - {{end}}{{$item.Value.EndDate -}} }{{end -}} } {{ $indexP1 := Inc $index 1 }}{{if eq $length $indexP1}}{{else}}\\ \\{{end}}{{end}}{{end}}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[PublicationsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Publications}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Publications: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.35in
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{4pt}
	{{ $length := len $customSection.Publications }}{{range $index, $item := $customSection.Publications}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{ $parsedPublication := ParsePublication $item $config }}{{if $parsedPublication}}{{ $parsedPublication }}{{end}}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[ProjectsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Projects}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Projects: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.15in
	{{range $index, $item := $customSection.Projects}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\hspace*{-0.25in}{\bf {{ $item.Value.Name -}} } \\
	\hspace*{-0.25in}{{if $item.Value.Team}}{{ SanitizeText $item.Value.Team }}{{end}}{{if $item.Value.Note}} ( {{- SanitizeText $item.Value.Note -}} ){{end}} \hfill {{if or $item.Value.StartDate $item.Value.EndDate}}{{$item.Value.StartDate}}{{if and $item.Value.StartDate $item.Value.EndDate -}} - {{- end}}{{$item.Value.EndDate }}{{end}}
	{{if $item.Value.Highlights}}{{ $length := len $item.Value.Highlights }}{{if gt $length 0}}\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm} \smallskip
	{{range $i, $subItem := $item.Value.Highlights}}
	\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{ SanitizeText $subItem.Detail }}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[AwardsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Awards}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Awards: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.35in
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}
	{{range $index, $item := $customSection.Awards}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{if $item.Value.Title}}{\bf {{ $item.Value.Title -}} }: {{end}}{{if $item.Value.Summary}} {{ SanitizeText $item.Value.Summary }}.{{end}} {{if $item.Value.Awarder}}{{ $item.Value.Awarder }}, {{end}}{{ $item.Value.Date }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[SkillsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Skills}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Skills: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.35in
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}
	{{range $index, $item := $customSection.Skills}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{if $item.Value.Name}}{\bf {{ $item.Value.Name -}} }: {{end}}{{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[LanguagesField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Languages}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Languages: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.35in
	{{ $languagesMap := MapLanguageFluency $customSection.Languages }}{{ $length := len $languagesMap }}{{ $count := 0 }}
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}{{range $key, $value := $languagesMap}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $count }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{ $key }}: {{ $value }}{{end}}
	\end{itemize}{{end}}
	{{end}}{{end}}
		`)

	case ProfileFieldNameMap[InterestsField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.Interests}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% Interests: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.35in
	{{ $length := len $customSection.Interests }}
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{0mm}{{range $index, $item := $customSection.Interests}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{if $item.Render}}{\bf {{ $item.Value.Name -}} :} {{ Join $item.Value.Keywords ", " }}{{end}}{{end}}
	\end{itemize}{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[ReferencesField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.References}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% References: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.15in
	{{range $index, $item := $customSection.References}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item {{ $index }} ~~~~~~~~~~~~~~~~~~~~~~~~~~%
	{\bf {{if $item.Value.URL}}\href{ {{- $item.Value.URL -}} }{ {{- $item.Value.Title }} {{ $item.Value.Name -}} }{{else}}{{ $item.Value.Title }} {{ $item.Value.Name }}{{end -}} } \\
	{{if $item.Value.Affiliation}}{{ $item.Value.Affiliation }} \\{{end}}
	{{if $item.Value.Address}}{{ $item.Value.Address }} \\{{end}}
	{{if $item.Value.City}}{{ $item.Value.City }}, {{end}}{{if $item.Value.PostalCode}}{{ $item.Value.PostalCode }}{{end}}{{if or $item.Value.City $item.Value.PostalCode}} \\{{end}}
	{{if $item.Value.Email}}\href{mailto: {{- $item.Value.Email -}} }{ {{- $item.Value.Email -}} }{{end}}
	{{end}}
	{{end}}{{end}}{{end}}
		`)

	case ProfileFieldNameMap[ListField]:
		f.WriteString(`
	{{if .Custom}}{{with $customSection := index .Custom ` + strconv.Itoa(index) + `}}{{if HasRender $customSection.List}}
	%%%%%%%%%%%%%%%%%%%%%%%%%%% List: Custom[` + strconv.Itoa(index) + `] %%%%%%%%%%%%%%%%%%%%%%%%%%%
	\section{\MakeUppercase{ {{- $customSection.Label -}} }} \vskip 0.35in
	\begin{itemize}[leftmargin=\parindent]
	\setlength{\itemsep}{6pt}
	{{range $index, $item := $customSection.List}}{{if $item.Render}}
	%~~~~~~~~~~~~~~~~~~~~~~~~~~ Item ~~~~~~~~~~~~~~~~~~~~~~~~~~~%
		\item {{- if $customSection.Meta}}{{if $customSection.Meta.ListStyleType}}{{if eq $customSection.Meta.ListStyleType "none"}}[]{{end}}{{end}}{{end}} {{if $item.Value.Brief}}{\bf {{ $item.Value.Brief -}} } {{end}}{{ $item.Value.Detail }}{{end}}{{end}}
	\end{itemize}
	{{end}}{{end}}{{end}}
		`)

	default:
		info := "Type: \"" + (*s)[index].Type + "\" in \"custom\" section is NOT AVAILABLE"
		panic(info)
	}
}
