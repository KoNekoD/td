// Code generated by mkentity, DO NOT EDIT.
package entity

import (
	"context"

	"github.com/KoNekoD/td/tg"
)

var (
	_ = tg.Invoker(nil)
	_ = context.Context(nil)
)

// Unknown creates Formatter of Unknown message entity.
//
// See https://core.telegram.org/constructor/messageEntityUnknown.
func Unknown() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityUnknown{
			Offset: offset,
			Length: length,
		}
	}
}

// Unknown adds and formats message as Unknown message entity.
//
// See https://core.telegram.org/constructor/messageEntityUnknown.
func (b *Builder) Unknown(s string) *Builder {
	return b.Format(s, Unknown())
}

// Mention creates Formatter of Mention message entity.
//
// See https://core.telegram.org/constructor/messageEntityMention.
func Mention() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityMention{
			Offset: offset,
			Length: length,
		}
	}
}

// Mention adds and formats message as Mention message entity.
//
// See https://core.telegram.org/constructor/messageEntityMention.
func (b *Builder) Mention(s string) *Builder {
	return b.Format(s, Mention())
}

// Hashtag creates Formatter of Hashtag message entity.
//
// See https://core.telegram.org/constructor/messageEntityHashtag.
func Hashtag() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityHashtag{
			Offset: offset,
			Length: length,
		}
	}
}

// Hashtag adds and formats message as Hashtag message entity.
//
// See https://core.telegram.org/constructor/messageEntityHashtag.
func (b *Builder) Hashtag(s string) *Builder {
	return b.Format(s, Hashtag())
}

// BotCommand creates Formatter of BotCommand message entity.
//
// See https://core.telegram.org/constructor/messageEntityBotCommand.
func BotCommand() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityBotCommand{
			Offset: offset,
			Length: length,
		}
	}
}

// BotCommand adds and formats message as BotCommand message entity.
//
// See https://core.telegram.org/constructor/messageEntityBotCommand.
func (b *Builder) BotCommand(s string) *Builder {
	return b.Format(s, BotCommand())
}

// URL creates Formatter of URL message entity.
//
// See https://core.telegram.org/constructor/messageEntityUrl.
func URL() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityURL{
			Offset: offset,
			Length: length,
		}
	}
}

// URL adds and formats message as URL message entity.
//
// See https://core.telegram.org/constructor/messageEntityUrl.
func (b *Builder) URL(s string) *Builder {
	return b.Format(s, URL())
}

// Email creates Formatter of Email message entity.
//
// See https://core.telegram.org/constructor/messageEntityEmail.
func Email() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityEmail{
			Offset: offset,
			Length: length,
		}
	}
}

// Email adds and formats message as Email message entity.
//
// See https://core.telegram.org/constructor/messageEntityEmail.
func (b *Builder) Email(s string) *Builder {
	return b.Format(s, Email())
}

// Bold creates Formatter of Bold message entity.
//
// See https://core.telegram.org/constructor/messageEntityBold.
func Bold() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityBold{
			Offset: offset,
			Length: length,
		}
	}
}

// Bold adds and formats message as Bold message entity.
//
// See https://core.telegram.org/constructor/messageEntityBold.
func (b *Builder) Bold(s string) *Builder {
	return b.Format(s, Bold())
}

// Italic creates Formatter of Italic message entity.
//
// See https://core.telegram.org/constructor/messageEntityItalic.
func Italic() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityItalic{
			Offset: offset,
			Length: length,
		}
	}
}

// Italic adds and formats message as Italic message entity.
//
// See https://core.telegram.org/constructor/messageEntityItalic.
func (b *Builder) Italic(s string) *Builder {
	return b.Format(s, Italic())
}

// Code creates Formatter of Code message entity.
//
// See https://core.telegram.org/constructor/messageEntityCode.
func Code() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityCode{
			Offset: offset,
			Length: length,
		}
	}
}

// Code adds and formats message as Code message entity.
//
// See https://core.telegram.org/constructor/messageEntityCode.
func (b *Builder) Code(s string) *Builder {
	return b.Format(s, Code())
}

// Pre creates Formatter of Pre message entity.
//
// See https://core.telegram.org/constructor/messageEntityPre.
func Pre(language string) Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityPre{
			Offset:   offset,
			Length:   length,
			Language: language,
		}
	}
}

// Pre adds and formats message as Pre message entity.
//
// See https://core.telegram.org/constructor/messageEntityPre.
func (b *Builder) Pre(s string, language string) *Builder {
	return b.Format(s, Pre(language))
}

// TextURL creates Formatter of TextURL message entity.
//
// See https://core.telegram.org/constructor/messageEntityTextUrl.
func TextURL(uRL string) Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityTextURL{
			Offset: offset,
			Length: length,
			URL:    uRL,
		}
	}
}

