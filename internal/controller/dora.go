package controller

import (
	"dora/internal/repository"
	"dora/internal/types/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DoraController struct {
	repository *repository.DoraRepository
}

func NewDoraController() *DoraController {
	doraRepository, err := repository.NewDoraRepository()
	if err != nil {
		return nil
	}
	return &DoraController{repository: doraRepository}
}

func (cc DoraController) Create(ctx *gin.Context) {
	var pullRequest github.PullRequestEvent
	if err := ctx.ShouldBindJSON(&pullRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := cc.repository.Create()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
