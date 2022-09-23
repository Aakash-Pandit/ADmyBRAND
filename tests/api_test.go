package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Aakash-Pandit/ADmyBRAND/models"
	"github.com/Aakash-Pandit/ADmyBRAND/routes"
	"github.com/Aakash-Pandit/ADmyBRAND/services"
	"github.com/Aakash-Pandit/ADmyBRAND/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBInstance() *gorm.DB {
	db, err := storage.NewConnection()
	if err != nil {
		log.Fatal("could not load database")
	}
	return db
}

func GetUserInstance() models.User {
	db := DBInstance()
	user := models.User{}
	db.Take(&user)
	return user
}

func CreateMockUser() models.User {
	db := DBInstance()
	user := models.User{
		Name:        "John Jake",
		DateOfBirth: "1998-04-04",
		Address:     "Virar",
		Description: "dev",
	}
	db.Create(&user)
	return user
}

func CreateMockPayload() string {
	payload := map[string]string{
		"name":          "Steve R",
		"date_of_birth": "1998-04-04",
		"address":       "Virar",
		"description":   "dev",
	}

	jsonByte, _ := json.Marshal(payload)
	return string(jsonByte[:])
}

func TestCreateUser(t *testing.T) {
	numberOfUser := 10
	for numberOfUser > 0 {
		CreateMockUser()
		numberOfUser--
	}
}

func TestHomeMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	route := "/api/v1/"
	r := gin.Default()
	r.GET(route, routes.ApiHome)

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetAllUsersMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	DBInstance()

	route := "/api/v1/users/"
	r := gin.Default()
	r.GET(route, services.GetUsersHandler)

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetUserMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	user := GetUserInstance()
	id := user.ID.String() + "/"

	route := "/api/v1/users/"
	r := gin.Default()
	r.GET(route+":id/", services.GetUserByIDHandler)

	req, err := http.NewRequest(http.MethodGet, route+id, nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	CreateMockUser()

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestCreateUserMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	route := "/api/v1/users/"
	r := gin.Default()
	r.POST(route, services.CreateUserHandler)

	req, err := http.NewRequest(http.MethodPost, route, strings.NewReader(CreateMockPayload()))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusCreated {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestUpdateUserMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	user := GetUserInstance()
	id := user.ID.String() + "/"

	route := "/api/v1/users/"
	r := gin.Default()
	r.PATCH(route+":id/", services.UpdateUserHandler)

	req, err := http.NewRequest(http.MethodPatch, route+id, strings.NewReader(CreateMockPayload()))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestDeleteUserMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	user := GetUserInstance()
	id := user.ID.String() + "/"

	route := "/api/v1/users/"
	r := gin.Default()
	r.DELETE(route+":id/", services.DeleteUserHandler)

	req, err := http.NewRequest(http.MethodDelete, route+id, nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusNoContent {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
