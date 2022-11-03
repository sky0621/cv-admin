package rest

import (
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/golang-utils/slice"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/ent/helper"
)

// PostSkills スキル新規登録
// (POST /skills)
func (s *ServerImpl) PostSkills(ctx echo.Context) error {
	var skill Skill
	if err := ctx.Bind(&skill); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := skill.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	rCtx := ctx.Request().Context()

	entSkills, err := s.dbClient.Skill.Query().All(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	if slice.Contains(helper.PickupSkillName(entSkills), *skill.Name) {
		return sendClientError(ctx, http.StatusBadRequest, "already registered under the same name")
	}
	if slice.Contains(helper.PickupSkillKey(entSkills), *skill.Key) {
		return sendClientError(ctx, http.StatusBadRequest, "already registered under the same key")
	}

	entSkill, err := ToEntSkillCreate(skill, s.dbClient.Skill.Create()).Save(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerSkill(entSkill))
}

// PostSkillrecords スキル新規一括登録
// (POST /skillrecords)
func (s *ServerImpl) PostSkillrecords(ctx echo.Context) error {
	var skills []Skill
	if err := ctx.Bind(&skills); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	// TODO: やり方考える。
	for _, skill := range skills {
		if err := skill.Validate(); err != nil {
			return sendClientError(ctx, http.StatusBadRequest, err.Error())
		}
	}

	rCtx := ctx.Request().Context()

	entSkills, err := s.dbClient.Skill.Query().All(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	// TODO: やり方考える。
	for _, skill := range skills {
		if slice.Contains(helper.PickupSkillName(entSkills), *skill.Name) {
			return sendClientError(ctx, http.StatusBadRequest, "already registered under the same name")
		}
		if slice.Contains(helper.PickupSkillKey(entSkills), *skill.Key) {
			return sendClientError(ctx, http.StatusBadRequest, "already registered under the same key")
		}
	}

	var createdEntSkills []*ent.Skill
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		var entSkillCreates []*ent.SkillCreate
		for _, skill := range skills {
			entSkillCreates = append(entSkillCreates, ToEntSkillCreate(skill, tx.Skill.Create()))
		}

		createdEntSkills, err = tx.Skill.CreateBulk(entSkillCreates...).Save(rCtx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerSkills(createdEntSkills))
}

func (s *ServerImpl) GetSkills(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServerImpl) GetSkillsBySkillId(ctx echo.Context, bySkillId SkillId) error {
	//TODO implement me
	panic("implement me")
}
