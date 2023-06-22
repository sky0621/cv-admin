package rest

import "context"

// PutUsersByUserIdNotesByNoteId 【未実装】注釈更新
// (PUT /users/{byUserId}/notes/{byNoteId})
func (s *strictServerImpl) PutUsersByUserIdNotesByNoteId(ctx context.Context, request PutUsersByUserIdNotesByNoteIdRequestObject) (PutUsersByUserIdNotesByNoteIdResponseObject, error) {
	// FIXME:
	return PutUsersByUserIdNotesByNoteId200JSONResponse{}, nil
}
