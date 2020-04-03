package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

// BigHuge ...
type BigHuge struct {
	APIKey string //1ad2d61d8ab6a8f8b8fd61ba200f6ca7
}
type synonyms struct {
	Adjective *words `json:"adjective"`
	Adverb    *words `json:"adverb"`
}
type words struct {
	Syn []string `json:"syn"`
}

// Synonyms ...
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string

	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, errors.New("bighuge: Failed when looking for  synonym for " + term + "" + err.Error())
	}

	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Adjective != nil {
		syns = append(syns, data.Adjective.Syn...)
	}
	if data.Adverb != nil {
		syns = append(syns, data.Adverb.Syn...)
	}
	return syns, nil
}
