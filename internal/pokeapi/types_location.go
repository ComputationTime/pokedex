package pokeapi

type LocationAreasResp struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Resource `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int64                 `json:"game_index"`
	ID                   int64                 `json:"id"`
	Location             Resource              `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod Resource                           `json:"encounter_method"`
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"`
}

type Resource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRateVersionDetail struct {
	Rate    int64    `json:"rate"`
	Version Resource `json:"version"`
}

type Name struct {
	Language Resource `json:"language"`
	Name     string   `json:"name"`
}

type PokemonEncounter struct {
	Pokemon        Resource                        `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	MaxChance        int64             `json:"max_chance"`
	Version          Resource          `json:"version"`
}

type EncounterDetail struct {
	Chance          int64         `json:"chance"`
	ConditionValues []interface{} `json:"condition_values"`
	MaxLevel        int64         `json:"max_level"`
	Method          Resource      `json:"method"`
	MinLevel        int64         `json:"min_level"`
}
