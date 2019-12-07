package models

import "time"

var stringGen = NewRandomStringGenerator("0123456789abcdefghijklmnopqrstuvwxyz")

// RGroup - structure of requests group storing.
type RGroup struct {
	Tag        		string `json:"tag"`
	Created     	int64  `json:"created"`
	Updated      	int64  `json:"updated"`
	Volume 			int    `json:"volume"`
}

// CreateRGroup - RGroup constructor
func CreateRGroup() *RGroup {
	now := time.Now().Unix()
	RGroup := RGroup{
		Tag:		stringGen.Generate(8),
		Created:	now,
		Updated:	now,
	}
	return &RGroup
}