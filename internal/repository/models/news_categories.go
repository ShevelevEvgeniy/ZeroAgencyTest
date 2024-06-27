//go:generate reform

package models

// reform:news_categories
type NewsCategories struct {
	NewsId     int64  `reform:"news_id,pk"`
	CategoryId *int64 `reform:"category_id"`
}
