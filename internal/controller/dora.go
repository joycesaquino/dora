package controller

import (
	"dora/internal/service"
	"dora/internal/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type DoraController struct {
	service *service.GithubService
}

func NewDoraController() *DoraController {

	githubService := service.NewGithubService()
	if githubService == nil {
		return nil
	}

	return &DoraController{service: githubService}
}

func (cc DoraController) Create(ctx *gin.Context) {
	var pullRequest types.PullRequestEvent
	if err := ctx.ShouldBindJSON(&pullRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(pullRequest)

	response, err := cc.service.Create(pullRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
