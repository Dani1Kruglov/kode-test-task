package speller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const BaseApiURL = "https://speller.yandex.net/services/spellservice.json/checkText?text="

type Speller struct {
	Code         int      `json:"code"`
	Pos          int      `json:"pos"`
	Row          int      `json:"row"`
	Col          int      `json:"col"`
	Len          int      `json:"len"`
	Word         string   `json:"word"`
	SpellerWords []string `json:"s"`
}

func GetCorrectedTextUsingSpeller(text string) (string, error) {
	regex := regexp.MustCompile(`\p{L}+`)
	words := regex.FindAllString(text, -1)

	apiUrlUsed := BaseApiURL
	apiUrlUsed = apiUrlUsed + words[0]
	for i, word := range words {
		if i != 0 {
			apiUrlUsed = apiUrlUsed + "+" + word
		}
	}

	res, err := http.Get(apiUrlUsed)
	if err != nil {
		return "", err
	}
	bodyURL, err := io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)

	var spellersJSON []Speller
	err = json.Unmarshal(bodyURL, &spellersJSON)
	if err != nil {
		return "", err
	}

	for _, spellerJSON := range spellersJSON {
		text = strings.Replace(text, spellerJSON.Word, spellerJSON.SpellerWords[0], 1)
	}
	return text, nil
}
