package rest

import (
	"context"
	"errors"

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

	// TODO: やり方考える。
	for _, skill_ := range skills {
		if err := skill_.Validate(); err != nil {
			return PostSkillrecords400JSONResponse{n400("validation failed")}, err
		}
	}

	entSkills, err := s.dbClient.Skill.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: やり方考える。
	for _, skill_ := range skills {
		if slice.Contains(helper.PickupSkillName(entSkills), *skill_.Name) {
			return PostSkillrecords400JSONResponse{n400("already registered under the same name")}, nil
		}
	}

	var createdEntSkills []*ent.Skill
	if err := helper.WithTransaction(ctx, s.dbClient, func(tx *ent.Tx) error {
		var entSkillCreates []*ent.SkillCreate
		for _, skill_ := range skills {
			entSkillTag, err := s.dbClient.SkillTag.Get(ctx, *skill_.SkillTagID)
			if err != nil {
				return err
			}
			if entSkillTag == nil {
				return errors.New("no skill tag")
			}

			entSkillCreates = append(entSkillCreates, ToEntSkillCreate(skill_, entSkillTag, s.dbClient.Skill.Create()))
		}
		createdEntSkills, err = s.dbClient.Skill.CreateBulk(entSkillCreates...).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return PostSkillrecords201JSONResponse(ToSwaggerSkills(createdEntSkills)), nil
}
