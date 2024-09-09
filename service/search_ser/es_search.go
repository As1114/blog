package search_ser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/axis1114/blog/global"
	"github.com/axis1114/blog/models"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"go.uber.org/zap"
)

func SearchDocumentTerm(field string, key string) (result []models.Article) {
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			Term: map[string]types.TermQuery{
				field: {Value: key},
			},
		}).
		Do(context.Background())
	if err != nil {
		global.Log.Error("search document failed, err:", zap.Error(err))
		return
	}
	for _, hit := range resp.Hits.Hits {
		var article models.Article
		err := json.Unmarshal(hit.Source_, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}
		result = append(result, article)
	}
	return result
}

func SearchDocumentTerms(field string, key []string) (result []models.Article) {
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			Terms: &types.TermsQuery{
				TermsQuery: map[string]types.TermsQueryField{
					field: key,
				},
			},
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("search document failed, err:%v\n", err)
		return
	}
	for _, hit := range resp.Hits.Hits {
		var article models.Article
		err := json.Unmarshal(hit.Source_, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}

		result = append(result, article)
	}
	return result
}

func SearchDocumentMultiMatch(fields []string, key string, pageInfo models.PageInfo) (result []models.Article) {
	form := (pageInfo.Page - 1) * pageInfo.PageSize
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Fields: fields,
				Query:  key,
			},
		}).From(form).Size(pageInfo.PageSize).
		Do(context.Background())
	if err != nil {
		global.Log.Error("search document failed", zap.Error(err))
		return
	}
	for _, hit := range resp.Hits.Hits {
		var article models.Article
		err := json.Unmarshal(hit.Source_, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}
		result = append(result, article)
	}
	return result
}

func GetDocumentById(id string) (result models.Article) {
	resp, err := global.Es.Get(models.Article{}.Index(), id).
		Do(context.Background())
	if err != nil {
		global.Log.Error("get document by id failed", zap.Error(err))
		return
	}
	var article models.Article
	err = json.Unmarshal(resp.Source_, &article)
	if err != nil {
		global.Log.Error("unmarshal json failed", zap.Error(err))
		return
	}

	return result
}

func SearchAllDocuments() (result []models.Article) {
	resp, err := global.Es.Search().
		Index(models.Article{}.Index()).
		Query(&types.Query{
			MatchAll: &types.MatchAllQuery{},
		}).Do(context.Background())
	if err != nil {
		global.Log.Error("search all documents failed", zap.Error(err))
		return
	}
	for _, hit := range resp.Hits.Hits {
		var article models.Article
		err := json.Unmarshal(hit.Source_, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			continue
		}
		result = append(result, article)
	}
	return result
}

func DocIsExistByTitle(title string) bool {
	res := SearchDocumentTerm("title.keyword", title)
	if len(res) == 0 {
		return false
	} else {
		return true
	}
}

func DocIsExistById(id string) bool {
	res, _ := global.Es.Exists(models.Article{}.Index(), id).Do(context.Background())
	return res
}
