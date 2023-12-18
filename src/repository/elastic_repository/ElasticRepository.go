package elastic_repository

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"go-fiber-api/src/models"
	"go-fiber-api/src/utils/response"
	"time"
)

type ElasticRepository[CType models.ElasticCollections] struct {
	client *elasticsearch.TypedClient
	ctx    context.Context
	index  string
}

func (repo *ElasticRepository[CType]) Get(id string) (*CType, *models.MyError) {
	res, err := repo.client.Get(repo.index, id).Do(repo.ctx)
	if err != nil {
		return nil, response.GetError(err)
	}
	if res.Found {
		var cType *CType
		err := json.Unmarshal(res.Source_, &cType)
		if err != nil {
			return nil, response.GetError(err)
		}
		return cType, nil
	}
	return nil, nil
}

func (repo *ElasticRepository[CType]) GetByFirst(key string, value any) (*CType, *models.MyError) {
	//TODO implement me
	panic("implement me")
}

func (repo *ElasticRepository[CType]) List(schema models.ListRequest) (*models.ListResponse[CType], *models.MyError) {
	schema.Page -= 1
	gte := schema.Gte.Format(time.RFC3339)
	lte := schema.Lte.Format(time.RFC3339)
	rang := map[string]types.RangeQuery{
		"CreatedAt": types.DateRangeQuery{
			Gte: &gte,
			Lte: &lte,
		},
	}
	match := map[string]types.MatchQuery{
		"TaskName": {
			Query: "test",
			Operator: &operator.Operator{
				Name: "or",
			},
		},
	}
	print(rang)
	print(match)
	query := repo.client.Search().
		Index("task_states").
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Range: rang,
						},
					},
					Should: []types.Query{
						{
							Term: map[string]types.TermQuery{
								"TaskName": {
									Value: schema.Keyword,
								},
							},
						},
					},
				},
			},
			From: &schema.Page,
			Size: &schema.PageSize,
		})
	res, err := query.Do(context.TODO())
	if err != nil {
		return nil, response.GetError(err)
	}

	list := []CType{}
	for _, hit := range res.Hits.Hits {
		var cType CType
		err := json.Unmarshal(hit.Source_, &cType)
		if err != nil {
			return nil, response.GetError(err)
		}
		list = append(list, cType)
	}

	return &models.ListResponse[CType]{
		Total:    int(res.Hits.Total.Value),
		PageSize: schema.PageSize,
		Page:     schema.Page + 1,
		List:     list,
	}, nil
}

func (repo *ElasticRepository[CType]) Add(schema CType) (*CType, *models.MyError) {
	//TODO implement me
	panic("implement me")
}

func (repo *ElasticRepository[CType]) Delete(id string) (bool, *models.MyError) {
	//TODO implement me
	panic("implement me")
}

func (repo *ElasticRepository[CType]) Update(id string, schema CType) (*CType, *models.MyError) {
	//TODO implement me
	panic("implement me")
}

func (repo *ElasticRepository[CType]) UpdateField(id string, field string, value any) (*CType, *models.MyError) {
	//TODO implement me
	panic("implement me")
}
