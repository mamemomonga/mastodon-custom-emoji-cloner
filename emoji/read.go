package emoji

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func (this *Emoji) Read() (err error) {
	err = nil

	// データベースから取得
	emjs := []CustomEmoji{}

	var rows *sql.Rows
	rows, err = this.db.Query("SELECT * FROM custom_emojis")
	if err != nil {
		return
	}

	// 取得
	for rows.Next() {
		var em CustomEmoji
		rows.Scan(&em.id,
			&em.shortcode,
			&em.domain,
			&em.image_file_name,
			&em.image_content_type,
			&em.image_file_size,
			&em.image_updated_at,
			&em.created_at,
			&em.updated_at,
			&em.disabled,
			&em.url,
			&em.image_remote_url,
			&em.visible_in_picker)
		emjs = append(emjs, em)
	}

	// shortcodeをキーにした配列に入れる処理
	appender := func(emj CustomEmoji, emr EmojisSC) {
		shortcode := emj.shortcode
		if _, ok := emr[shortcode]; ok {
			emr[shortcode] = append(emr[shortcode], emj)
		} else {
			emr[shortcode] = []CustomEmoji{emj}
		}
	}

	// LocalとRemoteを分ける
	for i := 0; i < len(emjs); i++ {
		if emjs[i].domain.Valid {
			appender(emjs[i], this.Remote)
		} else {
			appender(emjs[i], this.Local)
		}
	}

	if this.config.Verbose {
		log.Printf("Found %d custom_emojis", len(emjs))
	}

	if this.config.Verbose {
		log.Printf("Found %d custom_emojis", len(emjs))
		log.Printf("Local Emojis:  %d", len(this.Local))
		log.Printf("Remote Emojis: %d", len(this.Remote))
	}

	return
}