// TextURL adds and formats message as TextURL message entity.
//
// See https://core.telegram.org/constructor/messageEntityTextUrl.
func (b *Builder) TextURL(s string, uRL string) *Builder {
	return b.Format(s, TextURL(uRL))
}

// MentionName creates Formatter of MentionName message entity.
//
// See https://core.telegram.org/constructor/inputMessageEntityMentionName.
func MentionName(userID tg.InputUserClass) Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.InputMessageEntityMentionName{
			Offset: offset,
			Length: length,
			UserID: userID,
		}
	}
}

// MentionName adds and formats message as MentionName message entity.
//
// See https://core.telegram.org/constructor/inputMessageEntityMentionName.
func (b *Builder) MentionName(s string, userID tg.InputUserClass) *Builder {
	return b.Format(s, MentionName(userID))
}

// Phone creates Formatter of Phone message entity.
//
// See https://core.telegram.org/constructor/messageEntityPhone.
func Phone() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityPhone{
			Offset: offset,
			Length: length,
		}
	}
}

// Phone adds and formats message as Phone message entity.
//
// See https://core.telegram.org/constructor/messageEntityPhone.
func (b *Builder) Phone(s string) *Builder {
	return b.Format(s, Phone())
}

// Cashtag creates Formatter of Cashtag message entity.
//
// See https://core.telegram.org/constructor/messageEntityCashtag.
func Cashtag() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityCashtag{
			Offset: offset,
			Length: length,
		}
	}
}

// Cashtag adds and formats message as Cashtag message entity.
//
// See https://core.telegram.org/constructor/messageEntityCashtag.
func (b *Builder) Cashtag(s string) *Builder {
	return b.Format(s, Cashtag())
}

// Underline creates Formatter of Underline message entity.
//
// See https://core.telegram.org/constructor/messageEntityUnderline.
func Underline() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityUnderline{
			Offset: offset,
			Length: length,
		}
	}
}

// Underline adds and formats message as Underline message entity.
//
// See https://core.telegram.org/constructor/messageEntityUnderline.
func (b *Builder) Underline(s string) *Builder {
	return b.Format(s, Underline())
}

// Strike creates Formatter of Strike message entity.
//
// See https://core.telegram.org/constructor/messageEntityStrike.
func Strike() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityStrike{
			Offset: offset,
			Length: length,
		}
	}
}

// Strike adds and formats message as Strike message entity.
//
// See https://core.telegram.org/constructor/messageEntityStrike.
func (b *Builder) Strike(s string) *Builder {
	return b.Format(s, Strike())
}

// BankCard creates Formatter of BankCard message entity.
//
// See https://core.telegram.org/constructor/messageEntityBankCard.
func BankCard() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityBankCard{
			Offset: offset,
			Length: length,
		}
	}
}

// BankCard adds and formats message as BankCard message entity.
//
// See https://core.telegram.org/constructor/messageEntityBankCard.
func (b *Builder) BankCard(s string) *Builder {
	return b.Format(s, BankCard())
}

// Spoiler creates Formatter of Spoiler message entity.
//
// See https://core.telegram.org/constructor/messageEntitySpoiler.
func Spoiler() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntitySpoiler{
			Offset: offset,
			Length: length,
		}
	}
}

// Spoiler adds and formats message as Spoiler message entity.
//
// See https://core.telegram.org/constructor/messageEntitySpoiler.
func (b *Builder) Spoiler(s string) *Builder {
	return b.Format(s, Spoiler())
}

// CustomEmoji creates Formatter of CustomEmoji message entity.
//
// See https://core.telegram.org/constructor/messageEntityCustomEmoji.
func CustomEmoji(documentID int64) Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityCustomEmoji{
			Offset:     offset,
			Length:     length,
			DocumentID: documentID,
		}
	}
}

// CustomEmoji adds and formats message as CustomEmoji message entity.
//
// See https://core.telegram.org/constructor/messageEntityCustomEmoji.
func (b *Builder) CustomEmoji(s string, documentID int64) *Builder {
	return b.Format(s, CustomEmoji(documentID))
}

// Blockquote creates Formatter of Blockquote message entity.
//
// See https://core.telegram.org/constructor/messageEntityBlockquote.
func Blockquote() Formatter {
	return func(offset, length int) tg.MessageEntityClass {
		return &tg.MessageEntityBlockquote{
			Offset: offset,
			Length: length,
		}
	}
}

// Blockquote adds and formats message as Blockquote message entity.
//
// See https://core.telegram.org/constructor/messageEntityBlockquote.
func (b *Builder) Blockquote(s string) *Builder {
	return b.Format(s, Blockquote())
}
