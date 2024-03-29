package model

import "encoding/json"

// 用户模型，用于存储用户的各种信息
// 因为这个struct可能不仅仅用于珍爱网，所以不定义在zhenai文件夹下

type Profile struct {
	Name string
	Gender string
	Age int
	Height int
	Weight int
	Income string
	Marriage string
	Education string
	AncestralHome string // 籍贯
	Constellation string // 星座
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}