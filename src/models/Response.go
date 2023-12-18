package models

import (
	"encoding/json"
)

type Response struct {
	Code    int        `json:"code"`
	Content any        `json:"content"`
	Error   []*MyError `json:"error"`
}

type ListResponse[CType ElasticCollections | MongoCollections] struct {
	List     []CType `json:"list"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
	Total    int     `json:"total"`
}

type MyError struct {
	File     string `json:"file"`
	Function string `json:"function"`
	Detail   any    `json:"detail"`
	Line     int    `json:"line"`
	Code     string `json:"code"`
}

func (e *MyError) ToJson() []byte {
	b, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		return nil
	}
	return b
}

type ValidationError struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Tag   string      `json:"tag"`
	Param string      `json:"param"`
}

type UploadedFile struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Type   string `json:"type"`
}
