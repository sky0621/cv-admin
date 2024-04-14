package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"
)

// GetSkilltagsBySkillTagId 指定スキルタグ取得
// (GET /skilltags/{bySkillTagId})
func (s *strictServerImpl) GetSkilltagsBySkillTagId(ctx context.Context, request GetSkilltagsBySkillTagIdRequestObject) (GetSkilltagsBySkillTagIdResponseObject, error) {
	entSkillTag, err := s.dbClient.SkillTag.Get(ctx, request.BySkillTagId)
	if err != nil {
		if ent.IsNotFound(err) {
			return GetSkilltagsBySkillTagId404JSONResponse{n404("skillTag is none")}, nil
		}
		return nil, err
	}
	if entSkillTag == nil {
		return GetSkilltagsBySkillTagId404JSONResponse{n404("skillTag is none")}, nil
	}

	return GetSkilltagsBySkillTagId200JSONResponse(ToSwaggerSkillTag(entSkillTag)), nil
}
