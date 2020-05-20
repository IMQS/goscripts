package goscripts

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJSONFile(filename string) map[string]interface{} {
	r := map[string]interface{}{}
	raw, err := ioutil.ReadFile(filename)
	Check(err)
	Check(json.Unmarshal(raw, &r))
	return r
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

func WriteTextFile(filename string, txt string) {
	err := ioutil.WriteFile(filename, []byte(txt), 0644)
	Check(err)
}
