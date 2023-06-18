package rest

import "context"

// GetSkilltags 全スキルタグ取得
// 全スキルタグ取得
// (GET /skilltags)
func (s *strictServerImpl) GetSkilltags(ctx context.Context, _ GetSkilltagsRequestObject) (GetSkilltagsResponseObject, error) {
	entSkillTags, err := s.dbClient.SkillTag.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return GetSkilltags200JSONResponse{ToSwaggerSkillTags(entSkillTags)}, nil
}
