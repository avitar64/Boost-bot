package discord

import (
	"log"

	dg "github.com/bwmarrin/discordgo"
)

const (
	InstaMessage      = dg.InteractionResponseChannelMessageWithSource
	DefferSendMessage = dg.InteractionResponseDeferredChannelMessageWithSource
)

func Error(err error) {
	log.Printf("!! Error: %v\n", err)
}

func InteractionRespond(s *dg.Session, i *dg.Interaction, t dg.InteractionResponseType, d *dg.InteractionResponseData) error {
	return s.InteractionRespond(
		i,
		&dg.InteractionResponse{
			Type: t,
			Data: d,
		},
	)
}

func InteractionEdit(s *dg.Session, i *dg.Interaction, newresp *dg.WebhookEdit) error {
	_, err := s.InteractionResponseEdit(i, newresp)
	return err
}

func GetInteractionUser(i *dg.Interaction) *dg.User {
	u, m := i.User, i.Member
	if u != nil {
		return u
	}
	return m.User
}
