// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package rest

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// アクティビティアイコン
type ActivityIcon = string

// アクティビティ名
type ActivityName = string

// 所属
type BelongTo = string

// 生年月日
type BirthDay struct {
	// 生年月日の「日」
	Day *Day `json:"day,omitempty"`

	// 生年月日の「月」
	Month *Month `json:"month,omitempty"`

	// 生年月日の「年」
	Year *Year `json:"year,omitempty"`
}

// キャリア説明
type CareerDescription = []string

// キャリアグループを一意に識別するID
type CareerGroupId = int

// キャリアグループラベル
type CareerGroupLabel = string

// キャリア名
type CareerName = string

// キャリア期間年月From
type CareerPeriodFrom struct {
	// 生年月日の「月」
	Month *Month `json:"month,omitempty"`

	// 生年月日の「年」
	Year *Year `json:"year,omitempty"`
}

// キャリア期間年月To
type CareerPeriodTo struct {
	// 生年月日の「月」
	Month *Month `json:"month,omitempty"`

	// 生年月日の「年」
	Year *Year `json:"year,omitempty"`
}

// １キャリアのスキル
type CareerSkill struct {
	// スキル名
	Name *SkillName `json:"name,omitempty"`

	// URL
	Url *Url `json:"url,omitempty"`

	// スキルバージョン
	Version *SkillVersion `json:"version,omitempty"`
}

// １キャリアのスキルグループ
type CareerSkillGroup struct {
	// スキルグループラベル
	Label  *SkillGroupLabel `json:"label,omitempty"`
	Skills *[]CareerSkill   `json:"skills,omitempty"`
}

// １キャリアのタスク
type CareerTask struct {
	// タスク説明
	Description *TaskDescription `json:"description,omitempty"`

	// タスク名
	Name *TaskName `json:"name,omitempty"`
}

// ClientError defines model for ClientError.
type ClientError struct {
	// エラーメッセージ
	Message *ErrorMessage `json:"message,omitempty"`
}

// 生年月日の「日」
type Day = int

// エラーメッセージ
type ErrorMessage = string

// 職業
type Job = string

// 生年月日の「月」
type Month = int

// NotFoundError defines model for NotFoundError.
type NotFoundError struct {
	// エラーメッセージ
	Message *ErrorMessage `json:"message,omitempty"`
}

// 注釈を一意に識別するID
type NoteId = int

// 注釈内の１要素
type NoteItemText = string

// ラベル
type NoteLabel = string

// メモ
type NoteMemo = string

// PR
type PR = string

// 取得日
type QualificationGotDate = openapi_types.Date

// メモ
type QualificationMemo = string

// 資格名
type QualificationName = string

// 組織名
type QualificationOrg = string

// スキルグループラベル
type SkillGroupLabel = string

// スキル名
type SkillName = string

// スキルバージョン
type SkillVersion = string

// タスク説明
type TaskDescription = []string

// タスク名
type TaskName = string

// URL
type Url = string

// １ユーザーのアクティビティ
type UserActivity struct {
	// アクティビティアイコン
	Icon *ActivityIcon `json:"icon,omitempty"`

	// アクティビティ名
	Name *ActivityName `json:"name,omitempty"`

	// URL
	Url *Url `json:"url,omitempty"`
}

// １ユーザーの属性
type UserAttribute struct {
	// URL
	AvatarUrl *Url `json:"avatarUrl,omitempty"`

	// 所属
	BelongTo *BelongTo `json:"belongTo,omitempty"`

	// 生年月日
	Birthday *BirthDay `json:"birthday,omitempty"`

	// ユーザーを一意に識別するキー
	Id *UserId `json:"id,omitempty"`

	// 職業
	Job *Job `json:"job,omitempty"`

	// ユーザーの本名
	Name *UserName `json:"name,omitempty"`

	// ユーザーのニックネーム
	Nickname *UserNickName `json:"nickname,omitempty"`

	// PR
	Pr *PR `json:"pr,omitempty"`
}

// １キャリアグループのキャリア
type UserCareer struct {
	// キャリア説明
	Description *CareerDescription `json:"description,omitempty"`

	// キャリア期間年月From
	From *CareerPeriodFrom `json:"from,omitempty"`

	// キャリア名
	Name *CareerName `json:"name,omitempty"`

	// １キャリアのスキルグループ群
	SkillGroups *[]CareerSkillGroup `json:"skillGroups,omitempty"`

	// １キャリアのタスク群
	Tasks *[]CareerTask `json:"tasks,omitempty"`

	// キャリア期間年月To
	To *CareerPeriodTo `json:"to,omitempty"`
}

