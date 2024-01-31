package rest

import (
	"context"
	"errors"

	"github.com/sky0621/cv-admin/src/ent/skilltag"
)

// GetSkilltagsBySkillTagId 指定スキルタグ取得
// (GET /skilltags/{bySkillTagId})
func (s *strictServerImpl) GetSkilltagsBySkillTagId(ctx context.Context, request GetSkilltagsBySkillTagIdRequestObject) (GetSkilltagsBySkillTagIdResponseObject, error) {
	entSkillTag, err := s.dbClient.SkillTag.Query().Where(skilltag.ID(request.BySkillTagId)).Only(ctx)
	if err != nil {
		switch {
		case errors.As(err, &notFound):
			return GetSkilltagsBySkillTagId404JSONResponse{n404("skillTag is none")}, err
		}
		return nil, err
	}
	if entSkillTag == nil {
		return GetSkilltagsBySkillTagId404JSONResponse{n404("skillTag is none")}, nil
	}

	return GetSkilltagsBySkillTagId200JSONResponse(ToSwaggerSkillTag(entSkillTag)), nil
}
