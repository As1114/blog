package models

import (
	"blog/global"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"go.uber.org/zap"
)

type ArticleItem struct {
	ID      string `json:"id"`
	Article Article
}
type Article struct {
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间

	Title    string `json:"title"`    // 文章标题
	Abstract string `json:"abstract"` // 文章简介
	Content  string `json:"content" ` // 文章内容

	LookCount     int `json:"look_count" `    // 浏览量
	CommentCount  int `json:"comment_count" ` // 评论量
	DiggCount     int `json:"digg_count"`     // 点赞量
	CollectsCount int `json:"collects_count"` // 收藏量

	UserID     uint   `json:"user_id" `     // 用户id
	UserName   string `json:"user_name" `   //用户昵称
	UserAvatar string `json:"user_avatar" ` // 用户头像

	Category string `json:"category"` // 文章分类

	CoverID  uint   `json:"banner_id"`  // 封面id
	CoverURL string `json:"banner_url"` // 封面
}

func (a ArticleItem) Index() string {
	return "article_index"
}

func (a ArticleItem) CreateIndex() {
	exist := a.IndexExist()
	if exist {
		global.Log.Info("the index already exists")
		return
	}
	resp, err := global.Es.Indices.
		Create(a.Index()).
		Do(context.Background())
	if err != nil {
		global.Log.Error("create the index failed, err:%v\n", zap.Error(err))
		return
	}
	global.Log.Info("create the index successfully", zap.Any("index", resp))
}

func (a ArticleItem) CreateIndexByJson(index string) {
	exist := a.IndexExistByJson(index)
	if exist {
		global.Log.Info("the index already exists")
		return
	}
	resp, err := global.Es.Indices.
		Create(index).
		Do(context.Background())
	if err != nil {
		global.Log.Error("create the index failed, err:%v\n", zap.Error(err))
		return
	}
	global.Log.Info("create the index successfully", zap.Any("index", resp))
}

func (a ArticleItem) IndexExist() bool {
	resp, err := global.Es.Indices.Exists(a.Index()).Do(context.Background())
	if err != nil {
		global.Log.Error("detect the presence of the index", zap.Error(err))
	}
	return resp
}

func (a ArticleItem) IndexExistByJson(index string) bool {
	resp, err := global.Es.Indices.Exists(index).Do(context.Background())
	if err != nil {
		global.Log.Error("detect the presence of the index", zap.Error(err))
	}
	return resp
}

func (a ArticleItem) DocumentExist(title string) bool {
	res := a.SearchDocumentTerm("title", title)
	if len(res) == 0 {
		return false
	} else {
		return true
	}
}

func (a ArticleItem) DeleteIndex() {
	resp, err := global.Es.Indices.
		Delete(a.Index()).
		Do(context.Background())
	if err != nil {
		global.Log.Error("failed to delete the index, err", zap.Error(err))
		return
	}
	global.Log.Info("succeed to delete the index", zap.Any("delete", resp))
}

func (a ArticleItem) CreateDocument() {
	resp, err := global.Es.Index(a.Index()).Document(a).Do(context.Background())
	if err != nil {
		global.Log.Error("failed to create the document", zap.Error(err))
		return
	}
	global.Log.Info("succeed to create the document", zap.Any("doc", resp))
}

func (a ArticleItem) DeleteDocument() {
	resp, err := global.Es.Delete(a.Index(), "TWtA_JAB9Xws_a1XiaYk").Do(context.Background())
	if err != nil {
		global.Log.Error("delete document failed, err:%v\n", zap.Error(err))
		return
	}
	global.Log.Info("succeed to delete the document", zap.Any("delete", resp))
}

func (a ArticleItem) GetDocumentById() (result ArticleItem) {
	resp, err := global.Es.Get(a.Index(), a.ID).
		Do(context.Background())
	if err != nil {
		global.Log.Error("get document by id failed", zap.Error(err))
		return
	}
	var article Article
	var articleitem ArticleItem
	data := string(resp.Source_)
	bytes := []byte(data)
	err = json.Unmarshal(bytes, &article)
	if err != nil {
		global.Log.Error("unmarshal json failed", zap.Error(err))
		return
	}
	articleitem.Article = article
	articleitem.ID = resp.Id_
	return result
}

