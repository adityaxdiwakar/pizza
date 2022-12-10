package main

type tomlConfig struct {
	DiscordConfig discordCredentials
}

type discordCredentials struct {
	Token  string
	Prefix string
	Env    string
}
