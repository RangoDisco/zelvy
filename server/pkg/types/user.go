package types

type UserRequest struct {
	Username    string `json:"username"`
	DiscordID   string `json:"discordID"`
	PaypalEmail string `json:"paypalEmail"`
}

type Winner struct {
	DiscordID string `json:"discordID"`
}

type WinnerViewModel struct {
	Username string `json:"username"`
	Wins     int    `json:"wins"`
}