// １ユーザーのキャリアグループ
type UserCareerGroup struct {
	// １ユーザーのキャリアグループのキャリア群
	Careers *[]UserCareer `json:"careers,omitempty"`

	// キャリアグループラベル
	Label *CareerGroupLabel `json:"label,omitempty"`
}

// １ユーザーのキャリアグループ自身
type UserCareerGroupOwn struct {
	// キャリアグループラベル
	Label *CareerGroupLabel `json:"label,omitempty"`
}

// ユーザーを一意に識別するキー
type UserId = int

// ユーザーの本名
type UserName = string

// ユーザーのニックネーム
type UserNickName = string

// １ユーザーの注釈
type UserNote struct {
	Items *[]UserNoteItem `json:"items,omitempty"`

	// ラベル
	Label *NoteLabel `json:"label,omitempty"`

	// メモ
	Memo *NoteMemo `json:"memo,omitempty"`
}

// １ユーザーの注釈内の１要素
type UserNoteItem struct {
	// 注釈内の１要素
	Text *NoteItemText `json:"text,omitempty"`
}

// １ユーザーの注釈自身
type UserNoteOwn struct {
	// ラベル
	Label *NoteLabel `json:"label,omitempty"`

	// メモ
	Memo *NoteMemo `json:"memo,omitempty"`
}

// １ユーザーの資格情報
type UserQualification struct {
	// 取得日
	GotDate *QualificationGotDate `json:"gotDate,omitempty"`

	// メモ
	Memo *QualificationMemo `json:"memo,omitempty"`

	// 資格名
	Name *QualificationName `json:"name,omitempty"`

	// 組織名
	Organization *QualificationOrg `json:"organization,omitempty"`

	// URL
	Url *Url `json:"url,omitempty"`
}

// 生年月日の「年」
type Year = int

// N200OKUserActivities defines model for 200-OK-UserActivities.
type N200OKUserActivities = []UserActivity

// １ユーザーの属性
type N200OKUserAttribute = UserAttribute

// １ユーザーのキャリアグループ自身
type N200OKUserCareerGroupOwn = UserCareerGroupOwn

// N200OKUserCareerGroups defines model for 200-OK-UserCareerGroups.
type N200OKUserCareerGroups = []UserCareerGroup

// N200OKUserCareers defines model for 200-OK-UserCareers.
type N200OKUserCareers = []UserCareer

// N200OKUserNoteItems defines model for 200-OK-UserNoteItems.
type N200OKUserNoteItems = []UserNoteItem

// １ユーザーの注釈自身
type N200OKUserNoteOwn = UserNoteOwn

// N200OKUserNotes defines model for 200-OK-UserNotes.
type N200OKUserNotes = []UserNote

// N200OKUserQualifications defines model for 200-OK-UserQualifications.
type N200OKUserQualifications = []UserQualification

// １ユーザーの属性
type N201CreatedUserAttribute = UserAttribute

// N400BadRequest defines model for 400-BadRequest.
type N400BadRequest = ClientError

// N404NotFound defines model for 404-NotFound.
type N404NotFound = NotFoundError

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody = UserAttribute

// PutUsersByUserIdActivitiesJSONBody defines parameters for PutUsersByUserIdActivities.
type PutUsersByUserIdActivitiesJSONBody = []UserActivity

// PutUsersByUserIdAttributeJSONBody defines parameters for PutUsersByUserIdAttribute.
type PutUsersByUserIdAttributeJSONBody = UserAttribute

// PostUsersByUserIdCareergroupsJSONBody defines parameters for PostUsersByUserIdCareergroups.
type PostUsersByUserIdCareergroupsJSONBody = UserCareerGroup

// PutUsersByUserIdCareergroupsByCareerGroupIdJSONBody defines parameters for PutUsersByUserIdCareergroupsByCareerGroupId.
type PutUsersByUserIdCareergroupsByCareerGroupIdJSONBody = UserCareerGroupOwn

