package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetToken returns successfully if refresh_token is non-empty, and an auth error if empty
func GetToken(c *gin.Context) {
	switch token := c.PostForm("refresh_token"); token {
	case "connection1":
		log.Println("Auth token received with connection")
		file, err := ioutil.ReadFile("data/auth_connection1.json")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
			})
			return
		}
		c.Data(http.StatusOK, "application/json", file)
	case "connection2":
		log.Println("Auth token received with connection")
		file, err := ioutil.ReadFile("data/auth_connection2.json")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
			})
			return
		}
		c.Data(http.StatusOK, "application/json", file)
	case "":
		log.Println("Empty token received. Responding with auth error.")
		file, err := ioutil.ReadFile("data/auth_error.json")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
			})
			return
		}
		c.Data(http.StatusBadRequest, "application/json", file)
	default:
		log.Println("Auth token received. Responding with auth success.")
		file, err := ioutil.ReadFile("data/auth_success.json")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
			})
			return
		}
		c.Data(http.StatusOK, "application/json", file)
	}
}
