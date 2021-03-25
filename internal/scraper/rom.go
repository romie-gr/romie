package scraper

import "fmt"

type Rom struct {
	Name         string `json:"name"`
	Console      string `json:"console"`
	Language     string `json:"language"`
	Link         string	`json:"link"`
	DownloadLink string	`json:"downloadLink"`
	Version      string	`json:"version"`
	ReleaseYear  int	`json:"releaseYear"`
}

// NewRom factory constructor
func NewRom(name string, console string, language string, link string, downloadLink string) *Rom {
	return &Rom{
		Name:         name,
		Console:      console,
		Language:     language,
		Link:         link,
		DownloadLink: downloadLink,
		Version:      "0.0",
		ReleaseYear:  1970,
	}
}

func (r *Rom) GetName() string {
	return r.Name
}

func (r *Rom) GetVersion() string {
	return r.Version
}

func (r *Rom) SetVersion(version string) {
	r.Version = version
}

func (r *Rom) GetReleaseYear() int {
	return r.ReleaseYear
}

func (r *Rom) SetReleaseYear(releaseYear int) {
	r.ReleaseYear = releaseYear
}

func (r *Rom) Stringer() string {
	return fmt.Sprintf(`Game: %s (v. %s) for %s. Released in %d [Lang: %s]
Parsed via %s
Download via %s`,
		r.Name,
		r.Version,
		r.Console,
		r.ReleaseYear,
		r.Language,
		r.Link,
		r.DownloadLink,
	)
}
