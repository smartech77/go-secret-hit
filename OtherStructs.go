package main

type election struct {
	totalvotes          int
	ja                  int
	nein                int
	failedelections     int
	specialelection     bool
	lastnormalpresident string
}

type Player struct {
	password      string
	name          string
	party         string
	isHitler      bool
	isAlive       bool
	hasVoted      bool
	votedFor      string
	VotedFor      string
	Vetoing       bool
	isTermLimited bool
	intel         []string
	policies      []string
}

type request struct {
	playerpassword string
	gamepassword   string
	name           string
	action         string
	target         string
	policyindex    int
	category       string
}
