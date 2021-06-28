package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var channel Command = &discordgo.ApplicationCommand{
	Name: "channel",
	Description: "Manage channel",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:			discordgo.ApplicationCommandOptionSubCommand,
			Name:			"create",
			Description:	"Create channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:			discordgo.ApplicationCommandOptionString,
					Name:			"name",
					Description:	"channel name",
					Required:		true,
				},
				{
					Type:			discordgo.ApplicationCommandOptionBoolean,
					Name:			"voice",
                    Description:	"create voice channel (default: `False`)",
					Required:		false,
				},
				{
					Type:			discordgo.ApplicationCommandOptionBoolean,
					Name:			"nsfw",
                    Description:	"create nsfw channel (default: `False`)",
					Required:		false,
				},
				{
					Type:			discordgo.ApplicationCommandOptionBoolean,
					Name:			"private",
                    Description:	"create private channel (default: `False`)",
					Required:		false,
				},
				{
					Type:			discordgo.ApplicationCommandOptionChannel,
					Name:			"category",
					Description:	"create into the category",
					Required:		false,
				},
			},
		},
		{
			Type:			discordgo.ApplicationCommandOptionSubCommand,
			Name:			"edit",
			Description:	"Edit channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:			discordgo.ApplicationCommandOptionString,
					Name:			"name",
					Description:	"channel name",
					Required:		false,
				},
				{
					Type:			discordgo.ApplicationCommandOptionString,
					Name:			"topic",
                    Description:	"edit channel topic",
					Required:		false,
				},
				{
					Type:			discordgo.ApplicationCommandOptionInteger,
					Name:			"bitrate",
                    Description:	"edit channel bitrate",
					Required:		false,
				},
				{
					Type:			discordgo.ApplicationCommandOptionBoolean,
					Name:			"nsfw",
                    Description:	"edit nsfw flag",
					Required:		false,
				},
				{
					Type:			discordgo.ApplicationCommandOptionInteger,
					Name:			"position",
                    Description:	"edit channel position (in category)",
					Required:		false,
				},
			},
		},
		{
			Type:			discordgo.ApplicationCommandOptionSubCommand,
			Name:			"archive",
			Description:	"Toggle archive/expand channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:			discordgo.ApplicationCommandOptionChannel,
					Name:			"target",
					Description:	"select channel (no selected, target is here)",
					Required:		false,
				},
			},
		},
		{
			Type:			discordgo.ApplicationCommandOptionSubCommand,
			Name:			"hidden",
			Description:	"Toggle visible/hidden channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:			discordgo.ApplicationCommandOptionChannel,
					Name:			"target",
					Description:	"select channel (no selected, target is here)",
					Required:		false,
				},
			},
		},
		{
			Type:			discordgo.ApplicationCommandOptionSubCommand,
			Name:			"destroy",
			Description:	"Destroy channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:			discordgo.ApplicationCommandOptionChannel,
					Name:			"target",
					Description:	"select channel to destroy",
					Required:		true,
				},
			},
		},
	},
}

func channelHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
}

func init() {
	addCommand(channel, channelHandler)
}
