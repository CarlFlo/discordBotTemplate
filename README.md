<h1 align="center">
  Discord Bot Template
</h1>

<p align="center">
  <a href="#about">About</a>
  •
  <a href="#features">Features</a>
  •
  <a href="#commands">Commands</a>
  •
  <a href="#setup">Setup</a>
  •
  <a href="#todo">Todo</a>
</p>

## About
The project aims to create an easy to use project template for creating Discord bots in the [GO](https://golang.org/) programming language.

## Features

- Works on Windows, Linux and Mac.
    - Dockerfile included
- Easy to customize and setup
    - Strightforward process for adding new commands and custom functionality
- Powerful, flexable and detailed log messages with [Malm](https://github.com/CarlFlo/malm)
- Dynamic message handling
    - Channel binding
    - Toggle support for direct messages 
- Built in bot link invite creator
- User and admin commands
    - Permission check
- Help is generated automatically
    - Commands alphabetically sorted into categories for readability
    - Supports help syntax for individual commands
    - Admin command seperate and won't show for users
- Statistics
    - Number of servers bot is active on

## Commands

These are the commands the are currently included in the bot

### Admin
* reload - Reloads the configuration file, updating the settings
* debug - Currently outputs the operating system the bot is running on
* presence - Outputs how many servers the bot is running on and info about them

### Users
* ping - Pong!
* botinvite - Creates an invite link to the bot
* echo - Echoes back whatever message was sent with the command
* help - Prints a list with all the commands and how to use them


## Setup

Clone the repo

```
git clone https://github.com/CarlFlo/discordBotTemplate.git
```

Install all the requirements

```
go mod download
```


### Configuration

Running the bot will create a `config.json` file in the directory where it was run.

Insert into the configuration file your bot token into the **Token** slot, discord user ID in the **OwnerID** slot and the bots discord ID in the **ClientID** slot (Used for the bot invite).

### Running

You're able to build and run the bot with the included `makefile`.

It is also possible to build or run it yourself with the `go run main.go` and `go build main.go` command, respectively.

## Todo

