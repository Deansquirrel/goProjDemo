package object

import "encoding/json"

type SysConfig struct {
	Total total `toml:"total"`
}

type total struct {
	IsDebug bool `toml:"isDebug"`
}

func (sc *SysConfig) GetConfigStr() (string, error) {
	sConfig, err := json.Marshal(sc)
	if err != nil {
		return "", err
	}
	return string(sConfig), nil
}
