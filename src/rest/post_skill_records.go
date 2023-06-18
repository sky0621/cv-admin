package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/converter"
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/golang-utils/slice"
)

// PostSkillrecords スキル新規一括登録
// スキル一括登録
// (POST /skillrecords)
func (s *strictServerImpl) PostSkillrecords(ctx context.Context, request PostSkillrecordsRequestObject) (PostSkillrecordsResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	skills := *request.Body

	postSkillrecords400JSONResponse := func(msg string) PostSkillrecords400JSONResponse {
		return PostSkillrecords400JSONResponse{N400BadRequestJSONResponse{Message: converter.ToPtr(msg)}}
	}

	// TODO: やり方考える。
	for _, skill_ := range skills {
		if err := skill_.Validate(); err != nil {
			return postSkillrecords400JSONResponse("validation failed"), err
		}
	}

	entSkills, err := s.dbClient.Skill.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: やり方考える。
	for _, skill_ := range skills {
		if slice.Contains(helper.PickupSkillName(entSkills), *skill_.Name) {
			return postSkillrecords400JSONResponse("already registered under the same name"), nil
		}
		if slice.Contains(helper.PickupSkillKey(entSkills), *skill_.Key) {
			return postSkillrecords400JSONResponse("already registered under the same key"), nil
		}
	}

	var createdEntSkills []*ent.Skill
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		var entSkillCreates []*ent.SkillCreate
		for _, skill_ := range skills {
			entSkillCreates = append(entSkillCreates, ToEntSkillCreate(skill_, s.dbClient.Skill.Create()))
		}
		createdEntSkills, err = s.dbClient.Skill.CreateBulk(entSkillCreates...).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return PostSkillrecords201JSONResponse{ToSwaggerSkills(createdEntSkills)}, nil
}
