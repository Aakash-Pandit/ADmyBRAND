package services

import (
	"net/http"

	"github.com/Aakash-Pandit/ADmyBRAND/models"
	"github.com/Aakash-Pandit/ADmyBRAND/storage"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users := &[]models.User{}

	db := storage.GetDatabase()
	err := db.Find(users).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, users)
}

func GetUserByID(context *gin.Context) {
	id, _ := context.Params.Get("id")
	user := &models.User{}

	db := storage.GetDatabase()
	err := db.Where("id = ?", id).First(user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func CreateUser(context *gin.Context) {

	db := storage.GetDatabase()
	user := &models.User{}

	err := context.BindJSON(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	errors := models.ValidateUserStruct(*user)

	if errors != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": errors,
		})
		return
	}

	err = db.Create(user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func UpdateUser(context *gin.Context) {
	id, _ := context.Params.Get("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": "id cannot be empty",
		})
		return
	}

	user := &models.User{}

	err := context.BindJSON(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	errors := models.ValidateUserStruct(*user)

	if errors != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": errors,
		})
		return
	}

	db := storage.GetDatabase()

	err = db.Where("id = ?", id).Updates(user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	_ = db.Where("id = ?", id).First(user).Error
	context.JSON(http.StatusOK, user)
}

func DeleteUser(context *gin.Context) {
	user := &models.User{}
	id, _ := context.Params.Get("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": "id cannot be empty",
		})
		return
	}

	db := storage.GetDatabase()
	err := db.Where("id = ?", id).First(user).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"detail": err.Error(),
		})
		return
	}

	db.Where("id = ?", id).Delete(user)

	context.JSON(http.StatusNoContent, gin.H{})
}
