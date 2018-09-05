package main

import (
	"./config"
	"./emoji"
	"fmt"
	"log"
	"os"
	"time"
	// "github.com/davecgh/go-spew/spew"
)

var (
	commit  string
	builtAt string
	builtBy string
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	log.Printf("EMOJI CLONER")
	log.Printf("  commit: %s\n", commit)
	log.Printf("  build at: %s build by: %s\n", builtAt, builtBy)
	log.Print("STARTUP")

	if len(args) == 0 {
		usage()
		return 1
	}

	var err error

	// config.yaml読込
	var cnf config.Config
	cnf, err = config.ReadConfig(os.Args[1])
	if err != nil {
		log.Print(err)
		return 1
	}

	for {
		err = update(cnf)
		if err != nil {
			log.Print(err)
			return 1
		}
		if cnf.Verbose {
			log.Printf("Sleep %d minutes.", cnf.Duration)
		}
		for i := 0; i < cnf.Duration; i++ {
			time.Sleep(time.Second * 60)
		}
	}

	return 0
}

func update(cnf config.Config) (err error) {
	// DB接続
	var emj *emoji.Emoji
	emj, err = emoji.NewEmoji(cnf)
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
	err = emj.Copy()
	if err != nil {
		return
	}

	return nil
}

func usage() {
	fmt.Printf("\n\n  USAGE: %s config\n\n", os.Args[0])
}
