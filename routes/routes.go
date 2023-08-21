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

			player, err := client.Player.FindMany().Exec(ctx)

			if err != nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": player,
			})
		})

		v1.POST("/addPlayer", func(c *gin.Context) {

			playerInfo := c.Request.Body

			player := client.Player.CreateOne(
				db.Player.Nickname.Set(playerInfo.nickname),
				db.Player.Image.Set(""),
				db.Player.Country.Set(playerInfo.country),
				db.Player.Points.set(playerInfo.points),
				db.Player.LastPoints.Set("70"),
				db.Player.Name.Set(playerInfo.name),
				db.Player.LastName.Set(playerInfo.lastName),
				db.Player.Lane.Set(playerInfo.lane),
				db.Player.LaneImg.Set(""),
				db.Player.Team.Link(nil),
			)

		})
	}

	return v1

}
