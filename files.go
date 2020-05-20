package goscripts

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSONFile(filename string) map[string]interface{} {
	r := map[string]interface{}{}
	raw, err := ioutil.ReadFile(filename)
	Check(err)
	Check(json.Unmarshal(raw, &r))
	return r
}

func ReadJSONFileIfExists(filename string) map[string]interface{} {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return map[string]interface{}{}
	}
	return ReadJSONFile(filename)
}

func WriteJSONFile(filename string, j interface{}) {
	raw, err := json.MarshalIndent(j, "", "    ")
	Check(err)
	Check(ioutil.WriteFile(filename, raw, 0644))
}

func ReadTextFile(filename string) string {
	raw, err := ioutil.ReadFile(filename)
	Check(err)
	return string(raw)
}

func ReadTextFileIfExists(filename string) string {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return ""
	}
	return ReadTextFile(filename)
}

func WriteTextFile(filename string, txt string) {
	err := ioutil.WriteFile(filename, []byte(txt), 0644)
	Check(err)
}
