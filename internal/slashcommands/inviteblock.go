package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shinpuru/internal/services/database"
	"github.com/zekroTJA/shinpuru/internal/services/permissions"
	"github.com/zekroTJA/shinpuru/internal/util/static"
	"github.com/zekrotja/ken"
)

type Inviteblock struct {
	ken.EphemeralCommand
}

var (
	_ ken.SlashCommand        = (*Inviteblock)(nil)
	_ permissions.PermCommand = (*Inviteblock)(nil)
)

func (c *Inviteblock) Name() string {
	return "inviteblock"
}

func (c *Inviteblock) Description() string {
	return "Enable, disable or show state of invite blocking."
}

func (c *Inviteblock) Version() string {
	return "1.0.0"
}

func (c *Inviteblock) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *Inviteblock) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionBoolean,
			Name:        "enable",
			Description: "Set state to enabled or disabled.",
		},
	}
}

func (c *Inviteblock) Domain() string {
	return "sp.guild.mod.inviteblock"
}

func (c *Inviteblock) SubDomains() []permissions.SubPermission {
	return []permissions.SubPermission{
		{
			Term:        "send",
			Explicit:    true,
			Description: "Allows sending invites even if invite block is enabled",
		},
	}
}

func (c *Inviteblock) Run(ctx ken.Context) (err error) {
	if err = ctx.Defer(); err != nil {
		return
	}

	db, _ := ctx.Get(static.DiDatabase).(database.Database)

	stateV, ok := ctx.Options().GetByNameOptional("enable")

	if !ok {
		status, err := db.GetGuildInviteBlock(ctx.GetEvent().GuildID)
		if err != nil && !database.IsErrDatabaseNotFound(err) {
			return err
		}
		strStat := "disabled"
		color := static.ColorEmbedOrange
		if status != "" {
			strStat = "enabled"
			color = static.ColorEmbedGreen
		}

		return ctx.FollowUpEmbed(&discordgo.MessageEmbed{
			Description: fmt.Sprintf("Discord invite link blocking is currently **%s** on this guild.\n\n"+
				"*You can enable or disable this with the command `/inviteblock enable True/False`*.", strStat),
			Color: color,
		}).Send().Error
	}

	state := stateV.BoolValue()
	msg := "Discord invite links will **no more be blocked** on this guild now."
	stateStr := ""
	color := static.ColorEmbedOrange
	if state {
		msg = "Enabled invite link blocking."
		stateStr = "1"
		color = static.ColorEmbedGreen
	}

	err = db.SetGuildInviteBlock(ctx.GetEvent().GuildID, stateStr)
	if err != nil {
		return err
	}

	return ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Description: msg,
		Color:       color,
	}).Send().Error
}
