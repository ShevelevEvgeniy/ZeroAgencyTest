package models

import "github.com/lib/pq"

type NewsWithCategories struct {
	Id         int64         `reform:"id,pk"`
	Title      string        `reform:"title"`
	Content    string        `reform:"content"`
	Categories pq.Int64Array `reform:"categories" json:"categories"`
}
