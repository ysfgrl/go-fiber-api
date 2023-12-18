package helpers

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"go-fiber-api/src/config"
	response2 "go-fiber-api/src/models"
	"go-fiber-api/src/models/elastic_collections"
	"go-fiber-api/src/utils/response"
)

type elasticHelper struct {
	BaseHelper[elasticsearch.TypedClient]
}

func NewElasticHelper(ctx context.Context) (*elasticHelper, *response2.MyError) {

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{
			config.Elastic.Url,
		},
	})
	if err != nil {
		return nil, response.GetError(err)
	}
	return &elasticHelper{
		BaseHelper[elasticsearch.TypedClient]{
			client: client,
			ctx:    context.TODO(),
		},
	}, nil
}

func (helper *elasticHelper) SendStruck(index string, document any) *response2.MyError {
	_, err := helper.client.Index(index).Request(document).Do(helper.ctx)
	if err != nil {
		return response.GetError(err)
	}
	return nil
}
func (helper *elasticHelper) SendTaskState(index string, document elastic_collections.TaskState) *response2.MyError {
	_, err := helper.client.Index(index).Id(document.TaskUUID).Request(document).Do(helper.ctx)
	if err != nil {
		return response.GetError(err)
	}
	return nil
}

func (helper *elasticHelper) UpdateTaskState(index string, document elastic_collections.TaskState) *response2.MyError {
	//return helper.SendTaskState(index, document)
	jj, _ := json.Marshal(document)
	rawMessage := json.RawMessage(jj)
	_, err := helper.client.Update(index, document.TaskUUID).Request(&update.Request{
		Doc: rawMessage,
	}).Do(helper.ctx)
	if err != nil {
		return response.GetError(err)
	}
	return nil
}
