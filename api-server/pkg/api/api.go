package api

import (
	"github.com/gin-gonic/gin"
	"ds/collector/pkg/backend"
	"ds/collector/pkg/common"
	"context"
	"net/http"
)

// Login godoc
// @Summary API Authentication callback
// @Schemes
// @Description	login
// @Tags auth
// @Accept  json
// @Produce json
// @Success 200 {string} Authenticated
// @Router   /auth [post]
func Login(c *gin.Context) {
	userID, _ := c.Get("UserUID")
	if userID != nil {
		c.JSON(200, gin.H{})
	} else {
		c.JSON(302, gin.H{})
	}
}

// PostData godoc
// @Summary Write Data API
// @Schemes
// @Description	Write Data
// @Tags data
// @Accept  json
// @Produce json
// @Param data body common.Event true "Message Event"
// @Success 200 {string} committed
// @Router   /data [post]
func PostData(c *gin.Context) {
	//userID, _ := c.Get("UserUID")
	//if userID != nil {
	event := common.Event{}
	if err := c.BindJSON(&event); err!=nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx := context.Background()
	go backend.Produce(ctx, event)
	c.JSON(200, gin.H{})
	//} else {
	//	c.JSON(302, gin.H{})
	//}
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{})
}
