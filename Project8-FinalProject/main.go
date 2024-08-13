package main

// I installed gin to connect to server and run it locally "go get -u github.com/gin-gonic/gin"
// Then installed sql to connect to db "go get github.com/matten/go-sqlite3"
import (
	"example.com/project/db"
	"example.com/project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
