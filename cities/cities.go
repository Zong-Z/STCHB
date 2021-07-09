package cities

import (
	"encoding/json"
	"io/ioutil"
	"stchb/logger"
)

const File = "configs/cities.json"

var Cities []string

func init() {
	type Cities_ struct {
		Cities []string
	}

	b, err := ioutil.ReadFile(File)
	if err != nil {
		logger.ForWarning(err.Error())
	}

	var cities_ Cities_
	if err := json.Unmarshal(b, &cities_); err != nil {
		logger.ForWarning(err.Error())
	}

	Cities = cities_.Cities
}
