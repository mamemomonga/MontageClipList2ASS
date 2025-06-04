package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mamemomonga/MontageClipList2ASS/src/config"
	"github.com/mamemomonga/MontageClipList2ASS/src/montage"
)

var (
	version  string
	revision string
)

func main() {
	verString := fmt.Sprintf("%s-%s", version, revision)
	configPath := flag.String("c", "config.yaml", "Configファイルのパス")
	xmlPath := flag.String("f", "", "XMLファイルのパス")
	outPath := flag.String("o", "output.ass", "出力ファイル名")
	//	idxPath := flag.String("x", "index.txt", "YouTubeインデックスファイル名")
	showVersion := flag.Bool("v", false, "バージョン表示")
	flag.Parse()

	if *showVersion {
		fmt.Printf("montage_cliplist2ass version %s\n", verString)
		os.Exit(0)
	}

	cfg := config.NewConfig()
	cfg.Load(configPath)

	mcl := montage.NewMontageClipList(cfg)

	if err := mcl.MontageXMLLoad(xmlPath); err != nil {
		log.Fatal(fmt.Errorf("XMLエラー: %w", err))
	}

	if err := mcl.ASSWriteFile(outPath); err != nil {
		log.Fatal(fmt.Errorf("ASSファイル: %w", err))
	}

}
