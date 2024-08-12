package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
	"github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/bootstrap"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PromoteController struct {
	PromoteUseCase domain.PromoteUseCase
	Env            *bootstrap.Env
}

func (p *PromoteController) PromoteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	nErr := p.PromoteUseCase.PromoteUser(objectId, ctx)

	if nErr != nil {
		ctx.IndentedJSON(nErr.Code, gin.H{"error": nErr.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User promoted successfully"})

}