func (a ArticleItem) SearchAllDocuments() (result []ArticleItem) {
	resp, err := global.Es.Search().
		Index(a.Index()).
		Query(&types.Query{
			MatchAll: &types.MatchAllQuery{},
		}).Do(context.Background())
	if err != nil {
		global.Log.Error("search all documents failed", zap.Error(err))
		return
	}
	var article Article
	var articleitem ArticleItem
	for _, hit := range resp.Hits.Hits {
		data := string(hit.Source_)
		bytes := []byte(data)
		err := json.Unmarshal(bytes, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			return nil
		}
		articleitem.Article = article
		articleitem.ID = *hit.Id_
		result = append(result, articleitem)
	}
	return result
}

func (a ArticleItem) SearchDocumentMultiMatch(fields []string, key string, pageInfo PageInfo) (result []ArticleItem) {
	form := (pageInfo.Page - 1) * pageInfo.Limit
	resp, err := global.Es.Search().
		Index(a.Index()).
		Query(&types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Fields: fields,
				Query:  key,
			},
		}).From(form).Size(pageInfo.Limit).
		Do(context.Background())
	if err != nil {
		global.Log.Error("search document failed", zap.Error(err))
		return
	}
	var article Article
	var articleitem ArticleItem
	for _, hit := range resp.Hits.Hits {
		data := string(hit.Source_)
		bytes := []byte(data)
		err := json.Unmarshal(bytes, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			return nil
		}
		articleitem.Article = article
		articleitem.ID = *hit.Id_
		result = append(result, articleitem)
	}
	return result
}

func (a ArticleItem) SearchDocumentTerms(field string, key []string) (result []ArticleItem) {
	resp, err := global.Es.Search().
		Index(a.Index()).
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
	var article Article
	var articleitem ArticleItem
	for _, hit := range resp.Hits.Hits {
		data := string(hit.Source_)
		bytes := []byte(data)
		err := json.Unmarshal(bytes, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			return nil
		}
		articleitem.Article = article
		articleitem.ID = *hit.Id_
		result = append(result, articleitem)
	}
	return result
}

func (a ArticleItem) SearchDocumentTerm(field string, key string) (result []ArticleItem) {
	resp, err := global.Es.Search().
		Index(a.Index()).
		Query(&types.Query{
			Term: map[string]types.TermQuery{
				field: {Value: key},
			},
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("search document failed, err:%v\n", err)
		return
	}
	var article Article
	var articleitem ArticleItem
	for _, hit := range resp.Hits.Hits {
		data := string(hit.Source_)
		bytes := []byte(data)
		err := json.Unmarshal(bytes, &article)
		if err != nil {
			global.Log.Error("unmarshal json failed", zap.Error(err))
			return
		}
		articleitem.Article = article
		articleitem.ID = *hit.Id_
		result = append(result, articleitem)
	}
	return result
}

func (a ArticleItem) UpdateDocument() {
	resp, err := global.Es.Update(a.Index(), a.ID).Doc(a.Article).Do(context.Background())
	if err != nil {
		global.Log.Error("update document failed, err:%v\n", zap.Error(err))
		return
	}
	global.Log.Info("succeed to update the document", zap.Any("update", resp))
}

func (a ArticleItem) deleteDocument() {
	resp, err := global.Es.Delete(a.Index(), a.ID).
		Do(context.Background())
	if err != nil {
		global.Log.Error("delete document failed", zap.Error(err))
		return
	}
	global.Log.Info("succeed to delete the document", zap.Any("delete", resp))
}

func (a ArticleItem) DeleteMultipleDocuments(ids []string) error {
	bulkRequest := global.Es.Bulk()

	for _, id := range ids {
		op := types.DeleteOperation{Id_: &id}
		err := bulkRequest.DeleteOp(op)
		if err != nil {
			global.Log.Error("delete document failed", zap.Error(err))
			return err
		}
	}

	_, err := bulkRequest.Do(context.Background())
	if err != nil {
		global.Log.Error("bulk delete documents failed", zap.Error(err))
		return err
	}
	global.Log.Info("succeed to bulk delete documents")
	return nil
}
