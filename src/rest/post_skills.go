package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/converter"
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

	postSkills400JSONResponse := func(msg string) PostSkills400JSONResponse {
		return PostSkills400JSONResponse{N400BadRequestJSONResponse{Message: converter.ToPtr(msg)}}
	}

	if err := request.Body.Validate(); err != nil {
		return postSkills400JSONResponse("validation failed"), err
	}

	entSkills, err := s.dbClient.Skill.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	if slice.Contains(helper.PickupSkillName(entSkills), *skill_.Name) {
		return postSkills400JSONResponse("already registered under the same name"), nil
	}
	if slice.Contains(helper.PickupSkillKey(entSkills), *skill_.Key) {
		return postSkills400JSONResponse("already registered under the same key"), nil
	}

	entSkill, err := ToEntSkillCreate(skill_, s.dbClient.Skill.Create()).Save(ctx)
	if err != nil {
		return nil, err
	}

	return PostSkills201JSONResponse{N201CreatedSkillJSONResponse(ToSwaggerSkill(entSkill))}, nil
}
