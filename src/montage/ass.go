package montage

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func (t *MontageClipList) Convert() {
	var lines []ASSLine
	for _, c := range t.Clips {
		startDur := parseTimeToDuration(c.StartInMontage)
		// 表示時間はそのクリップが終わる3秒前まで
		endDur := startDur + parseTimeToDuration(c.Length) - time.Second*3
		lines = append(lines, ASSLine{
			Name:  cleanName(c.Name),
			Start: formatDurationToASS(startDur),
			End:   formatDurationToASS(endDur),
		})
	}
	t.assdata = lines
}

func (t *MontageClipList) Template(tplPath *string, outPath *string) error {

	// テンプレート読み込み
	tplText, err := os.ReadFile(*tplPath)
	if err != nil {
		return fmt.Errorf("テンプレート読み込みエラー: %w", err)
	}

	tpl, err := template.New("ass").Parse(string(tplText))
	if err != nil {
		return fmt.Errorf("テンプレート構文エラー: %w", err)
	}

	outFile, err := os.Create(*outPath)
	if err != nil {
		return fmt.Errorf("出力ファイル作成エラー: %w", err)
	}
	defer outFile.Close()

	if err := tpl.Execute(outFile, t.assdata); err != nil {
		return fmt.Errorf("テンプレート実行エラー: %w", err)
	}

	fmt.Printf("ASSファイルを出力しました: %s\n", *outPath)
	return nil
}
