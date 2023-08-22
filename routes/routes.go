package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"players_tblol/db"

	"github.com/gin-gonic/gin"
)

type Player struct {
	Nickname string `json:"nickname"`
	Country  string `json:"country"`
	Points   string `json:"points"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Lane     string `json:"lane"`
	TeamId   string `json:"teamId"`
}

type PlayerId struct {
	PlayerId string `json:"playerid"`
}

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

		v1.POST("/newPlayer", func(c *gin.Context) {

			var playerInfo Player

			if err := c.BindJSON(&playerInfo); err != nil {
				c.String(http.StatusBadRequest, "Bad request in BindJSON")
				return
			}

			player, err := client.Player.CreateOne(
				db.Player.Nickname.Set(playerInfo.Nickname),
				db.Player.Image.Set(""),
				db.Player.Country.Set(playerInfo.Country),
				db.Player.Points.Set(playerInfo.Points),
				db.Player.LastPoints.Set("70"),
				db.Player.Name.Set(playerInfo.Name),
				db.Player.LastName.Set(playerInfo.LastName),
				db.Player.Lane.Set(playerInfo.Lane),
				db.Player.LaneImg.Set(""),
				db.Player.Team.Link(db.Team.ID.Equals(playerInfo.TeamId)),
			).Exec(ctx)

			if err != nil {
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"Player ": player,
			})

		})

		v1.POST("/addPlayer", func(ctx *gin.Context) {
			var playerid PlayerId

			if err := ctx.BindJSON(&playerid); err != nil {
				ctx.String(http.StatusBadRequest, "Bad request saas")
				return
			}

			postBody, err := json.Marshal(playerid)

			if err != nil {
				ctx.String(http.StatusBadRequest, "Bad N")
				return
			}

			fmt.Println(bytes.NewBuffer(postBody))

			resp, err := http.Post("http://localhost:4000/teste", "application/json", bytes.NewBuffer(postBody))

			fmt.Printf("resp.Close: %v\n", resp.Close)

		})
	}

	return v1

}
