package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/Balaji01-4D/ecoware-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func generateExpiredAccessToken(userID uint) string {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(-1 * time.Minute).Unix(), // expired
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return signed
}

func insertValidRefreshToken(userID uint, refreshToken string) {
	initializer.DB.Create(&models.RefreshSession{
		UserID:    userID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	})
}

func TestRefreshTokenEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	// Test user
	userID := uint(1)
	expiredAccess := generateExpiredAccessToken(userID)
	refreshToken := "test-refresh-token-123"

	// Insert refresh session
	insertValidRefreshToken(userID, refreshToken)

	// Prepare request
	req, _ := http.NewRequest("GET", "/auth/refresh", nil)
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: expiredAccess,
	})
	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken,
	})

	// Send request
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Validate
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "token refreshed")
}


func TestLoginUser(t *testing.T) {
	router := routes.SetupRouter()

	email := fmt.Sprintf("login_test_%d@example.com", time.Now().UnixNano())

	// Register
	registerBody := map[string]interface{}{
		"name":     "Login User",
		"email":    email,
		"password": "testpass123",
		"role":     "user",
	}
	regJSON, _ := json.Marshal(registerBody)
	req1, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(regJSON))
	req1.Header.Set("Content-Type", "application/json")
	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)
	assert.Equal(t, http.StatusOK, resp1.Code, resp1.Body.String())

	// Login
	loginBody := map[string]interface{}{
		"email":    email,
		"password": "testpass123",
	}
	loginJSON, _ := json.Marshal(loginBody)
	req2, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(loginJSON))
	req2.Header.Set("Content-Type", "application/json")
	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)

	assert.Equal(t, http.StatusOK, resp2.Code, resp2.Body.String())
	assert.NotEmpty(t, resp2.Result().Cookies())
}

func TestLogoutEndpoint(t *testing.T) {
	router := routes.SetupRouter()

	req, _ := http.NewRequest("GET", "/auth/logout", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTE4MTkyODksInN1YiI6MTF9.kF3RlTuBuOpPnMefGpbrnMuxECKU_eCdwgnlo-gX0TM"})
	req.AddCookie(&http.Cookie{Name: "refresh_token", Value: "ee02e0ed-2fe2-4f42-bb38-8364a97a000e"})

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "logged out")
}

func TestGetMeEndpoint(t *testing.T) {
	router := routes.SetupRouter()

	req, _ := http.NewRequest("GET", "/auth/me", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: os.Getenv("TEST_TOKEN")})

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "email")
	assert.Contains(t, resp.Body.String(), "role")
}

func TestUpdatePasswordEndpoint(t *testing.T) {
	router := routes.SetupRouter()

	body := `{
		"currentPassword": "l@123",
		"newPassword": "t@123"
	}`
	req, _ := http.NewRequest("PUT", "/user/password", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: os.Getenv("TEST_TOKEN")})

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "successfully updated")
}

func TestUpdateProfileEndpoint(t *testing.T) {
	router := routes.SetupRouter()

	body := `{
		"name": "teo doss",
		"email": "teo@example.com"
	}`
	req, _ := http.NewRequest("PUT", "/user/profile", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: os.Getenv("TEST_TOKEN")})

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Contains(t, resp.Body.String(), "successfully updated")
}

