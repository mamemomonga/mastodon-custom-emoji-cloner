package main

import (
	"./config"
	"./store"
	"./don"
	"./emoji"
	"fmt"
	"github.com/comail/colog"
	"log"
	"os"
	"time"
//	"github.com/davecgh/go-spew/spew"
)

var (
	Version  string
	Revision string
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {

	// colog 設定
	if Version == "" {
		colog.SetDefaultLevel(colog.LDebug)
		colog.SetMinLevel(colog.LTrace)
		colog.SetFormatter(&colog.StdFormatter{
			Colors: true,
			Flag:   log.Ldate | log.Ltime | log.Lshortfile,
		})
	} else {
		colog.SetDefaultLevel(colog.LDebug)
		colog.SetMinLevel(colog.LInfo)
		colog.SetFormatter(&colog.StdFormatter{
			Colors: true,
			Flag:   log.Ldate | log.Ltime,
		})
	}
	colog.Register()

	log.Printf("info: mastodon-custom-emoji-cloner")
	log.Printf("info: https://github.com/mamemomonga/mastodon-custom-emoji-cloner/")
	log.Printf("info:   Version:  %s\n", Version)
	log.Printf("info:   Revision: %s\n", Revision)

	if len(args) == 0 {
		usage()
		return 1
	}

	var err error

	// config 読込
	var cnf config.Config
	cnf, err = config.Load(os.Args[1])
	if err != nil {
		log.Printf("alert: %s", err)
		return 1
	}

	// data読書
	var stor *store.Store
	stor, err = store.NewStore(cnf.DataFile)
	if err != nil {
		log.Print("alert: %s", err)
		return 1
	}

	for {
		err = update(&cnf, stor)
		if err != nil {
			log.Print(err)
			return 1
		}
		log.Printf("info: Sleep %d minutes.", cnf.Duration)
		for i := 0; i < cnf.Duration; i++ {
			time.Sleep(time.Second * 60)
		}
	}

	return 0
}

func update(cnf *config.Config, stor *store.Store) (err error) {

	// DB接続
	var emj *emoji.Emoji
	emj, err = emoji.NewEmoji(*cnf)
	defer emj.Finalize()
	if err != nil {
		return
	}

	// 絵文字取得
	err = emj.Read()
	if err != nil {
		return
	}

	// コピー処理
	var new_emojis []emoji.NewCustomEmoji
	new_emojis,err = emj.Copy()
	if err != nil {
		return
	}

	if len(new_emojis) > 0 {

		// トゥート文言作成
		message := "隊長！新しい絵文字が到着しました！\r\n"
		for _,em := range new_emojis {
			message = message + fmt.Sprintf(":%s: ",em.Shortcode)
		}
		message = message + "\r\n"

		// マストドン
		var dn *don.Don
		dn, err = don.NewDon(cnf, stor)
		if err != nil {
			return
		}

		// トゥート
		err = dn.Toot(message)
		if err != nil {
			return
		}
	}

	return nil
}

func usage() {
	fmt.Printf("\n  USAGE: %s config\n\n", os.Args[0])
}
