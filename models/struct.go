package models

type PokemonItem struct {
	ID    string
	Name  string
	Image string
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type TypeInfo `json:"type"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type Character struct {
	Id      int           `json:"id"`
	Name    string        `json:"name"`
	Sprites Sprites       `json:"sprites"`
	Types   []PokemonType `json:"types"`
	Stats   []PokemonStat `json:"stats"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type List struct {
	Results []Results `json:"results"`
}