// PutUsersByUserIdCareergroupsByCareerGroupIdCareersJSONBody defines parameters for PutUsersByUserIdCareergroupsByCareerGroupIdCareers.
type PutUsersByUserIdCareergroupsByCareerGroupIdCareersJSONBody = []UserCareer

// PostUsersByUserIdNotesJSONBody defines parameters for PostUsersByUserIdNotes.
type PostUsersByUserIdNotesJSONBody = UserNote

// PutUsersByUserIdNotesByNoteIdJSONBody defines parameters for PutUsersByUserIdNotesByNoteId.
type PutUsersByUserIdNotesByNoteIdJSONBody = UserNoteOwn

// PutUsersByUserIdNotesByNoteIdItemsJSONBody defines parameters for PutUsersByUserIdNotesByNoteIdItems.
type PutUsersByUserIdNotesByNoteIdItemsJSONBody = []UserNoteItem

// PutUsersByUserIdQualificationsJSONBody defines parameters for PutUsersByUserIdQualifications.
type PutUsersByUserIdQualificationsJSONBody = []UserQualification

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = PostUsersJSONBody

// PutUsersByUserIdActivitiesJSONRequestBody defines body for PutUsersByUserIdActivities for application/json ContentType.
type PutUsersByUserIdActivitiesJSONRequestBody = PutUsersByUserIdActivitiesJSONBody

// PutUsersByUserIdAttributeJSONRequestBody defines body for PutUsersByUserIdAttribute for application/json ContentType.
type PutUsersByUserIdAttributeJSONRequestBody = PutUsersByUserIdAttributeJSONBody

// PostUsersByUserIdCareergroupsJSONRequestBody defines body for PostUsersByUserIdCareergroups for application/json ContentType.
type PostUsersByUserIdCareergroupsJSONRequestBody = PostUsersByUserIdCareergroupsJSONBody

// PutUsersByUserIdCareergroupsByCareerGroupIdJSONRequestBody defines body for PutUsersByUserIdCareergroupsByCareerGroupId for application/json ContentType.
type PutUsersByUserIdCareergroupsByCareerGroupIdJSONRequestBody = PutUsersByUserIdCareergroupsByCareerGroupIdJSONBody

// PutUsersByUserIdCareergroupsByCareerGroupIdCareersJSONRequestBody defines body for PutUsersByUserIdCareergroupsByCareerGroupIdCareers for application/json ContentType.
type PutUsersByUserIdCareergroupsByCareerGroupIdCareersJSONRequestBody = PutUsersByUserIdCareergroupsByCareerGroupIdCareersJSONBody

// PostUsersByUserIdNotesJSONRequestBody defines body for PostUsersByUserIdNotes for application/json ContentType.
type PostUsersByUserIdNotesJSONRequestBody = PostUsersByUserIdNotesJSONBody

// PutUsersByUserIdNotesByNoteIdJSONRequestBody defines body for PutUsersByUserIdNotesByNoteId for application/json ContentType.
type PutUsersByUserIdNotesByNoteIdJSONRequestBody = PutUsersByUserIdNotesByNoteIdJSONBody

// PutUsersByUserIdNotesByNoteIdItemsJSONRequestBody defines body for PutUsersByUserIdNotesByNoteIdItems for application/json ContentType.
type PutUsersByUserIdNotesByNoteIdItemsJSONRequestBody = PutUsersByUserIdNotesByNoteIdItemsJSONBody

