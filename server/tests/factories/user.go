package factories

import pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"

func CreateAddUserRequest() *pb_usr.AddUserRequest {
	return &pb_usr.AddUserRequest{
		Username:    "User123",
		PaypalEmail: "email@email.com",
		DiscordId:   "123",
	}
}
