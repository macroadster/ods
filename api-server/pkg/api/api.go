package api

import (
	"github.com/gin-gonic/gin"
	"ds/api-server/pkg/backend"
	"ds/api-server/pkg/common"
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
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusMovedPermanently, gin.H{})
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
	c.JSON(http.StatusOK, gin.H{})
	//} else {
	//	c.JSON(302, gin.H{})
	//}
}

// CreatePipeline godoc
// @Summary Create Pipeline
// @Schemes
// @Description	Create a pipeline with Airflow DAG
// @Tags pipeline
// @Accept  json
// @Produce json
// @Param id path string true "pipeline ID"
// @Success 200 {string} committed
// @Router   /pipeline/{id} [post]
func CreatePipeline(c *gin.Context) {
  id := c.Param("id")
	ctx := context.Background()
	go backend.CreateDAG(ctx, id)
	c.JSON(http.StatusOK, gin.H{})
}

// GetPipeline godoc
// @Summary Get Pipeline
// @Schemes
// @Description	View a pipeline with Airflow DAG
// @Tags pipeline
// @Accept  json
// @Produce json
// @Param id path string true "pipeline ID"
// @Success 200 {string} committed
// @Router   /pipeline/{id} [get]
func GetPipeline(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	go backend.GetDAG(ctx, id)
	c.JSON(http.StatusOK, gin.H{})
}

// DeletePipeline godoc
// @Summary Delete Pipeline
// @Schemes
// @Description	Delete a pipeline with Airflow DAG
// @Tags pipeline
// @Accept  json
// @Produce json
// @Param id path string true "pipeline ID"
// @Success 200 {string} committed
// @Router   /pipeline/{id} [delete]
func DeletePipeline(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	go backend.DeleteDAG(ctx, id)
	c.JSON(http.StatusOK, gin.H{})
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
