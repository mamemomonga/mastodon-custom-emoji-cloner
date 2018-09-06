package emoji

import (
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"path/filepath"
)

type NewCustomEmoji struct {
	FromDomain string
	Shortcode  string
	ID         int
}

// ファイルが存在する
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// パスを得る
func (this *Emoji) ce2path(ce CustomEmoji, id int, image_type string) string {
	v := fmt.Sprintf("%09d", id)
	return this.config.EmojiImagesPath + "/" + v[0:3] + "/" + v[3:6] + "/" + v[6:9] + "/" + image_type + "/" + ce.image_file_name
}

// 絵文字のコピー
func (this *Emoji) copy_emoji(ce CustomEmoji, newid int, image_type string) error {

	src := this.ce2path(ce, ce.id, image_type)
	dst := this.ce2path(ce, newid, image_type)

	log.Printf("trace:  Src  %s\n", src)
	log.Printf("trace:  Dest %s\n", dst)

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()
	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}
	return nil
}

func (this *Emoji) Copy() (new_emojis []NewCustomEmoji, err error) {
	err = nil
	new_emojis = []NewCustomEmoji{}

	// コピー対象の収集
	targets := []CustomEmoji{}
	for key, val := range this.Remote {
		// ローカルにないものだけ集める
		if _, ok := this.Local[key]; !ok {
			// 任意(たまたま先頭にあった)インスタンスのものを取得
			src := val[0]
			// originalファイルがあるもののみ収集
			if bool, _ := exists(this.ce2path(src, src.id, "original")); bool {
				targets = append(targets, src)
			}
		}
	}
	// DBへ挿入
	for _, t := range targets {
		log.Printf("trace: [COPY]")
		log.Printf("trace:   ID:        %d", t.id)
		log.Printf("trace:   Shortcode: %s", t.shortcode)
		log.Printf("trace:   Domain:    %s", t.domain.String)
		log.Printf("trace:   Filename:  %s", t.image_file_name)
		var newid int
		err = this.db.QueryRow(`
INSERT INTO custom_emojis(
	shortcode,
	image_file_name,
	image_content_type,
	image_file_size,
	image_updated_at,
	created_at,
	updated_at,
	disabled,
	visible_in_picker
) VALUES (
	$1, $2, $3, $4, NOW(), NOW(), NOW(), false, true
) RETURNING id`,
			t.shortcode,
			t.image_file_name,
			t.image_content_type,
			t.image_file_size,
		).Scan(&newid)
		if err != nil {
			return
		}
		log.Printf("trace:  NewID:  %d", newid)

		this.copy_emoji(t, newid, "original")
		this.copy_emoji(t, newid, "static")

		new_emojis=append(new_emojis, NewCustomEmoji{
			FromDomain: t.domain.String,
			Shortcode:  t.shortcode,
			ID:         newid,
		})
	}
	return
}
