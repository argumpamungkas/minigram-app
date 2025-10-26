package auth_controller

import (
	"fmt"
	"log"
	"math/rand"
	"minigram-app-backend/config"
	"minigram-app-backend/helpers"
	"minigram-app-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func generateApiKey() (apiKey string) {
	// source akan selalu berubah2
	source := rand.NewSource(time.Now().UnixNano())

	// ambil random dari source
	random := rand.New(source)

	// min & max length
	minL := 18
	maxL := 32

	// data random diambil dari antara 18 -32
	randomNumber := random.Intn(maxL-minL+1) + minL

	// charset dari a-z A-Z 1-0
	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	// buat string sebanyak random number
	stringRandom := make([]byte, rand.Intn(randomNumber))
	for i := range stringRandom {
		stringRandom[i] = charset[rand.Intn(len(charset))]
	}

	// kembalikan
	apiKey = string(stringRandom)
	return
}

func AuthRegister(ctx *gin.Context) {
	// create variable
	var user models.User
	var responseMsg models.ResponseMessage
	var count int64

	// get db
	db := config.GetDb()

	// check bind yang dikirim client apakah appjson/form
	// gin sudah otomatis check by content type
	if err := ctx.ShouldBind(&user); err != nil {
		responseMsg.Message = fmt.Sprintf("Something Wrong %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseMsg)
		return
	}

	err := db.Debug().Model(&user).Where("username = ?", user.Username).Count(&count).Error
	if err != nil {
		responseMsg.Message = fmt.Sprintf("error Username %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseMsg)
		return
	}

	if count > 0 {
		responseMsg.Message = fmt.Sprintf("username %s has been registered", user.Username)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseMsg)
		return
	}

	// CEK EMAIL
	err = db.Debug().Model(&user).Where("email = ?", user.Email).Count(&count).Error
	// err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.Email).Scan(&exist)
	if err != nil {
		responseMsg.Message = fmt.Sprintf("error Email %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseMsg)
		return
	}

	if count > 0 {
		responseMsg.Message = fmt.Sprintf("email %s has been registered", user.Email)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseMsg)
		return
	}

	// _, err = user.BeforeCreate()
	// if err != nil {
	// 	responseMsg.Message = fmt.Sprintf("Something wrong : %s", err.Error())
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	// }

	// create user
	err = db.Debug().Create(&user).Error
	if err != nil {
		log.Println("EXEC", err)
		responseMsg.Message = fmt.Sprintf("Something wrong : %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseMsg)
		return
	}

	// return success json
	responseMsg.Message = fmt.Sprintf("User successfully registered")
	ctx.JSON(http.StatusCreated, responseMsg)

}

func AuthLogin(ctx *gin.Context) {
	var user models.User
	var userLogin models.UserLogin
	var responseMsg models.ResponseMessage

	// get db
	db := config.GetDb()

	// bind json dan cek apakah ada error ?
	if err := ctx.ShouldBind(&userLogin); err != nil {
		responseMsg.Message = fmt.Sprintf("Something wrong : %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseMsg)
		return
	}

	// cari ke table user berdasarkan username
	err := db.Debug().Where("username = ?", userLogin.Username).Model(&user).Take(&user).Error
	if err != nil {
		responseMsg.Message = fmt.Sprintf("User not exist")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseMsg)
		return
	}

	comparePwd := helpers.ComparePassword([]byte(user.Password), []byte(userLogin.Password))
	if !comparePwd {
		responseMsg.Message = fmt.Sprintf("Invalid username and password")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseMsg)
		return
	}

	token, err := helpers.GenerateToken(user.Username, user.ApiKey)
	if err != nil {
		responseMsg.Message = fmt.Sprintf("Something wrong %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseMsg)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"username":  user.Username,
	// 	"full_name": user.FullName,
	// 	"email":     user.Email,
	// 	"avatar":    user.Avatar,
	// 	"bio":       user.Bio,
	// })
}
