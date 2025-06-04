package montage

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/mamemomonga/MontageClipList2ASS/src/config"
)

func NewMontageClipList(cfg *config.Config) (t *MontageClipList) {
	t = new(MontageClipList)
	t.cfg = cfg
	return t
}

func (t *MontageClipList) MontageXMLLoad(fn *string) error {
	if *fn == "" {
		return fmt.Errorf("XMLファイルを -f で指定してください")
	}
	xmlFile, err := os.Open(*fn)
	if err != nil {
		return fmt.Errorf("XML読み込みエラー: %w", err)
	}
	defer xmlFile.Close()

	var c ClipList
	if err := xml.NewDecoder(xmlFile).Decode(&c); err != nil {
		return fmt.Errorf("XMLの読み込みに失敗: %w", err)
	}
	t.Clips = c.Clips

	if err := t.assFilter(); err != nil {
		return err
	}

	return nil
}
