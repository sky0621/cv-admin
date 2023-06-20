package rest

import "context"

// DeleteUsersByUserIdNotesByNoteId 【未実装】注釈削除
// (DELETE /users/{byUserId}/notes/{byNoteId})
func (s *strictServerImpl) DeleteUsersByUserIdNotesByNoteId(ctx context.Context, request DeleteUsersByUserIdNotesByNoteIdRequestObject) (DeleteUsersByUserIdNotesByNoteIdResponseObject, error) {
	// FIXME:
	return DeleteUsersByUserIdNotesByNoteId204Response{}, nil
}
