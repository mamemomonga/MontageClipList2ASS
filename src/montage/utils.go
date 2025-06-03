package montage

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// parseTime converts "2 mn 13 s 639 ms" → "0:02:13.64"
func parseTime(t string) string {
	min := 0
	sec := 0
	ms := 0

	re := regexp.MustCompile(`(\d+)\s*mn`)
	if match := re.FindStringSubmatch(t); len(match) > 1 {
		min, _ = strconv.Atoi(match[1])
	}
	re = regexp.MustCompile(`(\d+)\s*s`)
	if match := re.FindStringSubmatch(t); len(match) > 1 {
		sec, _ = strconv.Atoi(match[1])
	}
	re = regexp.MustCompile(`(\d+)\s*ms`)
	if match := re.FindStringSubmatch(t); len(match) > 1 {
		ms, _ = strconv.Atoi(match[1])
	}

	total := fmt.Sprintf("0:%02d:%02d.%02d", min, sec, ms/10)
	return total
}

// "2 mn 13 s 639 ms" → time.Duration
func parseTimeToDuration(s string) time.Duration {
	min := 0
	sec := 0
	ms := 0
	re := regexp.MustCompile(`(\d+)\s*mn`)
	if m := re.FindStringSubmatch(s); len(m) > 1 {
		min, _ = strconv.Atoi(m[1])
	}
	re = regexp.MustCompile(`(\d+)\s*s`)
	if m := re.FindStringSubmatch(s); len(m) > 1 {
		sec, _ = strconv.Atoi(m[1])
	}
	re = regexp.MustCompile(`(\d+)\s*ms`)
	if m := re.FindStringSubmatch(s); len(m) > 1 {
		ms, _ = strconv.Atoi(m[1])
	}
	return time.Duration(min)*time.Minute + time.Duration(sec)*time.Second + time.Duration(ms)*time.Millisecond
}

// time.Duration → "0:02:13.63"（ASS形式）
func formatDurationToASS(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	cs := (d.Milliseconds() % 1000) / 10 // centiseconds
	return fmt.Sprintf("%d:%02d:%02d.%02d", h, m, s, cs)
}

func cleanName(raw string) string {
	re := regexp.MustCompile(`^\d+\.\s*`)
	cleaned := re.ReplaceAllString(raw, "")
	return strings.TrimSpace(cleaned)
}
