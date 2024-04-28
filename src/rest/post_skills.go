package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/golang-utils/slice"
)

// PostSkillrecords スキル新規一括登録
// スキル登録
// (POST /skills)
func (s *strictServerImpl) PostSkills(ctx context.Context, request PostSkillsRequestObject) (PostSkillsResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	skill_ := *request.Body

	if err := request.Body.Validate(); err != nil {
		return PostSkills400JSONResponse{n400("validation failed")}, err
	}

	entSkills, err := s.dbClient.Skill.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	if slice.Contains(helper.PickupSkillName(entSkills), *skill_.Name) {
		return PostSkills400JSONResponse{n400("already registered under the same name")}, nil
	}

	entSkillTag, err := s.dbClient.SkillTag.Get(ctx, *skill_.SkillTagID)
	if err != nil {
		return nil, err
	}
	if entSkillTag == nil {
		return PostSkills400JSONResponse{n400("no skill tag")}, nil
	}

	entSkill, err := ToEntSkillCreate(skill_, entSkillTag, s.dbClient.Skill.Create()).Save(ctx)
	if err != nil {
		return nil, err
	}
	if entSkill == nil {
		return nil, errors.New("saved skill but return nil")
	}

	return PostSkills201JSONResponse(ToSwaggerSkill(*entSkill)), nil
}
