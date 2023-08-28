package database

import (
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
)

var languages = make(map[string]string)

func GetLanguages() map[string]string {
	return languages
}

func setLanguages() map[string]string {
	raw := tools.GetJSONRawMessage(filepath.Join(localesPath, "/languages.json"))
	err := json.Unmarshal(raw, &languages)
	if err != nil {
		panic(err)
	}
	return languages
}