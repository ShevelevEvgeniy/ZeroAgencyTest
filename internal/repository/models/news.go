//go:generate reform

package models

// reform:news
type News struct {
	Id      int64   `reform:"id,pk"`
	Title   *string `reform:"title"`
	Content *string `reform:"content"`
}
