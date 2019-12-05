package tests

import (
	"encoding/json"
	"errors"
	"github.com/brianvoe/gofakeit"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"hottub/handler"
	"hottub/types"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	echo     *echo.Echo
	handler  *handler.Handler
	testUser types.User
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `fake:"{person.email}"`
	Password string `fake:"{person.password}"`
}

func (s *Suite) SetupSuite() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&types.User{})

	s.handler = &handler.Handler{DB: db}
	s.handler.DB.LogMode(true)
	s.echo = echo.New()

	gofakeit.Seed(time.Now().UnixNano())

	s.testUser = types.User{
		Username: gofakeit.Username(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, false, 1),
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(s.testUser.Password), bcrypt.DefaultCost)

	s.handler.DB.Create(types.User{
		Username: s.testUser.Username,
		Email:    s.testUser.Email,
		Password: string(hash),
	})
}

func (s *Suite) TestEcho() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	// Router
	s.Assert().NotNil(s.echo.Router())

	// DefaultHTTPErrorHandler
	s.echo.DefaultHTTPErrorHandler(errors.New("error"), c)
	assert.Equal(s.T(), http.StatusInternalServerError, rec.Code)
}

func (s *Suite) TestCreateUserNil() {
	req := httptest.NewRequest(http.MethodGet, "/register", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	_ = s.handler.Register(c)
	assert.Equal(s.T(), http.StatusOK, rec.Code)
}

func (s *Suite) TestCreateUser() {
	gofakeit.Seed(time.Now().UnixNano())
	bytes, _ := json.Marshal(RegisterRequest{
		Username: gofakeit.Username(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, false, 1),
	})

	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(bytes)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	if assert.NoError(s.T(), s.handler.Register(c)) {
		var jsonRequest types.User
		var jsonResponse types.User

		_ = json.Unmarshal(bytes, &jsonRequest)
		_ = json.Unmarshal([]byte(rec.Body.String()), &jsonResponse)

		// Test returned user over HTTP
		assert.Equal(s.T(), http.StatusCreated, rec.Code)
		assert.Equal(s.T(), jsonRequest.Username, jsonResponse.Username)
		assert.Equal(s.T(), jsonRequest.Email, jsonResponse.Email)

		// test db
		var user types.User
		s.handler.DB.First(&user, jsonResponse.ID)
		assert.Equal(s.T(), jsonResponse.Username, user.Username)
		assert.Equal(s.T(), jsonResponse.Email, user.Email)
		assert.NoError(s.T(), bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonRequest.Password)))
	}
}

func (s *Suite) TestLoginUser() {
	bytes, _ := json.Marshal(&s.testUser)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(bytes)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	if assert.NoError(s.T(), s.handler.Login(c)) {
		var jsonResponse types.User

		_ = json.Unmarshal([]byte(rec.Body.String()), &jsonResponse)
		assert.NotNil(s.T(), jsonResponse.Token)
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
