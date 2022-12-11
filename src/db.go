package main

type database interface {
	getRegisteredGuilds() map[string]string
}

type testConnection struct{}

func (t testConnection) getRegisteredGuilds() map[string]string {
	return map[string]string{
		"1051215097265143919": "gte.dev",
		"1051214879488495738": "gte.dev.a",
		"1051214960711188520": "gte.dev.b",
	}
}
