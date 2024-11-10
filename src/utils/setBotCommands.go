package utils

import tgg "api-test/tgapi"

func SetBotCommands(bot *tgg.TelegramBot) error {
	commands := []tgg.BotCommand{
		{Command: "start", Description: "Start the bot"},
		{Command: "help", Description: "Get help"},
	}

	return bot.SetBotCommands(commands)
}
