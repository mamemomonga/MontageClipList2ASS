package montage

import "encoding/xml"

type MontageClipList struct {
	Clips   []Clip
	assdata []ASSLine
}

func NewMontageClipList() (t *MontageClipList) {
	t = new(MontageClipList)
	return t
}

type ClipList struct {
	XMLName xml.Name `xml:"ClipList"`
	Clips   []Clip   `xml:"Clip"`
}

type Clip struct {
	XMLName        xml.Name `xml:"Clip"`
	Name           string   `xml:"Name"`
	File           string   `xml:"File"`
	Track          string   `xml:"Track"`
	StartInMontage string   `xml:"StartInMontage"`
	StartInSource  string   `xml:"StarInSource"`
	Length         string   `xml:"Length"`
	FadeInLength   string   `xml:"FadeInLength"`
	FadeOutLength  string   `xml:"FadeOutLength"`
	PreGain        string   `xml:"PreGain"`
	PostGain       string   `xml:"PostGain"`
	Comment        string   `xml:"Comment"`
}

type ASSLine struct {
	Name  string
	Start string
	End   string
}
