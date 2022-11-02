package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *ServerImpl) PostSkills(ctx echo.Context) error {
	var skill Skill
	if err := ctx.Bind(&skill); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	rCtx := ctx.Request().Context()

	skills, err := s.dbClient.Skill.Query().All(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	if err := skill.Validate(skills); err != nil {
		return sendClientError(ctx, http.StatusBadRequest, err.Error())
	}

	entSkill, err := ToEntSkillCreate(skill, s.dbClient.Skill.Create()).Save(rCtx)
	if err != nil {
		return sendClientError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, ToSwaggerSkill(entSkill))
}
