package model

type GameMode string

const (
	GameModeCoop       GameMode = "coop"
	GameModeSteamCloud GameMode = "steam_cloud"
)

type Game struct {
	id            int        `json:"id"`
	Name          string     `json:"name"`
	DevStudioInfo DevStudio  `json:""`
	Publisher     string     `json:"publisher"`
	Announcement  string     `json:"announcement"`
	Release       string     `json:"release"`
	GenreInfo     Genre      `json:""`
	RRatingInfo   RRating    `json:""`
	ModeInfo      []GameMode `json:""`
	Rating        float32    `json:"rating"`
	Price         float32    `json:"price"`
	DLC           string     `json:"dlc"`
	Localization  string     `json:"localization"`
	DevPhazeInfo  DevPhaze   `json:""`
	SaleInfo      Sale       `json:""`
	IdShop        int        `json:"IdShop"`
	OrderInfo     Order      `json:""`
	Availability  string     `json:"availability"`
}
