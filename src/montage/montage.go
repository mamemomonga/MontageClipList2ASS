package montage

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/mamemomonga/MontageClipList2ASS/src/config"
)

type MontageClipList struct {
	cfg   *config.Config
	clips []Clip
	d     []TData
}

func NewMontageClipList(cfg *config.Config) (t *MontageClipList) {
	t = new(MontageClipList)
	t.cfg = cfg
	return t
}

func (t *MontageClipList) XMLLoad(fn *string) error {
	if *fn == "" {
		return fmt.Errorf("XMLファイルを -f で指定してください")
	}
	xmlFile, err := os.Open(*fn)
	if err != nil {
		return fmt.Errorf("XML読み込みエラー: %w", err)
	}
	defer xmlFile.Close()

	var cliplist ClipList

	if err := xml.NewDecoder(xmlFile).Decode(&cliplist); err != nil {
		return fmt.Errorf("XMLの読み込みに失敗: %w", err)
	}
	t.clips = cliplist.Clips

	if err := t.buildDatas(); err != nil {
		return err
	}

	return nil
}

func (t *MontageClipList) buildDatas() error {

	// クリップのテンプレート
	tpc, err := template.New("T").Parse(t.cfg.ASSFile.Template)
	if err != nil {
		return err
	}

	// Indexのテンプレート
	tpi, err := template.New("T").Parse(t.cfg.YouTubeIndexFile.Template)
	if err != nil {
		return err
	}

	// データの構築
	t.d = nil
	for i, c := range t.clips {
		startDur := parseTimeToDuration(c.StartInMontage)
		var endDur time.Duration
		hideBefore := time.Second * time.Duration(t.cfg.ASSFile.HideBefore)

		if i+1 < len(t.clips) {
			// 一つ前のクリップが始まる部分
			nextStart := parseTimeToDuration(t.clips[i+1].StartInMontage)
			endDur = nextStart - hideBefore
			if endDur < startDur {
				endDur = startDur // 最低限、EndがStartを下回らないように
			}
		} else {
			// 最後のクリップは自分の長さ
			endDur = startDur + parseTimeToDuration(c.Length) - hideBefore
		}

		// テンプレート用データ
		td := TLineTemplateData{
			Name:      strings.TrimSpace(c.Name),
			StartTime: formatDurationToYouTube(startDur),
		}

		var bfc bytes.Buffer
		if err := tpc.Execute(&bfc, td); err != nil {
			return err
		}
		var bfi bytes.Buffer
		if err := tpi.Execute(&bfi, td); err != nil {
			return err
		}

		// 出力用データの構築
		t.d = append(t.d, TData{
			Caption:      bfc.String(),
			YouTubeIndex: bfi.String(),
			EventStart:   formatDurationToASS(startDur),
			EventEnd:     formatDurationToASS(endDur),
		})
	}
	return nil
}

func (t *MontageClipList) ASSWrite(fn *string) error {
	b := `[Script Info]
Title: MontageClipList2ASS
ScriptType: v4.00+

` + t.assStyle() + `

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`
	for _, c := range t.d {
		b = b + "Dialogue: 0," + c.EventStart + "," + c.EventEnd + ",Default,,0,0,0,," + c.Caption + "\n"
	}

	outFile, err := os.Create(*fn)
	if err != nil {
		return fmt.Errorf("出力ファイル作成エラー: %w", err)
	}
	defer outFile.Close()

	outFile.WriteString(b)
	fmt.Printf("ASSファイルを出力しました: %s\n", *fn)
	return nil
}

func (t *MontageClipList) assStyle() string {
	b := "[V4+ Styles]\n"
	s := t.cfg.ASSFile.Style
	b = b + "Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding\n"
	b = b + fmt.Sprintf("Style: %s,%s,%d,%s,%s,%s,%s,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
		s.Name,
		s.Fontname,
		s.Fontsize,
		s.PrimaryColour,
		s.SecondaryColour,
		s.OutlineColour,
		s.BackColour,
		s.Bold,
		s.Italic,
		s.Underline,
		s.StrikeOut,
		s.ScaleX,
		s.ScaleY,
		s.Spacing,
		s.Angle,
		s.BorderStyle,
		s.Outline,
		s.Shadow,
		s.Alignment,
		s.MarginL,
		s.MarginR,
		s.MarginV,
		s.Encoding,
	)
	return b
}

func (t *MontageClipList) YouTubeIndexWrite(fn *string) error {
	var s string
	for _, c := range t.d {
		s = s + fmt.Sprintf("%s\n", c.YouTubeIndex)
	}

	outFile, err := os.Create(*fn)
	if err != nil {
		return fmt.Errorf("出力ファイル作成エラー: %w", err)
	}
	defer outFile.Close()

	outFile.WriteString(s)
	fmt.Printf("Indexファイルを出力しました: %s\n", *fn)
	return nil
}
