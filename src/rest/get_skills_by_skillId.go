package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent"

	"github.com/sky0621/cv-admin/src/ent/skill"
)

// GetSkillsBySkillId 指定スキル取得
// (GET /skills/{bySkillId})
func (s *strictServerImpl) GetSkillsBySkillId(ctx context.Context, request GetSkillsBySkillIdRequestObject) (GetSkillsBySkillIdResponseObject, error) {
	entSkill, err := s.dbClient.Skill.Query().Where(skill.ID(request.BySkillId)).WithSkillTag().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return GetSkillsBySkillId404JSONResponse{n404("skill is none")}, nil
		}
		return nil, err
	}
	if entSkill == nil {
		return GetSkillsBySkillId404JSONResponse{n404("skill is none")}, nil
	}

	return GetSkillsBySkillId200JSONResponse(ToSwaggerSkill(*entSkill)), nil
}
