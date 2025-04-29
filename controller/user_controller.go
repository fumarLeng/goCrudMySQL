package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"goCrudMySQL/model"
	"goCrudMySQL/service"
)

type UserController struct{ svc service.UserService }

func NewUserController(s service.UserService) *UserController {
	return &UserController{s}
}

func (ctl *UserController) Create(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctl.svc.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (ctl *UserController) List(c *gin.Context) {
	users, err := ctl.svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctl *UserController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := ctl.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctl *UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ctl.svc.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctl *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctl.svc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
