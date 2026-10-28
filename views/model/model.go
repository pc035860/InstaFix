package model

type ViewsData struct {
	Card        string
	Title       string `default:"InstaFix"`
	ImageURL    string `default:""`
	ImageURLs   []string
	VideoURL    string `default:""`
	URL         string
	Description string
	OEmbedURL   string
	Width       int `default:"400"`
	Height      int `default:"400"`
}

type OEmbedData struct {
	Text string
	URL  string
}

// Images returns every preview image, falling back to the single ImageURL.
func (v *ViewsData) Images() []string {
	if len(v.ImageURLs) > 0 {
		return v.ImageURLs
	}
	if v.ImageURL != "" {
		return []string{v.ImageURL}
	}
	return nil
}
