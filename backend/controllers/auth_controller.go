package controllers

import (
	"fmt"
	"log"
	"minigram-app-backend/models"
	"minigram-app-backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var listAuth = []models.User{}

func AuthRegister(ctx *gin.Context) {
	// create variable
	var user models.User
	var responseFailure models.ResponseFailure
	var count int64

	// get db
	db := repository.GetDb()

	// Binding data JSON & checking err
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Debug().Model(&user).Where("username = ?", user.Username).Count(&count).Error
	if err != nil {
		responseFailure.Message = fmt.Sprintf("error Username %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "gagal",
		})
		return
	}

	if count > 0 {
		responseFailure.Message = fmt.Sprintf("username %s has been registered", user.Username)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseFailure)
		return
	}

	// CEK EMAIL
	err = db.Debug().Model(&user).Where("email = ?", user.Email).Count(&count).Error
	// err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.Email).Scan(&exist)
	if err != nil {
		responseFailure.Message = fmt.Sprintf("error Email %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseFailure)
		return
	}

	if count > 0 {
		responseFailure.Message = fmt.Sprintf("email %s has been registered", user.Email)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseFailure)
		return
	}

	// _, err = user.BeforeCreate()
	// if err != nil {
	// 	responseFailure.Message = fmt.Sprintf("Something wrong : %s", err.Error())
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	// }

	// create user
	err = db.Debug().Create(&user).Error
	if err != nil {
		log.Println("EXEC", err)
		responseFailure.Message = fmt.Sprintf("Something wrong : %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User successfully registered",
		"data":    user,
	})

}

// func AuthLogin(ctx *gin.Context) {
// 	var user user.User

// }

func GetAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": listAuth,
	})
}
