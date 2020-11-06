package profileio

import (
	"math"
	"sort"
)

// GetSortedSectionArray returns array of SectionIndexRank type objects sorted by their ranks.
func GetSortedSectionArray(profile *Profile) []SectionIndexRank {
	sortedSectionList := []SectionIndexRank{}

	maxRank := math.MaxInt32
	var name string
	rank := 0
	defaultRanks := GetDefaultRanks(profile.Config.Theme.Value)

	name = profileFieldNameMap[WorkField]
	rank = profile.Work.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[EducationField]
	rank = profile.Education.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[ProjectsField]
	rank = profile.Projects.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[AwardsField]
	rank = profile.Awards.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[PublicationsField]
	rank = profile.Publications.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[SkillsField]
	rank = profile.Skills.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[LanguagesField]
	rank = profile.Languages.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[InterestsField]
	rank = profile.Interests.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[ReferencesField]
	rank = profile.References.Rank
	if rank <= 0 {
		rank = defaultRanks[name]
	}
	sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, rank: rank})

	name = profileFieldNameMap[CustomField]
	for index, customSection := range profile.Custom {
		rank = customSection.Rank
		if rank <= 0 {
			rank = maxRank
		}
		sortedSectionList = append(sortedSectionList, SectionIndexRank{name: name, index: index, rank: rank})
	}

	sort.SliceStable(sortedSectionList, func(i, j int) bool {
		return sortedSectionList[i].rank < sortedSectionList[j].rank
	})

	return sortedSectionList
}
