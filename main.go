package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type repos []struct {
	RepoName     string `json:"name"`
	LanguageName string `json:"language"`
}

func main() {
	id := "eprado99"
	var r repos
	// Github API docs: https://docs.github.com/en/rest?apiVersion=2022-11-28

	res, err := http.Get("https://api.github.com/users/" + id + "/repos")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &r); err != nil {
		fmt.Println(err)
		return
	}

	languageCount := map[string]int{}
	for _, repo := range r {
		languageCount[repo.LanguageName]++
	}

	languagePercent := map[string]float32{}
	reposCount := len(r)
	for language, count := range languageCount {
		languagePercent[language] = float32(count) / float32(reposCount) * 100
	}

	for lang, perc := range languagePercent {
		fmt.Printf("%s: %.2f%%\n", lang, perc)
	}

}
