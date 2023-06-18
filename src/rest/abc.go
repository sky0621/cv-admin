package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/converter"

	"github.com/sky0621/cv-admin/src/ent"
)

func n400(msg string) N400BadRequestJSONResponse {
	return N400BadRequestJSONResponse{Message: converter.ToPtr(msg)}
}

func n404(msg string) N404NotFoundJSONResponse {
	return N404NotFoundJSONResponse{Message: converter.ToPtr(msg)}
}

type strictServerImpl struct {
	dbClient *ent.Client
}

// 属性更新
// (PUT /users/{byUserId}/attribute)
func (s *strictServerImpl) PutUsersByUserIdAttribute(ctx context.Context, request PutUsersByUserIdAttributeRequestObject) (PutUsersByUserIdAttributeResponseObject, error) {

}

// キャリアグループ群取得
// (GET /users/{byUserId}/careergroups)
func (s *strictServerImpl) GetUsersByUserIdCareergroups(ctx context.Context, request GetUsersByUserIdCareergroupsRequestObject) (GetUsersByUserIdCareergroupsResponseObject, error) {

}

// キャリアグループ新規登録
// (POST /users/{byUserId}/careergroups)
func (s *strictServerImpl) PostUsersByUserIdCareergroups(ctx context.Context, request PostUsersByUserIdCareergroupsRequestObject) (PostUsersByUserIdCareergroupsResponseObject, error) {

}

// 【未実装】キャリアグループ削除
// (DELETE /users/{byUserId}/careergroups/{byCareerGroupId})
func (s *strictServerImpl) DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx context.Context, request DeleteUsersByUserIdCareergroupsByCareerGroupIdRequestObject) (DeleteUsersByUserIdCareergroupsByCareerGroupIdResponseObject, error) {

}

// 【未実装】キャリアグループ更新
// (PUT /users/{byUserId}/careergroups/{byCareerGroupId})
func (s *strictServerImpl) PutUsersByUserIdCareergroupsByCareerGroupId(ctx context.Context, request PutUsersByUserIdCareergroupsByCareerGroupIdRequestObject) (PutUsersByUserIdCareergroupsByCareerGroupIdResponseObject, error) {

}

// 【未実装】キャリアグループ内キャリア群最新化
// (PUT /users/{byUserId}/careergroups/{byCareerGroupId}/careers)
func (s *strictServerImpl) PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx context.Context, request PutUsersByUserIdCareergroupsByCareerGroupIdCareersRequestObject) (PutUsersByUserIdCareergroupsByCareerGroupIdCareersResponseObject, error) {

}

// 注釈群取得
// (GET /users/{byUserId}/notes)
func (s *strictServerImpl) GetUsersByUserIdNotes(ctx context.Context, request GetUsersByUserIdNotesRequestObject) (GetUsersByUserIdNotesResponseObject, error) {

}

// 注釈新規登録
// (POST /users/{byUserId}/notes)
func (s *strictServerImpl) PostUsersByUserIdNotes(ctx context.Context, request PostUsersByUserIdNotesRequestObject) (PostUsersByUserIdNotesResponseObject, error) {

}

// 【未実装】注釈削除
// (DELETE /users/{byUserId}/notes/{byNoteId})
func (s *strictServerImpl) DeleteUsersByUserIdNotesByNoteId(ctx context.Context, request DeleteUsersByUserIdNotesByNoteIdRequestObject) (DeleteUsersByUserIdNotesByNoteIdResponseObject, error) {

}

// 【未実装】注釈更新
// (PUT /users/{byUserId}/notes/{byNoteId})
func (s *strictServerImpl) PutUsersByUserIdNotesByNoteId(ctx context.Context, request PutUsersByUserIdNotesByNoteIdRequestObject) (PutUsersByUserIdNotesByNoteIdResponseObject, error) {

}

// 【未実装】注釈内要素群最新化
// (PUT /users/{byUserId}/notes/{byNoteId}/items)
func (s *strictServerImpl) PutUsersByUserIdNotesByNoteIdItems(ctx context.Context, request PutUsersByUserIdNotesByNoteIdItemsRequestObject) (PutUsersByUserIdNotesByNoteIdItemsResponseObject, error) {

}

// 資格情報群取得
// (GET /users/{byUserId}/qualifications)
func (s *strictServerImpl) GetUsersByUserIdQualifications(ctx context.Context, request GetUsersByUserIdQualificationsRequestObject) (GetUsersByUserIdQualificationsResponseObject, error) {

}

// 資格情報群最新化
// (PUT /users/{byUserId}/qualifications)
func (s *strictServerImpl) PutUsersByUserIdQualifications(ctx context.Context, request PutUsersByUserIdQualificationsRequestObject) (PutUsersByUserIdQualificationsResponseObject, error) {

}
