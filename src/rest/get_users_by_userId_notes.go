package rest

import (
	"context"

	"github.com/sky0621/cv-admin/src/ent/user"
	"github.com/sky0621/cv-admin/src/ent/usernote"
)

// GetUsersByUserIdNotes 注釈群取得
// (GET /users/{byUserId}/notes)
func (s *strictServerImpl) GetUsersByUserIdNotes(ctx context.Context, request GetUsersByUserIdNotesRequestObject) (GetUsersByUserIdNotesResponseObject, error) {
	// TODO: 直接 user_id 指定できるはず。。
	userNotes, err := s.dbClient.UserNote.Query().Where(usernote.HasUserWith(user.ID(request.ByUserId))).WithNoteItems().All(ctx)
	if err != nil {
		return nil, err
	}

	return GetUsersByUserIdNotes200JSONResponse{ToSwaggerUserNotes(userNotes)}, nil
}
