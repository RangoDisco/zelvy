package factories

import (
	"time"

	pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"
)

func CreateAddUserRequest() *pb_usr.AddUserRequest {
	return &pb_usr.AddUserRequest{
		Username:    "User123",
		PaypalEmail: "email@email.com",
		DiscordId:   "123",
	}
}

func CreateGetWinnersRequest() *pb_usr.GetWinnersRequest {
	return &pb_usr.GetWinnersRequest{
		StartDate: time.Now().Format("2006-01-02"),
		EndDate:   time.Now().Add(48 * time.Hour).Format("2006-01-02"),
		Limit:     2,
	}
}
