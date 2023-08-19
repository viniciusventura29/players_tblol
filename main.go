package main

import (
	"context"
	"net/http"
	"players_tblol/db"

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

	// createPost, err := client.Post.CreateOne(db.Post.Name.Set("Vinicius")).Exec(context.Background())

	// if err != nil {
	// 	return
	// }

	app.GET("/", func(c *gin.Context) {

		posts, err := client.Post.FindMany().Exec(context.Background())

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": posts,
		})
	})

	app.GET("/oi/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": ctx.Params.ByName("id"),
		})
	})

	app.Run("localhost:3000")
}
