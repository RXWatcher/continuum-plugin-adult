package tpdb

// Subset of the TPDB JSON schema we care about. Fields we don't use are
// intentionally omitted to keep the surface tight; add them as needed.

type sceneDTO struct {
	ID          string        `json:"id"`
	Slug        string        `json:"slug"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Trailer     string        `json:"trailer"`
	Date        string        `json:"date"`
	Duration    int           `json:"duration"`
	Poster      string        `json:"poster"`
	Background  imageDTO      `json:"background"`
	Site        siteRefDTO    `json:"site"`
	Performers  []performerDTO `json:"performers"`
	Directors   []directorDTO  `json:"directors"`
	Tags        []tagDTO       `json:"tags"`
}

type movieDTO struct {
	ID          string         `json:"id"`
	Slug        string         `json:"slug"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Trailer     string         `json:"trailer"`
	Date        string         `json:"date"`
	Duration    int            `json:"duration"`
	Poster      string         `json:"poster"`
	Background  imageDTO       `json:"background"`
	Site        siteRefDTO     `json:"site"`
	Performers  []performerDTO `json:"performers"`
	Directors   []directorDTO  `json:"directors"`
	Tags        []tagDTO       `json:"tags"`
}

type siteDTO struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
	Poster      string `json:"poster"`
	Network     string `json:"network"`
	Date        string `json:"date"`
	LastScene   string `json:"last_scene_date"`
}

type siteRefDTO struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	URL       string `json:"url"`
	Network   string `json:"network"`
	NetworkID int    `json:"network_id"`
	ParentID  int    `json:"parent_id"`
}

type performerDTO struct {
	ID         string   `json:"id"`
	Slug       string   `json:"slug"`
	Name       string   `json:"name"`
	Bio        string   `json:"bio"`
	Image      string   `json:"image"`
	Posters    imageDTO `json:"posters"`
	BirthDate  string   `json:"birthday"`
	DeathDate  string   `json:"death_date"`
	Birthplace string   `json:"birthplace"`
	Homepage   string   `json:"homepage"`
}

type directorDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type tagDTO struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type imageDTO struct {
	URL    string `json:"url"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Small  string `json:"small"`
}

type listResponse[T any] struct {
	Data []T `json:"data"`
}

type itemResponse[T any] struct {
	Data T `json:"data"`
}
