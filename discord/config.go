package discord

import (
	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Token    string
	ClientID string
	Secret   string
}

type Context struct {
	Session *discordgo.Session
	Config  *Config
}

func (c *Config) Client() (*Context, error) {
	session, err := discordgo.New(c.Token)
	if err != nil {
		return nil, err
	}

	return &Context{Config: c, Session: session}, nil
}
