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
type CareerDescriptions = []string

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
	Url *Url `json:"url"`

	// スキルバージョン
	Version *SkillVersion `json:"version"`
}

// CareerSkillGroup defines model for CareerSkillGroup.
type CareerSkillGroup struct {
	// スキルグループラベル
	Label *SkillGroupLabel `json:"label"`

	// １キャリアのスキル群
	Skills *CareerSkills `json:"skills,omitempty"`
}

// １キャリアのスキル群
type CareerSkills = []CareerSkill

// １キャリアのタスク
type CareerTask struct {
	// タスク説明
	Description *TaskDescriptions `json:"description"`

	// タスク名
	Name *TaskName `json:"name,omitempty"`
}

// １キャリアのタスク群
type CareerTasks = []CareerTask

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
type TaskDescriptions = []string

// タスク名
type TaskName = string

// URL
type Url = string

// １ユーザーのアクティビティ群
type UserActivities = []UserActivity

// １ユーザーのアクティビティ
type UserActivity struct {
	// アクティビティアイコン
	Icon *ActivityIcon `json:"icon"`

	// アクティビティ名
	Name *ActivityName `json:"name,omitempty"`

	// URL
	Url *Url `json:"url"`
}

// １ユーザーの属性
type UserAttribute struct {
	// URL
	AvatarUrl *Url `json:"avatarUrl"`

	// 所属
	BelongTo *BelongTo `json:"belongTo"`

	// 生年月日
	Birthday *BirthDay `json:"birthday,omitempty"`

	// ユーザーを一意に識別するキー
	Id *UserId `json:"id,omitempty"`

	// 職業
	Job *Job `json:"job"`

	// ユーザーの本名
	Name *UserName `json:"name,omitempty"`

	// ユーザーのニックネーム
	Nickname *UserNickName `json:"nickname"`

	// PR
	Pr *PR `json:"pr"`
}

// １キャリアグループのキャリア
type UserCareer struct {
	// キャリア説明
	Description *CareerDescriptions `json:"description"`

	// キャリア期間年月From
	From *CareerPeriodFrom `json:"from,omitempty"`

	// キャリア名
	Name *CareerName `json:"name,omitempty"`

	// １キャリアのスキルグループ群
	SkillGroups *[]CareerSkillGroup `json:"skillGroups,omitempty"`

	// １キャリアのタスク群
	Tasks *CareerTasks `json:"tasks,omitempty"`

	// キャリア期間年月To
	To *CareerPeriodTo `json:"to"`
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

// １ユーザーのキャリアグループ群
type UserCareerGroups = []UserCareerGroup

// ユーザーを一意に識別するキー
type UserId = int

// ユーザーの本名
type UserName = string

// ユーザーのニックネーム
type UserNickName = string

// １ユーザーの注釈
type UserNote struct {
	// １ユーザーの注釈内の要素群
	Items *UserNoteItems `json:"items,omitempty"`

	// ラベル
	Label *NoteLabel `json:"label,omitempty"`

	// メモ
	Memo *NoteMemo `json:"memo"`
}

// １ユーザーの注釈内の１要素
type UserNoteItem struct {
	// 注釈内の１要素
	Text *NoteItemText `json:"text,omitempty"`
}

// １ユーザーの注釈内の要素群
type UserNoteItems = []UserNoteItem

// １ユーザーの注釈自身
type UserNoteOwn struct {
	// ラベル
	Label *NoteLabel `json:"label,omitempty"`

	// メモ
	Memo *NoteMemo `json:"memo"`
}

// １ユーザーの注釈群
type UserNotes = []UserNote

// １ユーザーの資格情報
type UserQualification struct {
	// 取得日
	GotDate *QualificationGotDate `json:"gotDate"`

	// メモ
	Memo *QualificationMemo `json:"memo"`

	// 資格名
	Name *QualificationName `json:"name,omitempty"`

	// 組織名
	Organization *QualificationOrg `json:"organization"`

	// URL
	Url *Url `json:"url"`
}

// １ユーザーの資格情報群
type UserQualifications = []UserQualification

// 生年月日の「年」
type Year = int

// １ユーザーのアクティビティ群
type N200OKUserActivities = UserActivities

// １ユーザーの属性
type N200OKUserAttribute = UserAttribute

// １ユーザーのキャリアグループ自身
type N200OKUserCareerGroupOwn = UserCareerGroupOwn

// １ユーザーのキャリアグループ群
type N200OKUserCareerGroups = UserCareerGroups

// １ユーザーのキャリアグループのキャリア群
type N200OKUserCareers = []UserCareer

// １ユーザーの注釈内の要素群
type N200OKUserNoteItems = UserNoteItems

// １ユーザーの注釈自身
type N200OKUserNoteOwn = UserNoteOwn

// １ユーザーの注釈群
type N200OKUserNotes = UserNotes

// １ユーザーの資格情報群
type N200OKUserQualifications = UserQualifications

// １ユーザーの属性
type N201CreatedUserAttribute = UserAttribute

// N400BadRequest defines model for 400-BadRequest.
type N400BadRequest = ClientError

// N404NotFound defines model for 404-NotFound.
type N404NotFound = NotFoundError

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody = UserAttribute

// PutUsersByUserIdActivitiesJSONBody defines parameters for PutUsersByUserIdActivities.
type PutUsersByUserIdActivitiesJSONBody = UserActivities

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
type PutUsersByUserIdNotesByNoteIdItemsJSONBody = UserNoteItems

// PutUsersByUserIdQualificationsJSONBody defines parameters for PutUsersByUserIdQualifications.
type PutUsersByUserIdQualificationsJSONBody = UserQualifications

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
	// キャリアグループ削除
	// (DELETE /users/{byUserId}/careergroups/{byCareerGroupId})
	DeleteUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error
	// キャリアグループ更新
	// (PUT /users/{byUserId}/careergroups/{byCareerGroupId})
	PutUsersByUserIdCareergroupsByCareerGroupId(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error
	// キャリアグループ内キャリア群最新化
	// (PUT /users/{byUserId}/careergroups/{byCareerGroupId}/careers)
	PutUsersByUserIdCareergroupsByCareerGroupIdCareers(ctx echo.Context, byUserId UserId, byCareerGroupId CareerGroupId) error
	// 注釈群取得
	// (GET /users/{byUserId}/notes)
	GetUsersByUserIdNotes(ctx echo.Context, byUserId UserId) error
	// 注釈新規登録
	// (POST /users/{byUserId}/notes)
	PostUsersByUserIdNotes(ctx echo.Context, byUserId UserId) error
	// 注釈削除
	// (DELETE /users/{byUserId}/notes/{byNoteId})
	DeleteUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId UserId, byNoteId NoteId) error
	// 注釈更新
	// (PUT /users/{byUserId}/notes/{byNoteId})
	PutUsersByUserIdNotesByNoteId(ctx echo.Context, byUserId UserId, byNoteId NoteId) error
	// 注釈内要素群最新化
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
	"nxw91P4oIxW8frr+vwAAAP//uHBJZcBSAAA=",
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
