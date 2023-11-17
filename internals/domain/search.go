package domain

type PexelsResponse struct {
	TotalResults int `json:"total_results"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	Photos       []struct {
		ID              int    `json:"id"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
		URL             string `json:"url"`
		Photographer    string `json:"photographer"`
		PhotographerURL string `json:"photographer_url"`
		PhotographerID  int    `json:"photographer_id"`
		AvgColor        string `json:"avg_color"`
		Src             struct {
			Original  string `json:"original"`
			Large2x   string `json:"large2x"`
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Small     string `json:"small"`
			Portrait  string `json:"portrait"`
			Landscape string `json:"landscape"`
			Tiny      string `json:"tiny"`
		} `json:"src"`
		Liked bool   `json:"liked"`
		Alt   string `json:"alt"`
	} `json:"photos"`
	NextPage string `json:"next_page"`
}
