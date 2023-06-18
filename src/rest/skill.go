package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/cv-admin/src/converter"
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/cv-admin/src/ent/skill"
	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usercareergroup"
	"github.com/sky0621/golang-utils/slice"
	"net/http"
)

// PostSkilltags スキルタグ新規登録
// (POST /skilltags)
func (s *ServerImpl) PostSkilltags(ctx echo.Context) (PostSkilltagsResponseObject, error) {
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
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	if slice.Contains(helper.PickupSkillTagName(entSkillTags), *skillTag.Name) {
		return sendClientError(ctx, http.StatusBadRequest, "already registered under the same name")
	}
	if slice.Contains(helper.PickupSkillTagKey(entSkillTags), *skillTag.Key) {
		return sendClientError(ctx, http.StatusBadRequest, "already registered under the same key")
	}

	entSkillTag, err := ToEntSkillTagCreate(skillTag, s.dbClient.SkillTag.Create()).Save(rCtx)
	if err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	return PostSkilltagsResponseObject{ToSwaggerSkillTag(entSkillTag)}, nil
}

// GetSkilltags 全スキルタグ取得
// (GET /skilltags)
func (s *ServerImpl) GetSkilltags(ctx echo.Context) (GetSkilltagsResponseObject, error) {
	entSkillTags, err := s.dbClient.SkillTag.Query().All(ctx.Request().Context())
	if err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	return GetSkilltags200JSONResponse{ToSwaggerSkillTags(entSkillTags)}, nil
}

// PostSkills スキル新規登録
// (POST /skills)
func (s *ServerImpl) PostSkills(ctx echo.Context) (PostSkillsResponseObject, error) {
	postSkills400JSONResponse := func(msg string) PostSkills400JSONResponse {
		return PostSkills400JSONResponse{N400BadRequestJSONResponse{Message: converter.ToPtr(msg)}}
	}

	var skill_ Skill
	if err := ctx.Bind(&skill_); err != nil {
		return postSkills400JSONResponse("bind failed"), err
	}

	if err := skill_.Validate(); err != nil {
		return postSkills400JSONResponse("validation failed"), err
	}

	rCtx := ctx.Request().Context()

	entSkills, err := s.dbClient.Skill.Query().All(rCtx)
	if err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	if slice.Contains(helper.PickupSkillName(entSkills), *skill_.Name) {
		return postSkills400JSONResponse("already registered under the same name"), nil
	}
	if slice.Contains(helper.PickupSkillKey(entSkills), *skill_.Key) {
		return postSkills400JSONResponse("already registered under the same key"), nil
	}

	entSkill, err := ToEntSkillCreate(skill_, s.dbClient.Skill.Create()).Save(rCtx)
	if err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	return PostSkills201JSONResponse{N201CreatedSkillJSONResponse(ToSwaggerSkill(entSkill))}, nil
}

// PostSkillrecords スキル新規一括登録
// (POST /skillrecords)
func (s *ServerImpl) PostSkillrecords(ctx echo.Context) (PostSkillrecordsResponseObject, error) {
	postSkillrecords400JSONResponse := func(msg string) PostSkillrecords400JSONResponse {
		return PostSkillrecords400JSONResponse{N400BadRequestJSONResponse{Message: converter.ToPtr(msg)}}
	}

	var skills []Skill
	if err := ctx.Bind(&skills); err != nil {
		return postSkillrecords400JSONResponse("bind failed"), err
	}

	// TODO: やり方考える。
	for _, skill_ := range skills {
		if err := skill_.Validate(); err != nil {
			return postSkillrecords400JSONResponse("validation failed"), err
		}
	}

	rCtx := ctx.Request().Context()

	entSkills, err := s.dbClient.Skill.Query().All(rCtx)
	if err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
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
	if err := helper.WithTransaction(rCtx, s.dbClient, func(tx *ent.Tx) error {
		var entSkillCreates []*ent.SkillCreate
		for _, skill_ := range skills {
			entSkillCreates = append(entSkillCreates, ToEntSkillCreate(skill_, s.dbClient.Skill.Create()))
		}
		createdEntSkills, err = s.dbClient.Skill.CreateBulk(entSkillCreates...).Save(rCtx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	return PostSkillrecords201JSONResponse{ToSwaggerSkills(createdEntSkills)}, nil
}

// GetSkills 全スキル取得
// (GET /skills)
func (s *ServerImpl) GetSkills(ctx echo.Context) (GetSkillsResponseObject, error) {
	entSkills, err := s.dbClient.Skill.Query().All(ctx.Request().Context())
	if err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	return ctx.JSON(http.StatusOK, ToSwaggerSkills(entSkills))
}

func (s *ServerImpl) GetSkillsBySkillId(ctx echo.Context, bySkillId SkillId) (GetSkillsBySkillIdRequestObject, error) {
	// FIXME: implement me
	panic("implement me")
}

// GetUsersByUserIdSkills スキル群取得
// (GET /users/{byUserId}/skills)
func (s *ServerImpl) GetUsersByUserIdSkills(ctx echo.Context, byUserId UserId) (GetUsersByUserIdSkillsResponseObject, error) {
	rCtx := ctx.Request().Context()
	entSkillTags, err := s.dbClient.SkillTag.Query().All(rCtx)
	if err != nil {
		return nil, sendStatusInternalServerError(ctx, err.Error())
	}

	var userSkillTags []UserSkillTag
	for _, entSkillTag := range entSkillTags {
		entSkills, err := s.dbClient.Skill.Query().Where(skill.TagKey(entSkillTag.Key)).WithCareerSkills(func(q *ent.CareerSkillQuery) {
			q.WithCareerSkillGroup(func(q *ent.CareerSkillGroupQuery) {
				q.WithCareer(func(q *ent.UserCareerQuery) {
					q.WithCareerGroup(func(q *ent.UserCareerGroupQuery) {
						q.Where(usercareergroup.HasUserWith(user.ID(byUserId)))
					})
				})
			})
		}).All(rCtx)
		if err != nil {
			return nil, sendStatusInternalServerError(ctx, err.Error())
		}

		var userSkills []UserSkill
		for _, entSkill := range entSkills {
			var versions []UserSkillVersion
			for _, entCareerSkill := range entSkill.Edges.CareerSkills {
				versions = append(versions, UserSkillVersion{
					Version: entCareerSkill.Version,
					From:    ToSwaggerUserCareerPeriodFrom(entCareerSkill.Edges.CareerSkillGroup.Edges.Career.From),
					To:      ToSwaggerUserCareerPeriodTo(entCareerSkill.Edges.CareerSkillGroup.Edges.Career.To),
				})
			}
			userSkills = append(userSkills, UserSkill{
				Name:     &entSkill.Name,
				Key:      &entSkill.Key,
				Url:      entSkill.URL,
				Versions: &versions,
			})
		}

		userSkillTags = append(userSkillTags, UserSkillTag{
			TagName: &entSkillTag.Name,
			TagKey:  &entSkillTag.Key,
			Skills:  &userSkills,
		})
	}

	return ctx.JSON(http.StatusOK, userSkillTags)
}
