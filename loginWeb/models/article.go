package models

import (
	"go.uber.org/zap"
	"netgo/loginWeb/dao"
	"netgo/loginWeb/logger"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime string `db:"create_time"`
	Status     int    //0正常，1删除，冻结
}

const (
	pageSize = 4 //每页有4条数据
)

//---数据处理---
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

func insertArticle(article Article) (int64, error) {
	logger.Debug("insertArticle", zap.Any("article", article))
	//logger.Debug("insertArticle", zap.Any("article",article))
	return dao.ModifyDB("insert into z_article(title,tags,short,content,author,create_time,status) values (?,?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime, article.Status)
}

//查询所有文章

func QueryArticleWithPage(pageNum int) (articleList []*Article, err error) {

	sqlStr := "select id,title,tags,short,author,create_time from  z_article limit ?,?"

	articleList, err = QueryArticleWithCon(pageNum, sqlStr)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func QueryCurrUserArticleWithPage(username string, pageNum int) (articleList []*Article, err error) {
	sqlStr := "select id,title,tags,short,content,author,create_time from z_article where author=? limit ?,?"
	articleList, err = QueryArticleWithCon(pageNum, sqlStr, username)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

func QueryArticleWithCon(pageNum int, sqlStr string, args ...interface{}) (articleList []*Article, err error) {
	pageNum--
	args = append(args, pageNum*pageSize, pageSize)
	err = dao.QueryRows(&articleList, sqlStr, args...)
	return
}
