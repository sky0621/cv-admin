package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/golang-utils/slice"
	"net/http"
)

// PostSkilltags スキルタグ新規登録
// (POST /skilltags)
func (s *ServerImpl) PostSkilltags(ctx echo.Context) error {
	var skillTag SkillTag
	if err := ctx.Bind(&skillTag); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := skillTag.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	rCtx := ctx.Request().Context()

	entSkillTags, err := s.dbClient.SkillTag.Query().All(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	if slice.Contains(helper.PickupSkillTagName(entSkillTags), *skillTag.Name) {
		return sendClientError(ctx, http.StatusBadRequest, "already registered under the same name")
	}
	if slice.Contains(helper.PickupSkillTagKey(entSkillTags), *skillTag.Key) {
		return sendClientError(ctx, http.StatusBadRequest, "already registered under the same key")
	}

	entSkillTag, err := ToEntSkillTagCreate(skillTag, s.dbClient.SkillTag.Create()).Save(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerSkillTag(entSkillTag))
}

// GetSkilltags 全スキルタグ取得
// (GET /skilltags)
func (s *ServerImpl) GetSkilltags(ctx echo.Context) error {
	entSkillTags, err := s.dbClient.SkillTag.Query().All(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerSkillTags(entSkillTags))
}

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
			entSkillCreates = append(entSkillCreates, ToEntSkillCreate(skill, s.dbClient.Skill.Create()))
		}
		createdEntSkills, err = s.dbClient.Skill.CreateBulk(entSkillCreates...).Save(rCtx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerSkills(createdEntSkills))
}

// GetSkills 全スキル取得
// (GET /skills)
func (s *ServerImpl) GetSkills(ctx echo.Context) error {
	entSkills, err := s.dbClient.Skill.Query().All(ctx.Request().Context())
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerSkills(entSkills))
}

func (s *ServerImpl) GetSkillsBySkillId(ctx echo.Context, bySkillId SkillId) error {
	// FIXME: implement me
	panic("implement me")
}

func (s *ServerImpl) GetUsersByUserIdSkills(ctx echo.Context, byUserId UserId) error {
	//TODO implement me
	panic("implement me")
}
