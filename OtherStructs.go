package main

import (
	"fmt"
	"math/rand"
)

type Match struct {
	id                string
	password          string
	stage             string // this will dictate what can be done
	substage          int
	president         *Player
	chancellor        *Player
	hitler            *Player
	players           map[string]*Player
	playernames       []string
	policies          []string
	discardedpolicies []string
	libDecs           int
	fashDecs          int
	failedElections   int
	vetoEnabled       bool
	game_stage_enum   game_stage_enum
	fash_stage_enum   fash_stage_enum
	election          election
	waitingfor        string
	currentaction     string
	chancpolicies     []string
	presidentpolicies []string
}

type election struct {
	totalvotes      int
	ja              int
	nein            int
	failedelections int
}

type Player struct {
	//username string
	password       string
	name           string
	party          string
	isHitler       bool
	isAlive        bool
	hasVoted       bool
	isTermLimited  bool
	intel          []string
	canseepolicies bool
}

type request struct {
	playerpassword string
	gamepassword   string
	name           string
	action         string
	target         string
	index          int
}
type chancpolicies struct {
}
type presidentpolicies struct {
}

// func (match *Match) checkIfPowerEnabled() []string {}
func (match *Match) getCandidatesForChancellor() []string {

	var playernames2 []string

	for _, player := range match.players {
		if player.isTermLimited == false {
			playernames2 = append(playernames2, player.name)
		}
	}
	return playernames2
}
func (match *Match) collapsegovernment() {
	//also gotta do the random policy here

	if match.policies[0] == "fascist" {
		match.fashDecs++
	}
	if match.policies[0] == "liberal" {
		match.libDecs++
	}
	for _, player := range match.players {
		player.isTermLimited = false
	}
	match.policies = append(match.policies[:0], match.policies[1:]...)
	match.currentaction = "waiting for president to pick a chancellor"
	match.stage = match.game_stage_enum.election()
	match.substage = 1
}
func (match *Match) addplayer(player Player) {
}
func (match *Match) getPresidents(president string) (int, int, []string) {
	var playernames2 []string
	for i := 0; i < len(match.playernames); i++ {
		if match.players[match.playernames[i]].isAlive == true {
			playernames2 = append(playernames2, match.playernames[i])
		}

	}
	presidentIndex := 0
	nextpresidentIndex := 1
	for i := 0; i < len(playernames2); i++ {
		if playernames2[i] == president {
			presidentIndex = i
			break
		}
	}
	if presidentIndex == len(playernames2)-1 {
		nextpresidentIndex = 1
	} else {
		nextpresidentIndex = presidentIndex + 1
	}

	return presidentIndex, nextpresidentIndex, playernames2
}

func (match *Match) getlivingplayers() int {
	var aliveNumbers int = 0
	for _, v := range match.players {
		if v.isAlive == true {
			aliveNumbers++
		}
	}
	return aliveNumbers
}

func (match *Match) central_method(request request) {
	if match.contains_player(match.players, request.name) {
		if match.players[request.name].password == request.playerpassword && match.password == request.gamepassword {
			match.central_methodv2(request)
		}
	}

}

