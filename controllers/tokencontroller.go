package controllers

import (
	auth "jwttuts/controllers/Auth"
	"jwttuts/dbmodels"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var req TokenRequest
	var user dbmodels.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// check if email exists and pass is corrct

	err := dbmodels.Users(qm.Where("email =?", req.Email)).Bind(c, boil.GetContextDB(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	credentialserror := user.CheckPassword(req.Password)
	if credentialserror != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		c.Abort()
		return
	}

	tokenstring, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenstring})

}
