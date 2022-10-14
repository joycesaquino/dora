package controller

import (
	"dora/internal/repository"
	"dora/internal/types/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DoraController struct {
	repository repository.DoraRepository
}

func (cc DoraController) Create(c *gin.Context) {
	var pullRequest github.PullRequestEvent
	if err := c.ShouldBindJSON(&pullRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := cc.repository.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}
