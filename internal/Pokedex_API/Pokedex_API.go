package PokeDex_API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	PokeCache "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/PokeCache"
)

// Declare config struct
type Config struct {
	Initial  string
	Next     string
	Previous string
	Pokedex  []string
	Cache    *PokeCache.Cache
}

// Declare API struct
type LocationAreaResponse struct {
	Next     string
	Previous string
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       any    `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  any    `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      any    `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale any    `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string `json:"back_default"`
				BackFemale       any    `json:"back_female"`
				BackShiny        string `json:"back_shiny"`
				BackShinyFemale  any    `json:"back_shiny_female"`
				FrontDefault     string `json:"front_default"`
				FrontFemale      any    `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale any    `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       any    `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  any    `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      any    `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale any    `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       any    `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  any    `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      any    `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale any    `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
}

type ThreadSafeCache struct {
	data map[string][]byte
	mu   sync.Mutex
}

func NewThreadSafeCache() *ThreadSafeCache {
	return &ThreadSafeCache{
		data: make(map[string][]byte),
	}
}

// Adds a new entry to the cache
func (c *ThreadSafeCache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Gets an entry from the cache if it exists
func (c *ThreadSafeCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val, ok
}

func GetNextLocations(cfg *Config) ([]string, error) {
	var url string
	// Determine the URL: is it the first call or subsequent request?
	if cfg.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = cfg.Next
	}

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var names []string
		if err := json.Unmarshal(val, &names); err != nil {
			return nil, err
		}
		return names, nil
	}
	// If the data is not in the cache, make the API call
	response := LocationAreaResponse{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Check if HTTP response is successful (status code 200)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	// Update the config with the next and previous URLs from this response
	cfg.Next = response.Next
	cfg.Previous = response.Previous

	names := []string{}
	for _, loc := range response.Results {
		names = append(names, loc.Name)
	}
	// Store the data in the cache
	cachedValue, err := json.Marshal(names)
	if err != nil {
		return nil, err
	}
	cfg.Cache.Add(url, cachedValue)

	// Return the data
	return names, nil
}

func GetPreviousLocations(cfg *Config) ([]string, error) {
	var url string // IF no URL is stored yet (first call), use the base URL
	if cfg.Previous == "" {
		return nil, fmt.Errorf("you're on the first page")
	} else {
		url = cfg.Previous
	}

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var names []string
		if err := json.Unmarshal(val, &names); err != nil {
			return nil, err
		}
		return names, nil
	}

	// If the data is not in the cache, make the API call
	response := LocationAreaResponse{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Check if HTTP response is successful (status code 200)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	// Update the config with the next and previous URLs from this response
	cfg.Next = response.Next
	cfg.Previous = response.Previous

	names := []string{}
	for _, loc := range response.Results {
		names = append(names, loc.Name)
	}

	// Store the data in the cache
	cachedValue, err := json.Marshal(names)
	if err != nil {
		return nil, err
	}
	cfg.Cache.Add(url, cachedValue)

	// Return the data
	return names, nil
}

func GetPokemon(cfg *Config, loc string) ([]string, error) {
	// Add the location to the base URL
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", loc)

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var pokemonName []string
		if err := json.Unmarshal(val, &pokemonName); err != nil {
			return nil, err
		}
		return pokemonName, nil
	}

	// If the data is not in the cache, make the API call
	response := LocationAreaResponse{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Check if HTTP response is successful (status code 200)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	pokemonNames := []string{}
	for _, pokemon := range response.PokemonEncounters {
		pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
	}

	// Store the data in the cache
	cachedValue, err := json.Marshal(pokemonNames)
	if err != nil {
		return nil, err
	}
	cfg.Cache.Add(url, cachedValue)

	// Return the data
	return pokemonNames, nil
}

func CatchPokemon(cfg *Config, pokemon string) (int, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var response PokemonResponse // Unmarshal into the correct struct
		if len(val) == 0 {           // Validate empty data
			return 0, fmt.Errorf("cached data is empty for URL: %s", url)
		}
		// Unmarshal the cached JSON data into a PokemonResponse struct
		if err := json.Unmarshal(val, &response); err != nil {
			return 0, fmt.Errorf("failed to decode cached data: %w", err)
		}
		// Return the BaseExperience value
		return response.BaseExperience, nil
	}

	// If the data is not in the cache, make the API call
	response := PokemonResponse{}

	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	// Check if HTTP response is successful (status code 200)
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API request failed with status: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return 0, err
	}

	// Store the data in the cache
	responseData := PokemonResponse{
		Name:           pokemon,
		BaseExperience: response.BaseExperience,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		return 0, err
	}
	cfg.Cache.Add(url, jsonData)

	// Return the data
	return response.BaseExperience, nil
}
