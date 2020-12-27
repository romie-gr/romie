package scraper

import "fmt"

type Rom struct {
	name         string
	console      string
	language     string
	link         string
	downloadLink string
	version      string
	releaseYear  int
}

// NewRom factory constructor
func NewRom(name string, console string, language string, link string, downloadLink string) *Rom {
	return &Rom{
		name:         name,
		console:      console,
		language:     language,
		link:         link,
		downloadLink: downloadLink,
		version:      "0.0",
		releaseYear:  1970,
	}
}

func (r *Rom) Name() string {
	return r.name
}

func (r *Rom) Version() string {
	return r.version
}

func (r *Rom) SetVersion(version string) {
	r.version = version
}

func (r *Rom) ReleaseYear() int {
	return r.releaseYear
}

func (r *Rom) SetReleaseYear(releaseYear int) {
	r.releaseYear = releaseYear
}

func (r *Rom) Stringer() string {
	return fmt.Sprintf(`Game: %s (v. %s) for %s. Released in %d [Lang: %s]
Parsed via %s
Download via %s`,
		r.name,
		r.version,
		r.console,
		r.releaseYear,
		r.language,
		r.link,
		r.downloadLink,
	)
}
