package helper

import "github.com/sky0621/cv-admin/src/ent"

func PickupUserQualificationIDs(qualifications []*ent.UserQualification) []int {
	results := make([]int, len(qualifications))
	for _, q := range qualifications {
		results = append(results, q.ID)
	}
	return results
}

func PickupUserCareerGroupIDs(careerGroups []*ent.UserCareerGroup) []int {
	results := make([]int, len(careerGroups))
	for _, cg := range careerGroups {
		results = append(results, cg.ID)
	}
	return results
}

func PickupUserCareerIDs(careers []*ent.UserCareer) []int {
	results := make([]int, len(careers))
	for _, cg := range careers {
		results = append(results, cg.ID)
	}
	return results
}
