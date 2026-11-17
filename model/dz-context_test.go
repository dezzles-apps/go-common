package model

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestGetContextCreatesAndReusesContext(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/", nil)
	logger := zap.NewNop()
	c.Set("logger", logger)

	got := GetContext(c)
	if got.Logger != logger || got.Ctx == nil || *got.Ctx != c.Request.Context() {
		t.Fatal("new DAContext did not preserve the request context and logger")
	}
	userID := "user"
	got.UserId = &userID
	if again := GetContext(c); again != got || again.UserId != &userID {
		t.Fatal("GetContext did not reuse the existing DAContext")
	}
}

func TestGetContextReusesStoredContext(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	want := &DAContext{}
	c.Set(DAContextKey, want)

	if got := GetContext(c); got != want {
		t.Fatal("GetContext did not return the stored DAContext")
	}
}

func TestGetContextReplacesInvalidStoredValue(t *testing.T) {
	for _, value := range []any{"invalid", (*DAContext)(nil)} {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = httptest.NewRequest("GET", "/", nil)
		c.Set(DAContextKey, value)

		got := GetContext(c)
		if got == nil || got.Logger == nil || GetContext(c) != got {
			t.Fatal("GetContext did not create and store a valid DAContext")
		}
	}
}
