package sitemap

import (
	"encoding/xml"
	"time"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	ChangeFreq string    `xml:"changefreq"`
	Priority   float64   `xml:"priority"`
}

func NewURLSet(urls []URL) *URLSet {
	return &URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}
}

func (urlSet *URLSet) ToXML() ([]byte, error) {
	xmlBytes, err := xml.MarshalIndent(urlSet, "", "  ")
	if err != nil {
		return nil, err
	}
	return []byte(xml.Header + string(xmlBytes)), nil
}
