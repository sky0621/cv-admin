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

func PickupCareerTaskIDs(tasks []*ent.CareerTask) []int {
	results := make([]int, len(tasks))
	for _, t := range tasks {
		results = append(results, t.ID)
	}
	return results
}

func PickupCareerSkillGroupIDs(skillGroups []*ent.CareerSkillGroup) []int {
	results := make([]int, len(skillGroups))
	for _, sg := range skillGroups {
		results = append(results, sg.ID)
	}
	return results
}
