package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent/skill"
	"github.com/sky0621/cv-admin/src/ent/skilltag"

	"github.com/sky0621/cv-admin/src/ent"
)

// GetSkills 全スキル取得
// 全スキル取得
// (GET /skills)
func (s *strictServerImpl) GetSkills(ctx context.Context, request GetSkillsRequestObject) (GetSkillsResponseObject, error) {
	var entSkills []*ent.Skill
	var err error
	if request.Params.Tag == nil {
		entSkills, err = s.dbClient.Skill.Query().All(ctx)
	} else {
		entSkills, err = s.dbClient.Skill.Query().Where(skill.HasSkillTagWith(skilltag.ID(*request.Params.Tag))).All(ctx)
	}
	if err != nil {
		return nil, err
	}

	return GetSkills200JSONResponse(ToSwaggerSkills(entSkills)), nil
}
