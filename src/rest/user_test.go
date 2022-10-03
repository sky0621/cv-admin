package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/ent/enttest"
)

var (
	testDBFilePath = "../../db/enttest"
	testDBSrcName  = fmt.Sprintf("file:%s?cache=shared&_fk=1", testDBFilePath)

	userJSON = `{"name":"ユーザー１","nickname":"sky0621","avatarUrl":"https://example.com","birthday":{"year":1900,"month":1,"day":1},"job":"職業","belongTo":"所属","pr":"PR"}`
)

func testExecRestWrapper(t *testing.T, method string, body io.Reader, fn func(server ServerInterface, echoCtx echo.Context) error) (*httptest.ResponseRecorder, error, func()) {
	e := echo.New()
	req := httptest.NewRequest(method, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	dbCli := enttest.Open(t, "sqlite3", testDBSrcName, []enttest.Option{enttest.WithOptions(ent.Log(t.Log))}...)
	s := NewRESTService(dbCli)

	err := fn(s, c)
	return rec, err, func() {
		dbCli.Close()
		if err := os.Remove(testDBFilePath); err != nil {
			fmt.Println(err)
		}
	}
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestPostUsers(t *testing.T) {
	resRecorder, err, dbCloseFunc := testExecRestWrapper(t, http.MethodPost, strings.NewReader(userJSON),
		func(server ServerInterface, echoCtx echo.Context) error {
			return server.PostUsers(echoCtx)
		})
	defer dbCloseFunc()
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, resRecorder.Code)
		// TODO: parse and check
		fmt.Println(resRecorder.Body)
	}
}

func TestGetUsersByUserIdAttribute(t *testing.T) {
	resRecorder, err, dbCloseFunc := testExecRestWrapper(t, http.MethodGet, nil,
		func(server ServerInterface, echoCtx echo.Context) error {
			return server.GetUsersByUserIdAttribute(echoCtx, 1)
		})
	defer dbCloseFunc()
	fmt.Println(resRecorder)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusNotFound, resRecorder.Code)
		// TODO: parse and check
		fmt.Println(resRecorder.Body)
	}
}
