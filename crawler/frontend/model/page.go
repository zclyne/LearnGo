package model

import "learngo.com/crawler/engine"

type SearchResult struct {
	Hits int64
	Start int
	Query string
	PrevFrom int
	NextFrom int
	Items []engine.Item
}