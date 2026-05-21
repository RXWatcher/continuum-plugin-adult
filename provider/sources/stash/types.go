package stash

// Subset of the Stash GraphQL schema we care about. Field names mirror the
// upstream camelCase / snake_case as exposed by Stash. Unused fields are
// intentionally omitted to keep the surface tight.

type sceneDTO struct {
	ID         string         `json:"id"`
	Title      string         `json:"title"`
	Details    string         `json:"details"`
	Date       string         `json:"date"`
	Paths      scenePathsDTO  `json:"paths"`
	Studio     *studioDTO     `json:"studio"`
	Tags       []tagDTO       `json:"tags"`
	Performers []performerDTO `json:"performers"`
}

type scenePathsDTO struct {
	Screenshot string `json:"screenshot"`
}

type studioDTO struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	ImagePath     string     `json:"image_path"`
	Details       string     `json:"details"`
	ParentStudio  *studioDTO `json:"parent_studio"`
}

type performerDTO struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Disambiguation string   `json:"disambiguation"`
	Details        string   `json:"details"`
	ImagePath      string   `json:"image_path"`
	Birthdate      string   `json:"birthdate"`
	DeathDate      string   `json:"death_date"`
	Country        string   `json:"country"`
	AliasList      []string `json:"alias_list"`
	Tags           []tagDTO `json:"tags"`
}

type tagDTO struct {
	Name string `json:"name"`
}

// GraphQL response envelopes.

type findScenesResp struct {
	FindScenes struct {
		Scenes []sceneDTO `json:"scenes"`
	} `json:"findScenes"`
}

type findSceneResp struct {
	FindScene *sceneDTO `json:"findScene"`
}

type findStudiosResp struct {
	FindStudios struct {
		Studios []studioDTO `json:"studios"`
	} `json:"findStudios"`
}

type findStudioResp struct {
	FindStudio *studioDTO `json:"findStudio"`
}

type findPerformerResp struct {
	FindPerformer *performerDTO `json:"findPerformer"`
}
