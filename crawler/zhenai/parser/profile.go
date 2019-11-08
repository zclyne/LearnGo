package parser

import (
	"LearnGo/crawler/engine"
	"LearnGo/crawler/model"
	"regexp"
	"strconv"
)

// 用户页面的Parser，提取出用户的详细信息

// 匹配用的正则表达式，这里预先编译好，以节省时间
var ageRe = regexp.MustCompile(`<div[^>]*>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div[^>]*>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div[^>]*>([\d]+)kg</div>`)
var constellationRe = regexp.MustCompile(`<div[^>]*>([^>]+座)\([\d]+\.[\d]+-[\d]+\.[\d]+\)</div>`)
var ancestralHomeRe = regexp.MustCompile(`div[^>]*>籍贯:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div[^>]*>月收入:([^<]+)</div>`)
var educationRe = regexp.MustCompile(`<div[^>]*>(高中及以下|中专|大专|大学本科|硕士|博士)</div>`)
var marriageRe = regexp.MustCompile(`<div[^>]*>(未婚|离异|丧偶)</div>`)

// 第二、三个参数是从city的parser那边传来的用户的名称和性别
func ParseProfile(contents []byte, name string, gender string) engine.ParseResult {
	// 创建用户profile
	profile := model.Profile{
		Name:          "",
		Gender:        "",
		Age:           0,
		Height:        0,
		Weight:        0,
		Income:        "",
		Marriage:      "",
		Education:     "",
		AncestralHome: "",
		Constellation: "",
	}

	// 匹配年龄
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	// 匹配婚姻情况
	profile.Marriage = extractString(contents, marriageRe)
	// 匹配身高
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	// 匹配体重
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	// 匹配籍贯
	profile.AncestralHome = extractString(contents, ancestralHomeRe)
	// 匹配星座
	profile.Constellation = extractString(contents, constellationRe)
	// 匹配收入
	profile.Income = extractString(contents, incomeRe)
	// 匹配教育状况
	profile.Education = extractString(contents, educationRe)
	// 从参数中获得其他的用户信息
	profile.Name = name
	profile.Gender = gender

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

// 从正则表达式中获取匹配到的字符串并返回
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1]) // match[1]是匹配到的字符串对应的正则表达式括号内部的内容
	}
	return ""
}
