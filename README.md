# Zelvy
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/rangodisco/zelvy)
![GitHub commits since latest release](https://img.shields.io/github/commits-since/rangodisco/zelvy/latest)
![GitHub last commit](https://img.shields.io/github/last-commit/rangodisco/zelvy)


# What it is

Discord bot created in order to keep me accountable for my daily fitness related tasks. \
Each day I can call the bot the check if I have completed my goals for the day. These goals are:

- 1h gym session
- 1h of cardio (can be soft, like walking)
- Less than 2100kcal consumed
- More than 1000kcal burned
- More than 2L of water drank

If one of these goals is not met, the bot will pick a winner from the registered users and I will this person 5€.\
Each goal is dynamic and can be changed at any time (depending on my objectives). \

# Technical details

This project is composed of 4 parts each containing a `README.md` file with more information:

- 🤖 [Bot](https://github.com/RangoDisco/zelby/tree/main/bot) written in Go with
  the [discordgo](https://github.com/bwmarrin/discordgo) library
- 🖥 [Server](https://github.com/RangoDisco/zelby/tree/main/server) written in Go
  with [Gin](https://github.com/gin-gonic/gin), [Gorm](https://github.com/go-gorm/gorm)
  and [PostgreSQL](https://www.postgresql.org/).
- 🌐 The server also includes a frontend written
  with [htmx](https://github.com/bigskysoftware/htmx), [TailwindCSS](https://github.com/tailwindlabs/tailwindcss)
  and [DaisyUI](https://github.com/saadeghi/daisyui)
- 📱 [Companion app](https://github.com/RangoDisco/zelby-companion) Written with Swift, for now on its own repository but
  will be moved here later

# Backstory

The name comes from JLeague1 football team [Machida Zelvia](https://www.zelvia.co.jp/)'s mascot Zelvy.

