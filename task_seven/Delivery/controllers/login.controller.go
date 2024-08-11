package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
)

type LoginController struct {
	LoginUseCase domain.LoginUseCase
	Env          *bootstrap.Env
}

func (l *LoginController) Login(ctx *gin.Context) {
	var user domain.LoginRequest
	secret, _ := l.LoginUseCase.GetJwtSecret()

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	existUser, err := l.LoginUseCase.GetUserEmail(ctx, user.Email)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	if validePassword := l.LoginUseCase.ValidatePassword(user.Password, existUser.Password); !validePassword {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email and password loop"})
		return
	}

	curUser := domain.UserDTO{
		ID:    existUser.ID,
		Email: existUser.Email,
		Role:  existUser.Role,
	}
	token, tErr := l.LoginUseCase.CreateAccessToken(curUser, secret)

	if tErr != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "error while generating token"})
		return
	}

	response := domain.LoginResponse{
		User:        curUser,
		AccessToken: token,
	}

	ctx.IndentedJSON(http.StatusOK, response)

}
