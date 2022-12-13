package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "crud/api/docs"
	"crud/api/handler"
)

func SetUpApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandlerV1(db)

	r.POST("/user", handlerV1.Create)
	r.GET("/user/:id", handlerV1.GetById)
	r.GET("/user", handlerV1.GetList)
	r.PUT("/user", handlerV1.Update)
	r.PATCH("/user", handlerV1.Patch)
	r.DELETE("/user/:id", handlerV1.Delete)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
