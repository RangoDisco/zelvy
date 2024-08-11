# Zelvy

![GitHub commit activity](https://img.shields.io/github/commit-activity/w/rangodisco/zelvy)
![GitHub commits since latest release](https://img.shields.io/github/commits-since/rangodisco/zelvy/latest)
![GitHub last commit](https://img.shields.io/github/last-commit/rangodisco/zelvy)

# What it is

Discord bot created in order to keep me accountable for my daily fitness related tasks. \
Each day the bot checks if each goals has been completed. These goals are:

- `1h` gym session
- `1h` of cardio (can be soft, like walking)
- Less than `2100kcal` consumed
- More than `1000kcal` burned
- More than `2L` of water drank

If one of these goals is not met, the bot will mentionned the pre-determined winner and I will owe this person 5€.\
Each value is a variable and therefore can be changed at any time (depending on my objectives).

# Technical details

This project is composed of 4 parts each containing a `README.md` file with more information:

- 🤖 [Bot](https://github.com/RangoDisco/zelby/tree/main/bot) written in Go with
  the [discordgo](https://github.com/bwmarrin/discordgo) library
- 🖥 [Server](https://github.com/RangoDisco/zelby/tree/main/server) written in Go
  with [Gin](https://github.com/gin-gonic/gin), [Gorm](https://github.com/go-gorm/gorm)
  and [PostgreSQL](https://www.postgresql.org/).
- 🌐 The server also includes a frontend written
  with [Templ](https://github.com/a-h/templ), [htmx](https://github.com/bigskysoftware/htmx), [TailwindCSS](https://github.com/tailwindlabs/tailwindcss)
  and [DaisyUI](https://github.com/saadeghi/daisyui)
- 📱 [Companion app](https://github.com/RangoDisco/zelby-companion) Written with Swift, for now on its own repository but
  will be moved here later

# Backstory

The name comes from the J1 League football team [Machida Zelvia](https://www.zelvia.co.jp/)'s mascot Zelvy.
