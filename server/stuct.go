package main

type Suggestion struct {
	Id int64 `json:"id"`
	Type int64 `json:"type"`
	Time int64 `json:"time"`
	Content string `json:"content"`
}
