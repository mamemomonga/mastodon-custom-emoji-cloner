package emoji

import (
	"../config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type CustomEmoji struct {
	id                 int
	shortcode          string
	domain             sql.NullString
	image_file_name    string
	image_content_type string
	image_file_size    string
	image_updated_at   string
	created_at         string
	updated_at         string
	disabled           bool
	url                string
	image_remote_url   string
	visible_in_picker  bool
}

type EmojisSC map[string][]CustomEmoji

type Emoji struct {
	db     *sql.DB
	config config.Config
	Local  EmojisSC
	Remote EmojisSC
}

func NewEmoji(config config.Config) (this *Emoji, err error) {
	err = nil
	this = new(Emoji)

	this.Local = make(EmojisSC)
	this.Remote = make(EmojisSC)

	this.config = config
	dcnf := config.Database

	if this.config.Verbose {
		log.Printf("Connecting DB: %s", dcnf.DBname)
	}
	this.db, err = sql.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dcnf.Host, dcnf.Port, dcnf.User, dcnf.Password, dcnf.DBname))

	return
}

func (this *Emoji) Finalize() {
	if this.config.Verbose {
		log.Print("Disconnect DB")
	}
	defer this.db.Close()
}
