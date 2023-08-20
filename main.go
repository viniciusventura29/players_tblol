package main

import (
	"players_tblol/db"
	"players_tblol/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	routes.AppRouter(app, client)

	// createPost, err := client.Post.CreateOne(db.Post.Name.Set("Vinicius")).Exec(context.Background())

	// if err != nil {
	// 	return
	// }

	app.Run("localhost:3000")
}
