package rest

import "context"

// GetSkills 全スキル取得
// 全スキル取得
// (GET /skills)
func (s *strictServerImpl) GetSkills(ctx context.Context, _ GetSkillsRequestObject) (GetSkillsResponseObject, error) {
	entSkills, err := s.dbClient.Skill.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return GetSkills200JSONResponse{ToSwaggerSkills(entSkills)}, nil
}
