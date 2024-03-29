package clients

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ysfgrl/go-fiber-api/src/config"
)

func initElastic(ctx context.Context) (*elasticsearch.TypedClient, error) {

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{
			config.AppConf.Elastic.Url,
		},
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

//func (helper *elasticHelper) SendStruck(index string, document any) *response2.Error {
//	_, err := helper.client.Index(index).Request(document).Do(helper.ctx)
//	if err != nil {
//		return response.GetError(err)
//	}
//	return nil
//}
//func (helper *elasticHelper) SendTaskState(index string, document elastic_collections.TaskState) *response2.Error {
//	_, err := helper.client.Index(index).Id(document.TaskUUID).Request(document).Do(helper.ctx)
//	if err != nil {
//		return response.GetError(err)
//	}
//	return nil
//}
//
//func (helper *elasticHelper) UpdateTaskState(index string, document elastic_collections.TaskState) *response2.Error {
//	//return helper.SendTaskState(index, document)
//	jj, _ := json.Marshal(document)
//	rawMessage := json.RawMessage(jj)
//	_, err := helper.client.Update(index, document.TaskUUID).Request(&update.Request{
//		Doc: rawMessage,
//	}).Do(helper.ctx)
//	if err != nil {
//		return response.GetError(err)
//	}
//	return nil
//}
