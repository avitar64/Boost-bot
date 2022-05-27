package commands

import "github.com/avitar64/Boost-bot/discord"

type exportedCommands map[string]discord.SlashCommandHandlerType

func ExportCommands() exportedCommands {
	ec := make(exportedCommands)

	ec["ping"] = Ping
	ec["8ball"] = Ball8

	return ec
}