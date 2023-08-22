package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"players_tblol/db"
	"strings"

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

type User struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Coins     float32 `json:"coins"`
	Level     int     `json:"level"`
	Image     string  `json:"image"`
	Adm       bool    `json:"adm"`
	Nickname  string  `json:"nickname"`
	CreatedAt string  `json:"createdAt"`
	UpdateAt  string  `json:"updateAt"`
	PlayersId string  `json:"playersId"`
}

type Team struct {
	Name string `json:"name"`
}

func AppRouter(Router *gin.Engine, client *db.PrismaClient) *gin.RouterGroup {

	ctx := context.Background()

	v1 := Router.Group("/")
	{
		v1.GET("/allPlayers", func(c *gin.Context) {

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

		v1.POST("/addPlayer", func(c *gin.Context) {
			var playerid PlayerId

			if err := c.BindJSON(&playerid); err != nil {
				c.String(http.StatusBadRequest, "Bad request saas")
				return
			}

			postBody, err := json.Marshal(playerid)

			if err != nil {
				c.String(http.StatusBadRequest, "Bad N")
				return
			}

			fmt.Println(bytes.NewBuffer(postBody))

			player, err := client.Player.FindUnique(db.Player.ID.Equals(playerid.PlayerId)).Exec(ctx)

			if err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			resp, err := http.Post("http://localhost:4000/addPlayer", "application/json", bytes.NewBuffer(postBody))

			fmt.Printf("resp.Close: %v\n", resp.Close)

			fmt.Printf("resp.body: %v\n", bytes.NewBuffer(postBody))

			c.JSON(resp.StatusCode, player)

		})

		v1.GET("/getMyPlayers", func(c *gin.Context) {
			resp, err := http.Get("http://localhost:4000/me")

			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)

			var user User

			json.Unmarshal(body, &user)

			if user.Id == "" {
				c.String(http.StatusBadRequest, "User do not exists!")
				return
			}

			playersIds := strings.Split(user.PlayersId, ",")

			players, err := client.Player.FindMany(db.Player.ID.In(playersIds)).Exec(ctx)

			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			fmt.Println(players)

			c.JSON(http.StatusOK, players)
		})

		v1.POST("/newTeam", func(c *gin.Context) {
			var team Team

			if err := c.BindJSON(&team); err != nil {
				c.String(http.StatusBadRequest, err.Error())
				return
			}

			newTeam, err := client.Team.CreateOne(db.Team.Name.Set(team.Name)).Exec(ctx)

			if err != nil {
				return
			}

			c.JSON(http.StatusOK, newTeam)

		})
	}

	return v1

}
