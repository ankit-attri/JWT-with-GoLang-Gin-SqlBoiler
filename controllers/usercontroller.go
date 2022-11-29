package controllers

import (
	"jwttuts/dbmodels"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func RegisterUser(c *gin.Context) {
	var user dbmodels.User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	err = user.HashPassword(user.Pass)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err = user.InsertG(c, boil.Infer()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"userid": user.ID, "email": user.Email, "username": user.Username})

}
