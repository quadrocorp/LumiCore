package db

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
