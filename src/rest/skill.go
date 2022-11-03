package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sky0621/cv-admin/src/ent/helper"
)

func (s *ServerImpl) PostSkills(ctx echo.Context) error {
	var skill Skill
	if err := ctx.Bind(&skill); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	rCtx := ctx.Request().Context()

	if err := skill.Validate(); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	skills, err := s.dbClient.Skill.Query().All(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}
	for _, s := range helper.PickupSkillName(skills) {
		if s == *skill.Name {
			return sendClientError(ctx, http.StatusBadRequest, "already registered under the same name")
		}
	}

	entSkill, err := ToEntSkillCreate(skill, s.dbClient.Skill.Create()).Save(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerSkill(entSkill))
}
