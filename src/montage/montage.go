package montage

import (
	"encoding/xml"
	"fmt"
	"os"
)

func (t *MontageClipList) LoadFile(fn *string) error {
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
	return nil
}
