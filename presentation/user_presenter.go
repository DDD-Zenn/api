package presentation

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/DDD-Zenn/api/external/serviceIF"
	"github.com/DDD-Zenn/api/application/user"
)

type UserPresenter struct{
	userUsecase serviceIF.User
}

func NewUserPresenter(userUsecase serviceIF.User) *UserPresenter {
	return &UserPresenter{userUsecase: userUsecase}
}

func (h *UserPresenter) Create(c *gin.Context) {
	var req user.UserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUsecase.Create(req.UID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *UserPresenter) FindByUID(c *gin.Context) {
	uid := c.Param("uid")

	user, err := h.userUsecase.FindByUID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserPresenter) Update(c *gin.Context) {
	var req user.UserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUsecase.Update(req.UID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *UserPresenter) Delete(c *gin.Context) {
	uid := c.Param("uid")

	err := h.userUsecase.Delete(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
