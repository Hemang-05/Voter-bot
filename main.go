package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

var eligibility bool

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5054387922661-5057752233106-HkYa5wAAMuoLiW908SvTSttY")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A051S5YFZ52-5060204313892-cb8c2b6e939f715da86e6669a20b6c504aa214191c1ea3c974526fcbf86dafbe")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "Voter ID eligiblity",
		Handler: func(botctx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
			year := req.Param("year")

			yob, err := strconv.Atoi(year)

			if err != nil {
				println("error")
			}

			if 2024-yob >= 18 {
				eligibility = true
			} else {
				eligibility = false
			}

			if eligibility {
				res.Reply("You are Eligible to vote :)")
			} else {
				res.Reply("You are not Eligible to vote :)")
			}

		},
	})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
