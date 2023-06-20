package rest

import "context"

// PutUsersByUserIdNotesByNoteIdItems 【未実装】注釈内要素群最新化
// (PUT /users/{byUserId}/notes/{byNoteId}/items)
func (s *strictServerImpl) PutUsersByUserIdNotesByNoteIdItems(ctx context.Context, request PutUsersByUserIdNotesByNoteIdItemsRequestObject) (PutUsersByUserIdNotesByNoteIdItemsResponseObject, error) {
	return PutUsersByUserIdNotesByNoteIdItems200JSONResponse{}, nil
}
