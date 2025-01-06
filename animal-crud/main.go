package main

import (
	"animal-crud/db"
	"animal-crud/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	db.Connect()

	// Crear un nuevo router
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Permitir solo tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Definir las rutas
	r.GET("/animals", routes.GetAnimals)
	r.POST("/animals", routes.CreateAnimal)
	r.PUT("/animals/:id", routes.UpdateAnimal)
	r.DELETE("/animals/:id", routes.DeleteAnimal)
	r.GET("/cat-images", routes.GetCatImages)

	// Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
