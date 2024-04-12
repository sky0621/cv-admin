package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/ent/helper"
	"github.com/sky0621/golang-utils/slice"
)

// PostSkilltags スキルタグ新規登録
// スキルタグ登録
// (POST /skilltags)
func (s *strictServerImpl) PostSkilltags(ctx context.Context, request PostSkilltagsRequestObject) (PostSkilltagsResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is nil")
	}
	skillTag := *request.Body

	if err := skillTag.Validate(); err != nil {
		return PostSkilltags400JSONResponse{n400("validation failed")}, err
	}

	entSkillTags, err := s.dbClient.SkillTag.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	if slice.Contains(helper.PickupSkillTagName(entSkillTags), *skillTag.Name) {
		return PostSkilltags400JSONResponse{n400("already registered under the same name")}, nil
	}

	entSkillTag, err := ToEntSkillTagCreate(skillTag, s.dbClient.SkillTag.Create()).Save(ctx)
	if err != nil {
		return nil, err
	}

	return PostSkilltags201JSONResponse(ToSwaggerSkillTag(entSkillTag)), nil
}
