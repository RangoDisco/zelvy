package types

type UserRequest struct {
	Username    string `json:"username"`
	DiscordID   string `json:"discordID"`
	PaypalEmail string `json:"paypalEmail"`
}
