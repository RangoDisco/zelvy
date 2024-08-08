# Bot
![](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

Source code for the bot itself. See [here](../README.md) for more general information.

## Commands

### `/get summary {date}`

*admin only*

> Returns a summary of the day's transactions.

### `/set offday`

*admin only*

> Disable goals for the goals selected.

### `/set paypal {email}`

*everyone*

> Create User in the database with their paypal email in order facilitate transactions in case they win. <br>
> If the user already exist, it simply updates the email.
