package main

import (
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/CarlFlo/discordBotTemplate/bot"
	"github.com/CarlFlo/discordBotTemplate/config"
	"github.com/CarlFlo/discordBotTemplate/utils"
	"github.com/CarlFlo/malm"
)

// https://discordapp.com/oauth2/authorize?&client_id=643191140849549333&scope=bot&permissions=37211200

func init() {

	utils.Clear()
	rand.Seed(time.Now().UTC().UnixNano())

	malm.Debug("Running on %s", runtime.GOOS)

	if err := config.LoadConfiguration(); err != nil {
		malm.Fatal("Error loading configuration: %v", err)
	}

	malm.Debug("Version %s", config.CONFIG.Version)
}

func main() {

	session := bot.StartBot()

	time.Sleep(500 * time.Millisecond) // Added this sleep so the messages below will come last
	// Keeps bot from closing. Waits for CTRL-C
	malm.Info("Press CTRL-C to initiate shutdown")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	malm.Info("Shutting down!")

	// Run cleanup code here
	close(sc)
	session.Close() // Stops the discord bot
}
