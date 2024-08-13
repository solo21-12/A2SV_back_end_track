package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
)

type SignupController struct {
	SignupUsecase domain.SignUpUseCase
	Env           *bootstrap.Env
}

func (s SignupController) SignUp(ctx *gin.Context) {
	var user domain.UserCreateRequest

	secret, _ := s.SignupUsecase.GetJwtSecret()

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	encryptedPassword, paErr := s.SignupUsecase.EncryptPassword(user.Password)

	if paErr != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting the password"})
		return
	}

	user.Password = encryptedPassword

	newUser, nErr := s.SignupUsecase.CreateUser(ctx, user)

	if nErr != nil {
		ctx.IndentedJSON(nErr.Code, gin.H{"error": nErr.Message})
		return
	}

	token, tErr := s.SignupUsecase.CreateAccessToken(newUser, secret)

	if tErr != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error generating the token" + tErr.Error()})
		return
	}

	response := domain.SignUpResponse{
		User:        newUser,
		AccessToken: token,
	}

	ctx.IndentedJSON(http.StatusCreated, response)

}
