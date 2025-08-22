module github.com/rangodisco/zelvy/bot

go 1.24.1

require (
	github.com/bwmarrin/discordgo v0.28.1
	github.com/go-resty/resty/v2 v2.16.5
	github.com/joho/godotenv v1.5.1
	github.com/rangodisco/zelvy v0.0.0
	google.golang.org/grpc v1.74.2
)

replace github.com/rangodisco/zelvy => ../

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/jonboulle/clockwork v0.5.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/protobuf v1.36.7 // indirect
)

require (
	github.com/go-co-op/gocron/v2 v2.16.1
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
)
