package routes

import (
	"context"
	"net/http"
	"players_tblol/db"

	"github.com/gin-gonic/gin"
)

func AppRouter(Router *gin.Engine, client *db.PrismaClient) *gin.RouterGroup {

	ctx := context.Background()

	v1 := Router.Group("/")
	{
		v1.GET("/", func(c *gin.Context) {

			posts, err := client.Post.FindMany().Exec(ctx)

			if err != nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": posts,
			})
		})
	}

	return v1

}
