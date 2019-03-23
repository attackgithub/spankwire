package spankwire

type SpankwireVideo struct {
	ID           float64 `json:"video_id,omitempty"`
	Duration     string  `json:"duration,omitempty"`
	Views        float64 `json:"views,omitempty"`
	Rating       string  `json:"rating,omitempty"`
	Ratings      string  `json:"ratings,omitempty"`
	Title        string  `json:"title,omitempty"`
	URL          string  `json:"url,omitempty"`
	DefaultThumb string  `json:"default_thumb,omitempty"`
	Thumb        string  `json:"thumb,omitempty"`
	PublishDate  string  `json:"publish_date,omitempty"`
	Thumbs       []SpankwireThumb
}

type SpankwireThumb struct {
	Size   string `json:"size,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
	Src    string `json:"src,omitempty"`
}

type SpankwireTag struct {
	Name string `json:"tag_name,omitempty"`
}
