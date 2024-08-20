package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context){
	user := c.MustGet("user").(string)
	c.JSON(http.StatusOK, gin.H{"user": user})
}