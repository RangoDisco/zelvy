package helpers

import "github.com/bwmarrin/discordgo"

/**
 * Not used for now but could be useful for later features
 */
func GetAvatarByUserID(session *discordgo.Session, userID string) (string, error) {
	user, err := session.User(userID)
	if err != nil {
		return "", err
	}

	avatarURL := user.AvatarURL("")
	if avatarURL == "" {
		// If the user has no custom avatar, use the default avatar URL
		avatarURL = user.AvatarURL("0")
	}

	return avatarURL, nil
}
