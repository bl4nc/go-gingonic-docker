package main

import (
	"net/http"
	"os"
	"tecnicos_service/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := setupRouter()
	_ = r.Run(":" + os.Getenv("APP_PORT"))

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	tecnicosRepo := controllers.New()
	r.GET("/tecnico_id/:id", tecnicosRepo.GetTecnicoId)
	r.GET("/tecnico_login/:login", tecnicosRepo.GetTecnicoLogin)
	r.GET("/tecnicos", tecnicosRepo.GetTecnicos)

	return r

}
