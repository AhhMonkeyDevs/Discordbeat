package beater

import (
	"fmt"
	"github.com/AhhMonkeyDevs/discordgo-lite"
	"regexp"
	"strings"
	"time"
)

type MessageFormatter struct {
	msg            discordgo.Message
	content        *[]string
	hasAttachments bool
	hasEmbeds      bool
	links          *[]string
}

func GetMessageFormatter(msg discordgo.Message) *MessageFormatter {
	mf := MessageFormatter{
		msg:            msg,
		hasAttachments: len(msg.Attachments) > 0,
		hasEmbeds:      len(msg.Embeds) > 0,
	}
	return &mf
}

func (f *MessageFormatter) getAuthorType() int {
	if f.msg.WebhookID == "" {
		return 2
	} else if f.msg.Author.Bot {
		return 1
	} else {
		return 0
	}
}

func (f *MessageFormatter) getChannelMentions() []string {
	content := f.getContent()
	return parseMentions(content, "<#([0-9]+)>")
}

func (f *MessageFormatter) getRoleMentions() []string {
	content := f.getContent()
	return parseMentions(content, "<@&([0-9]+)>")
}

func (f *MessageFormatter) getUserMentions() []string {
	content := f.getContent()
	return parseMentions(content, "<@!?([0-9]+)>")
}

func (f *MessageFormatter) getContent() []string {
	if f.content != nil {
		return *f.content
	}

	var content []string
	content = AppendNonZero(content, f.msg.Content)
	for _, e := range f.msg.Embeds {
		content = AppendNonZero(content, e.Title, e.Description)
		if e.Footer != nil {
			content = AppendNonZero(content, e.Footer.Text)
		}
		if e.Author != nil {
			content = AppendNonZero(content, e.Author.Name)
		}
		if e.Fields != nil {
			for _, f := range *e.Fields {
				content = AppendNonZero(content, f.Name, f.Value)
			}
		}
	}

	f.content = &content

	return content
}

func (f *MessageFormatter) extractLinks() []string {
	if f.links != nil {
		return *f.links
	}

	content := f.getContent()
	var links []string
	re := regexp.MustCompile(`(^|\s)https?:\/\/([A-Za-z0-9-.]+)`)
	for _, text := range content {
		matches := re.FindAllStringSubmatch(text, -1)
		for _, element := range matches {
			links = append(links, element[2])
		}
	}

	f.links = &links

	return links
}

func (f *MessageFormatter) getHasArray() []int {
	var has []int
	if f.hasAttachments {
		has = append(has, 2)
	}
	if f.hasEmbeds {
		has = append(has, 1)
	}
	if len(*f.links) > 0 {
		has = append(has, 0)
	}
	return has
}

func (f *MessageFormatter) getAttachmentFilenames() []string {
	var attachmentFilenames []string
	for _, a := range f.msg.Attachments {
		attachmentFilenames = append(attachmentFilenames, a.Filename)
	}
	return attachmentFilenames
}

func (f *MessageFormatter) getAttachmentExtensions() []string {
	var attachmentExtensions []string
	for _, a := range f.msg.Attachments {
		pieces := strings.Split(a.Filename, ".")
		attachmentExtensions = append(attachmentExtensions, pieces[len(pieces)-1])
	}
	return attachmentExtensions
}

func (f *MessageFormatter) getMessageReference() *string {
	if f.msg.MessageReference != nil {
		return &f.msg.MessageReference.MessageId
	}
	return nil
}

func (f *MessageFormatter) getAuthorID() *string {
	if f.msg.Author != nil {
		id := f.msg.Author.Id
		return &id
	}
	return nil
}

func parseMentions(content []string, format string) []string {
	var mentions []string
	re := regexp.MustCompile(format)
	for _, text := range content {
		matches := re.FindAllStringSubmatch(text, -1)
		for _, element := range matches {
			fmt.Println("Match")
			mentions = append(mentions, element[1])
		}
	}

	return mentions
}

func AppendNonZero(slice []string, elements ...string) []string {
	for _, v := range elements {
		if v != "" {
			slice = append(slice, v)
		}
	}
	return slice
}

func GetTimestamp(aTime *time.Time) *int64 {
	if aTime != nil {
		nanos := aTime.UnixNano() / int64(time.Millisecond)
		return &nanos
	}
	return nil
}
