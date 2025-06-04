package montage

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (t *MontageClipList) ASSWriteFile(fn *string) error {
	b := `[Script Info]
Title: MontageClipList2ASS
ScriptType: v4.00+

` + t.assStyle() + `

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`
	for _, c := range t.assFilter() {
		b = b + "Dialogue: 0," + c.Start + "," + c.End + ",Default,,0,0,0,," + c.Name + "\n"
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

func (t *MontageClipList) assFilter() []ASSLine {
	var lines []ASSLine
	for i, c := range t.Clips {
		startDur := parseTimeToDuration(c.StartInMontage)
		var endDur time.Duration
		hideBefore := time.Second * time.Duration(t.cfg.ASSFile.HideBefore)

		if i+1 < len(t.Clips) {
			// 一つ前のクリップが始まる部分
			nextStart := parseTimeToDuration(t.Clips[i+1].StartInMontage)
			endDur = nextStart - hideBefore
			if endDur < startDur {
				endDur = startDur // 最低限、EndがStartを下回らないように
			}
		} else {
			// 最後のクリップは自分の長さ
			endDur = startDur + parseTimeToDuration(c.Length) - hideBefore
		}
		lines = append(lines, ASSLine{
			Name:  t.cfg.ASSFile.Filter.Pre + strings.TrimSpace(c.Name) + t.cfg.ASSFile.Filter.Post,
			Start: formatDurationToASS(startDur),
			End:   formatDurationToASS(endDur),
		})
	}
	return lines
}
