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

	// createTeam, err := client.Team.CreateOne(db.Team.Name.Set("Loud")).Exec(context.Background())

	// log.Printf("format", createTeam)

	// if err != nil {
	// 	return
	// }

	app.Run("0.0.0.0:3000")
}
