package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	. "golang-dev-logic-challenge-devphilip/controllers"
	_ "golang-dev-logic-challenge-devphilip/docs"
	. "golang-dev-logic-challenge-devphilip/model"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		var contracts []OptionsContract

		if err := c.ShouldBindJSON(&contracts); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
			return
		}

		// Your code here
		analysisResponse := AnalysisHandler(contracts)
		c.JSON(http.StatusOK, analysisResponse)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
