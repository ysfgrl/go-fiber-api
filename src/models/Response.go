package models

import "github.com/ysfgrl/go-fiber-api/src/repository"

type Response struct {
	Code    int      `json:"code"`
	Content any      `json:"content"`
	Error   []*Error `json:"error"`
}

type ListResponse[CType repository.ElasticCollections | repository.MongoCollections] struct {
	List     []CType `json:"list"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
	Total    int     `json:"total"`
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
