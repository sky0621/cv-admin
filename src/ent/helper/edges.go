package helper

import "github.com/sky0621/cv-admin/src/ent"

func PickupUserQualificationIDs(qualifications []*ent.UserQualification) []int {
	results := make([]int, len(qualifications))
	for _, q := range qualifications {
		results = append(results, q.ID)
	}
	return results
}
