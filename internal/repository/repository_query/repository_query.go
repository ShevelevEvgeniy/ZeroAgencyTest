package repository_query

import _ "embed"

var (
	//go:embed news/get_news_list.sql
	GetNewsList string
)
