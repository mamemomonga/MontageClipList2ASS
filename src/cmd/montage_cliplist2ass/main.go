package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mamemomonga/MontageClipList2ASS/src/montage"
)

var (
	version  string
	revision string
)

func main() {
	verString := fmt.Sprintf("%s-%s", version, revision)
	xmlPath := flag.String("f", "", "XMLファイルのパス")
	tplPath := flag.String("t", "ass.tpl", "ASSテンプレートファイル")
	outPath := flag.String("o", "output.ass", "出力ファイル名")
	showVersion := flag.Bool("v", false, "バージョン表示")
	flag.Parse()

	if *showVersion {
		fmt.Printf("montage_cliplist2ass version %s\n", verString)
		os.Exit(0)
	}
	mcl := montage.NewMontageClipList()
	if err := mcl.LoadFile(xmlPath); err != nil {
		log.Fatal(fmt.Errorf("XMLエラー: %w", err))
	}

	mcl.Convert()
	if err := mcl.Template(tplPath, outPath); err != nil {
		log.Fatal(err)
	}

	// spew.Dump(mcl.Clips)

}