func (match *Match) central_methodv2(request request) {
	if match.stage == match.game_stage_enum.election() && request.action == "nominatechancellor" {
		if match.substage == 1 && match.president.password == request.playerpassword && match.password == request.gamepassword {
			match.chancellor = match.players[request.target]
			match.waitingfor = "all"
			match.currentaction = "voting on the new government"
			match.substage = 2
		}

		if match.substage == 2 && match.players[request.name].hasVoted == false && request.action == "vote" {
			if request.target == "ja" {
				match.election.ja++
			}
			if request.target == "nein" {
				match.election.nein++
			}
			match.players[request.name].hasVoted = true
			match.election.totalvotes++
			if match.getlivingplayers() == match.election.totalvotes {
				// if the election succeeds
				if match.election.ja > match.election.nein {
					match.waitingfor = match.chancellor.name
					match.currentaction = "chancellor is looking at top 3 cards of the deck"
					match.stage = match.game_stage_enum.policy()
					match.substage = 1

					for _, player := range match.players {
						player.isTermLimited = false
					}
					match.chancellor.isTermLimited = true
					if match.getlivingplayers() > 5 {
						match.president.isTermLimited = true
					}
					match.election.failedelections = 0
				}
				// if the election fails
				if match.election.ja < match.election.nein {
					match.failedElections++
					_, nextpresidentIndex, playernames2 := match.getPresidents(match.president.name)

					match.substage = 1
					match.stage = match.game_stage_enum.election()
					match.president = match.players[playernames2[nextpresidentIndex]]
					// if the government is collapsed
					if match.election.failedelections == 3 {
						match.collapsegovernment()
					}
					match.election.failedelections = 0
				}
				// cleanup happens here
				match.election.totalvotes = 0
				match.election.ja = 0
				match.election.nein = 0

				for _, player := range match.players {
					player.hasVoted = false
				}

			}

		}
	}
}

func (match *Match) calcFashNumbers() int {
	playersize := len(match.playernames)
	if playersize > 8 {
		return 3
	} else if playersize < 7 {
		return 1
	} else {
		return 2
	}
	return 2
}
func (match *Match) contains_player(m map[string]*Player, suspect string) bool {

	if _, ok := m[suspect]; ok {
		//	fmt.Printf("Value is : %s \n", v)
		return true
	} else {
		fmt.Println("Key not found")
		return false
	}

	return false
}

func (match *Match) randomize() {
	rand.Shuffle(len(match.playernames), func(i, j int) {
		match.playernames[i], match.playernames[j] = match.playernames[j], match.playernames[i]
	})
}
func (match *Match) randomizepolicies() {

	// rand.Seed()
	rand.Shuffle(len(match.policies), func(i, j int) {
		match.policies[i], match.policies[j] = match.policies[j], match.policies[i]
	})
}

func (match *Match) LaunchGame() {
	// this commented piece of code needs to be in the constructor
	//	match.players = make(map[string]*Player)
	for k := range match.players {
		match.playernames = append(match.playernames, k)
	}

	//	for i := 0; i < len(match.playernames); i++ {
	//		match.players[match.playernames[i]] = &Player{name: match.playernames[i], party: "liberal"}
	//	}

	match.randomize()

	fashnumbers := match.calcFashNumbers()
	match.hitler = match.players[match.playernames[0]]
	match.players[match.playernames[0]].isHitler = true
	match.players[match.playernames[0]].party = "fascist"

	for helpers := 1; helpers < fashnumbers+1; helpers++ {
		match.players[match.playernames[helpers]].party = "fascist"
	}
	match.randomize()
	match.president = match.players[match.playernames[0]]

	for i := 0; i < 6; i++ {
		match.policies = append(match.policies, "liberal")
	}
	for i := 0; i < 11; i++ {
		match.policies = append(match.policies, "fascist")
	}

	match.randomizepolicies()
	match.waitingfor = match.president.name
	match.currentaction = "waiting for president to pick a chancellor"
	match.stage = match.game_stage_enum.election()
	match.substage = 1
	var fascists []string
	for _, player := range match.players {

		if len(match.players) > 6 {
			if player.party == "fascist" && player.isHitler == false {
				fascists = append(fascists, player.name)
			}
		}

		if len(match.players) < 7 {
			if player.party == "fascist" {
				fascists = append(fascists, player.name)
			}
		}

	}
	var intel = " The fascists are : "
	var hitlerintel = "Hitler is actually : " + match.hitler.name
	for i := range fascists {
		intel = intel + fascists[i] + " , "
	}
	for i := range fascists {
		match.players[fascists[i]].intel = append(match.players[fascists[i]].intel, intel)
		match.players[fascists[i]].intel = append(match.players[fascists[i]].intel, hitlerintel)
	}

	//	fmt.Print(match.players)

	//	fascistsSize := match.calcFashNumbers()
	//	match.playernames[0]

}
