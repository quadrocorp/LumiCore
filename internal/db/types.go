package db

// Author == quadrocorp
// Version == v1.0-alpha || Photon
// NOTE: Photon is the alpha version of this project, and implies an MVP.

// This file will contain types and models for the db package.

type user struct {
	id        int
	username  string
	firstName string
	lastName  string
	email     string
	password  string
	role      int
}

type containers struct {
	id          int
	botName     string
	botId       string
	containerId string
}