// PutUsersByUserIdQualificationsJSONRequestBody defines body for PutUsersByUserIdQualifications for application/json ContentType.
type PutUsersByUserIdQualificationsJSONRequestBody = PutUsersByUserIdQualificationsJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ユーザーアカウント登録
	// (POST /users)
	PostUsers(ctx echo.Context) error
	// 指定ユーザーアカウント削除
	// (DELETE /users/{byUserId})
	DeleteUsersByUserId(ctx echo.Context, byUserId UserId) error
	// アクティビティ群取得
	// (GET /users/{byUserId}/activities)
	GetUsersByUserIdActivities(ctx echo.Context, byUserId UserId) error
	// アクティビティ群最新化
	// (PUT /users/{byUserId}/activities)
	PutUsersByUserIdActivities(ctx echo.Context, byUserId UserId) error
	// 属性取得
	// (GET /users/{byUserId}/attribute)
	GetUsersByUserIdAttribute(ctx echo.Context, byUserId UserId) error
	// 属性更新
	// (PUT /users/{byUserId}/attribute)
	PutUsersByUserIdAttribute(ctx echo.Context, byUserId UserId) error
	// キャリアグループ群取得
	// (GET /users/{byUserId}/careergroups)
	GetUsersByUserIdCareergroups(ctx echo.Context, byUserId UserId) error
	// キャリアグループ新規登録
	// (POST /users/{byUserId}/careergroups)
	PostUsersByUserIdCareergroups(ctx echo.Context, byUserId UserId) error
	// 【未実装】キャリアグループ削除
	// (DELETE /users/{byUserId}/careergroups/{byCareerGroupId})
	DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error
	// 【未実装】キャリアグループ更新
	// (PUT /users/{byUserId}/careergroups/{byCareerGroupId})
	PutUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error
	// 【未実装】キャリアグループ内キャリア群最新化
	// (PUT /users/{byUserId}/careergroups/{byCareerGroupId}/careers)
	PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error
	// 注釈群取得
	// (GET /users/{byUserId}/notes)
	GetUsersByUserIdNotes(ctx echo.Context, byUserId UserId) error
	// 注釈新規登録
	// (POST /users/{byUserId}/notes)
	PostUsersByUserIdNotes(ctx echo.Context, byUserId UserId) error
	// 【未実装】注釈削除
	// (DELETE /users/{byUserId}/notes/{byNoteId})
	DeleteUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId UserId, byNoteId NoteId) error
	// 【未実装】注釈更新
	// (PUT /users/{byUserId}/notes/{byNoteId})
	PutUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId UserId, byNoteId NoteId) error
	// 【未実装】注釈内要素群最新化
	// (PUT /users/{byUserId}/notes/{byNoteId}/items)
	PutUsersByUserIdNotesByNoteIdItems(ctx echo.Context, byUserId UserId, byNoteId NoteId) error
	// 資格情報群取得
	// (GET /users/{byUserId}/qualifications)
	GetUsersByUserIdQualifications(ctx echo.Context, byUserId UserId) error
	// 資格情報群最新化
	// (PUT /users/{byUserId}/qualifications)
	PutUsersByUserIdQualifications(ctx echo.Context, byUserId UserId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// DeleteUsersByUserId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersByUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUsersByUserId(ctx, byUserId)
	return err
}

// GetUsersByUserIdActivities converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersByUserIdActivities(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersByUserIdActivities(ctx, byUserId)
	return err
}

// PutUsersByUserIdActivities converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersByUserIdActivities(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUsersByUserIdActivities(ctx, byUserId)
	return err
}

// GetUsersByUserIdAttribute converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersByUserIdAttribute(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersByUserIdAttribute(ctx, byUserId)
	return err
}

// PutUsersByUserIdAttribute converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersByUserIdAttribute(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUsersByUserIdAttribute(ctx, byUserId)
	return err
}

// GetUsersByUserIdCareergroups converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersByUserIdCareergroups(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersByUserIdCareergroups(ctx, byUserId)
	return err
}

// PostUsersByUserIdCareergroups converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersByUserIdCareergroups(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUsersByUserIdCareergroups(ctx, byUserId)
	return err
}

// DeleteUsersByUserIdCareergroupsByCareerGroupId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// ------------- Path parameter "byCareerGroupId" -------------
	var byCareerGroupId CareerGroupId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byCareerGroupId", runtime.ParamLocationPath, ctx.Param("byCareerGroupId"), &byCareerGroupId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byCareerGroupId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx, byUserId, byCareerGroupId)
	return err
}

// PutUsersByUserIdCareergroupsByCareerGroupId converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// ------------- Path parameter "byCareerGroupId" -------------
	var byCareerGroupId CareerGroupId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byCareerGroupId", runtime.ParamLocationPath, ctx.Param("byCareerGroupId"), &byCareerGroupId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byCareerGroupId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUsersByUserIdCareergroupsByCareerGroupId(ctx, byUserId, byCareerGroupId)
	return err
}

// PutUsersByUserIdCareergroupsByCareerGroupIdCareers converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// ------------- Path parameter "byCareerGroupId" -------------
	var byCareerGroupId CareerGroupId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byCareerGroupId", runtime.ParamLocationPath, ctx.Param("byCareerGroupId"), &byCareerGroupId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byCareerGroupId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx, byUserId, byCareerGroupId)
	return err
}

