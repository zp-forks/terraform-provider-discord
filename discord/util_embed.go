package discord

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
)

type UnmappedEmbed struct {
	Title       string                             `json:"title,omitempty"`       // title of embed
	Description string                             `json:"description,omitempty"` // description of embed
	URL         string                             `json:"url,omitempty"`         // url of embed
	Timestamp   string                             `json:"timestamp,omitempty"`   // timestamp	timestamp of embed content
	Color       int                                `json:"color,omitempty"`       // color code of the embed
	Footer      []*discordgo.MessageEmbedFooter    `json:"footer,omitempty"`      // embed footer object	footer information
	Image       []*discordgo.MessageEmbedImage     `json:"image,omitempty"`       // embed image object	image information
	Thumbnail   []*discordgo.MessageEmbedThumbnail `json:"thumbnail,omitempty"`   // embed thumbnail object	thumbnail information
	Video       []*discordgo.MessageEmbedVideo     `json:"video,omitempty"`       // embed video object	video information
	Provider    []*discordgo.MessageEmbedProvider  `json:"provider,omitempty"`    // embed provider object	provider information
	Author      []*discordgo.MessageEmbedAuthor    `json:"author,omitempty"`      // embed author object	author information
	Fields      []*discordgo.MessageEmbedField     `json:"fields,omitempty"`      //	array of embed field objects	fields information
}

func buildEmbed(embedList []interface{}) (*discordgo.MessageEmbed, error) {
	embedMap := embedList[0].(map[string]interface{})

	embed := &discordgo.MessageEmbed{
		Title:       embedMap["title"].(string),
		Description: embedMap["description"].(string),
		URL:         embedMap["url"].(string),
		Color:       embedMap["color"].(int),
		Timestamp:   embedMap["timestamp"].(string),
	}

	if len(embedMap["footer"].([]interface{})) > 0 {
		footerMap := embedMap["footer"].([]interface{})[0].(map[string]interface{})
		embed.Footer = &discordgo.MessageEmbedFooter{
			Text:    footerMap["text"].(string),
			IconURL: footerMap["icon_url"].(string),
		}
	}

	if len(embedMap["image"].([]interface{})) > 0 {
		imageMap := embedMap["image"].([]interface{})[0].(map[string]interface{})
		embed.Image = &discordgo.MessageEmbedImage{
			URL:    imageMap["url"].(string),
			Width:  imageMap["width"].(int),
			Height: imageMap["height"].(int),
		}
	}

	if len(embedMap["thumbnail"].([]interface{})) > 0 {
		thumbnailMap := embedMap["thumbnail"].([]interface{})[0].(map[string]interface{})
		embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL:    thumbnailMap["url"].(string),
			Width:  thumbnailMap["width"].(int),
			Height: thumbnailMap["height"].(int),
		}
	}

	if len(embedMap["video"].([]interface{})) > 0 {
		videoMap := embedMap["video"].([]interface{})[0].(map[string]interface{})
		embed.Video = &discordgo.MessageEmbedVideo{
			URL:    videoMap["url"].(string),
			Width:  videoMap["width"].(int),
			Height: videoMap["height"].(int),
		}
	}

	if len(embedMap["provider"].([]interface{})) > 0 {
		providerMap := embedMap["provider"].([]interface{})[0].(map[string]interface{})
		embed.Provider = &discordgo.MessageEmbedProvider{
			URL:  providerMap["url"].(string),
			Name: providerMap["name"].(string),
		}
	}

	if len(embedMap["author"].([]interface{})) > 0 {
		authorMap := embedMap["author"].([]interface{})[0].(map[string]interface{})
		embed.Author = &discordgo.MessageEmbedAuthor{
			Name:    authorMap["name"].(string),
			URL:     authorMap["url"].(string),
			IconURL: authorMap["icon_url"].(string),
		}
	}

	for _, field := range embedMap["fields"].([]interface{}) {
		fieldMap := field.(map[string]interface{})

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   fieldMap["name"].(string),
			Value:  fieldMap["value"].(string),
			Inline: fieldMap["inline"].(bool),
		})
	}

	return embed, nil
}

func unbuildEmbed(embed *discordgo.MessageEmbed) []interface{} {
	var ret interface{}

	e := &UnmappedEmbed{
		Title:       embed.Title,
		Description: embed.Description,
		URL:         embed.URL,
		Timestamp:   embed.Timestamp,
		Color:       embed.Color,
		Fields:      embed.Fields,
	}

	if embed.Footer != nil {
		e.Footer = []*discordgo.MessageEmbedFooter{embed.Footer}
	}
	if embed.Image != nil {
		e.Image = []*discordgo.MessageEmbedImage{embed.Image}
	}
	if embed.Thumbnail != nil {
		e.Thumbnail = []*discordgo.MessageEmbedThumbnail{embed.Thumbnail}
	}
	if embed.Video != nil {
		e.Video = []*discordgo.MessageEmbedVideo{embed.Video}
	}
	if embed.Provider != nil {
		e.Provider = []*discordgo.MessageEmbedProvider{embed.Provider}
	}
	if embed.Author != nil {
		e.Author = []*discordgo.MessageEmbedAuthor{embed.Author}
	}

	j, _ := json.MarshalIndent(e, "", "    ")
	_ = json.Unmarshal(j, &ret)

	return []interface{}{ret}
}
