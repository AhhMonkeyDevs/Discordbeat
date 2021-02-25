package beater

import (
	"encoding/json"
	"fmt"
	"github.com/AhhMonkeyDevs/discordgo-lite"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"time"

	"github.com/AhhMonkeyDevs/discordbeat/config"
)

// discordbeat configuration.
type discordbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of discordbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &discordbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts discordbeat.
func (bt *discordbeat) Run(b *beat.Beat) error {
	logp.Info("discordbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	//guilds := bt.config.AllowedGuilds

	c, _ := discordgo.ConnectToGateway(bt.config.Token, 512, func(eventName string, eventData json.RawMessage) {
		switch eventName{
		case "MESSAGE_CREATE":
			var msg discordgo.Message
			json.Unmarshal(eventData, &msg)

			eventFields := messageToFields(&msg)
			eventFields.Put("event", "MESSAGE_CREATE")

			bt.SendEvent(eventFields)
		case "MESSAGE_UPDATE":
			var msg discordgo.Message
			json.Unmarshal(eventData, &msg)

			var restMessage = make(chan []byte)
			discordgo.NewRestRequest().
				Token(bt.config.Token).
				Route("channels").
				Channel(msg.ChannelID).
				Route("messages").
				Id(msg.Id).
				Callback(restMessage).
				Enqueue()
			json.Unmarshal(<- restMessage, &msg)

			eventFields := messageToFields(&msg)
			eventFields.Put("event", "MESSAGE_UPDATE")

			bt.SendEvent(eventFields)
		case "MESSAGE_DELETE":
			var e discordgo.MessageDeleteEvent
			json.Unmarshal(eventData, &e)

			event := common.MapStr{
				"guild_id": e.GuildID,
				"channel_id": e.ChannelID,
				"id": e.Id,
				"event": "MESSAGE_DELETE",
			}
			bt.SendEvent(event)

		case "MESSAGE_DELETE_BULK":
			var e discordgo.MessageDeleteBulkEvent
			json.Unmarshal(eventData, &e)

			for _, v := range e.Ids{
				event := common.MapStr{
					"guild_id": e.GuildID,
					"channel_id": e.ChannelID,
					"id": v,
					"event": "MESSAGE_DELETE",
				}
				bt.SendEvent(event)
			}
		}
	})

	<- bt.done

	c.Close()

	return nil

}

// Stop stops discordbeat.
func (bt *discordbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *discordbeat) SendEvent(fields common.MapStr){
	event := beat.Event{
		Timestamp: time.Now(),
		Fields: fields,
	}

	bt.client.Publish(event)
}

func messageToFields(msg *discordgo.Message) common.MapStr {

	mf := GetMessageFormatter(*msg)

	return common.MapStr{
		"id":               msg.Id,
		"channel_id":       msg.ChannelID,
		"guild_id":         msg.GuildID,
		"author_id":        mf.getAuthorID(),
		"author_type": 		mf.getAuthorType(),
		"type": 			msg.Type,
		"user_mentions": 	mf.getUserMentions(),
		"role_mentions": 	msg.RoleMentions,
		"channel_mentions":	mf.getChannelMentions(),
		"content":          mf.getContent(),
		"tts":              msg.Tts,
		"mention_everyone": msg.MentionsEveryone,
		"link_hostnames":	mf.extractLinks(),
		"has": 				mf.getHasArray(),
		"attachment_filenames": mf.getAttachmentFilenames(),
		"attachment_extensions": mf.getAttachmentExtensions(),
		"referenced_message":	mf.getMessageReference(),
		"created_timestamp": GetTimestamp(msg.Timestamp),
		"edited_timestamp": GetTimestamp(msg.EditedTimestamp),
	}
}