// GetUsersByUserIdNotes converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersByUserIdNotes(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersByUserIdNotes(ctx, byUserId)
	return err
}

// PostUsersByUserIdNotes converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersByUserIdNotes(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUsersByUserIdNotes(ctx, byUserId)
	return err
}

// DeleteUsersByUserIdNotesByNoteId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersByUserIdNotesByNoteId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// ------------- Path parameter "byNoteId" -------------
	var byNoteId NoteId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byNoteId", runtime.ParamLocationPath, ctx.Param("byNoteId"), &byNoteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byNoteId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUsersByUserIdNotesByNoteId(ctx, byUserId, byNoteId)
	return err
}

// PutUsersByUserIdNotesByNoteId converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersByUserIdNotesByNoteId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// ------------- Path parameter "byNoteId" -------------
	var byNoteId NoteId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byNoteId", runtime.ParamLocationPath, ctx.Param("byNoteId"), &byNoteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byNoteId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUsersByUserIdNotesByNoteId(ctx, byUserId, byNoteId)
	return err
}

// PutUsersByUserIdNotesByNoteIdItems converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersByUserIdNotesByNoteIdItems(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// ------------- Path parameter "byNoteId" -------------
	var byNoteId NoteId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byNoteId", runtime.ParamLocationPath, ctx.Param("byNoteId"), &byNoteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byNoteId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUsersByUserIdNotesByNoteIdItems(ctx, byUserId, byNoteId)
	return err
}

// GetUsersByUserIdQualifications converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersByUserIdQualifications(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersByUserIdQualifications(ctx, byUserId)
	return err
}

// PutUsersByUserIdQualifications converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersByUserIdQualifications(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "byUserId" -------------
	var byUserId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "byUserId", runtime.ParamLocationPath, ctx.Param("byUserId"), &byUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter byUserId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutUsersByUserIdQualifications(ctx, byUserId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.DELETE(baseURL+"/users/:byUserId", wrapper.DeleteUsersByUserId)
	router.GET(baseURL+"/users/:byUserId/activities", wrapper.GetUsersByUserIdActivities)
	router.PUT(baseURL+"/users/:byUserId/activities", wrapper.PutUsersByUserIdActivities)
	router.GET(baseURL+"/users/:byUserId/attribute", wrapper.GetUsersByUserIdAttribute)
	router.PUT(baseURL+"/users/:byUserId/attribute", wrapper.PutUsersByUserIdAttribute)
	router.GET(baseURL+"/users/:byUserId/careergroups", wrapper.GetUsersByUserIdCareergroups)
	router.POST(baseURL+"/users/:byUserId/careergroups", wrapper.PostUsersByUserIdCareergroups)
	router.DELETE(baseURL+"/users/:byUserId/careergroups/:byCareerGroupId", wrapper.DeleteUsersByUserIdCareergroupsByCareerGroupId)
	router.PUT(baseURL+"/users/:byUserId/careergroups/:byCareerGroupId", wrapper.PutUsersByUserIdCareergroupsByCareerGroupId)
	router.PUT(baseURL+"/users/:byUserId/careergroups/:byCareerGroupId/careers", wrapper.PutUsersByUserIdCareergroupsByCareerGroupIdCareers)
	router.GET(baseURL+"/users/:byUserId/notes", wrapper.GetUsersByUserIdNotes)
	router.POST(baseURL+"/users/:byUserId/notes", wrapper.PostUsersByUserIdNotes)
	router.DELETE(baseURL+"/users/:byUserId/notes/:byNoteId", wrapper.DeleteUsersByUserIdNotesByNoteId)
	router.PUT(baseURL+"/users/:byUserId/notes/:byNoteId", wrapper.PutUsersByUserIdNotesByNoteId)
	router.PUT(baseURL+"/users/:byUserId/notes/:byNoteId/items", wrapper.PutUsersByUserIdNotesByNoteIdItems)
	router.GET(baseURL+"/users/:byUserId/qualifications", wrapper.GetUsersByUserIdQualifications)
	router.PUT(baseURL+"/users/:byUserId/qualifications", wrapper.PutUsersByUserIdQualifications)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
