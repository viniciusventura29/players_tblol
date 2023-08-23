# players_tblol

Tblol is a Web app that you could create your dream team with real professionals e-sports players and to compete against others teams!

## players
This part is responsible for all the logics that involves the player (the professional e-sport player)! This microservice communicates with auth_tblol

## endpoints
Follow below, all the endpoints that there is in this microservices. 

#### Get all players
"*/allPlayers"

#### Create a new player (professional e-sport player)
"*/newPlayer"

Params: 
   	- Nickname: string
	  - Country: string 
	  - Points: string 
	  - Name: string 
	  - LastName: string 
	  - Lane: string
	  - TeamId: string 

#### To assign a player to a user team
"*/addPlayer"

Params:
     - playerId: string

#### Get the players that the user have 
"*/getMyPlayers"

#### to add a player in your team
"*/newTeam"

Params:
     - name: string

#### Get all teams
"*/getAllTeams"

#### to remove a player in your team
"*/removePlayer"

Params:
     - playerId: string

