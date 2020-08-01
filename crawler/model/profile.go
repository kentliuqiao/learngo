package model

import (
	"encoding/json"
)

// Profile user profile
type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Education  string
	Marriage   string
	Occupation string
	Hukou      string
	Xinzuo     string
	House      string
	Car        string
}

// FromJSONObj 将json反序列化得到的map转成Profile
func FromJSONObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
