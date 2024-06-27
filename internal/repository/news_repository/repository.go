package news_repository

import (
	newsConverter "ZeroAgencyTest/internal/converter/news_converter"
	"ZeroAgencyTest/internal/repository/models"
	repositoryQuery "ZeroAgencyTest/internal/repository/repository_query"
	servModel "ZeroAgencyTest/internal/service/news_service/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

type NewsRepository struct {
	db *reform.DB
}

func NewNewsRepository(db *reform.DB) *NewsRepository {
	return &NewsRepository{
		db: db,
	}
}

func (r *NewsRepository) GetNews(ctx context.Context) ([]servModel.News, error) {
	rows, err := r.db.QueryContext(ctx, repositoryQuery.GetNewsList)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching news")
	}
	defer rows.Close()

	var newsList []models.NewsWithCategories

	for rows.Next() {
		news := models.NewsWithCategories{}

		err = rows.Scan(&news.Id, &news.Title, &news.Content, &news.Categories)
		if err != nil {
			return nil, errors.Wrap(err, "Error scanning news")
		}

		newsList = append(newsList, news)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "Error reading rows")
	}

	return newsConverter.RepoToServModels(newsList), nil
}

func (r *NewsRepository) UpdateNews(ctx context.Context, news *models.News, newsCategories *models.NewsCategories) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	updateColumns := repositoryQuery.ColumnsToUpdate(news)

	if len(updateColumns) != 0 {
		if err = tx.UpdateColumns(news, updateColumns...); err != nil {
			if errs := tx.Rollback(); errs != nil {
				return errors.Wrap(errs, "failed to rollback transaction")
			}

			return errors.Wrap(err, "failed to update news")
		}
	}

	if newsCategories != nil {
		fmt.Println(newsCategories)
		if err = tx.Save(newsCategories); err != nil {
			if err = tx.Rollback(); err != nil {
				return errors.Wrap(err, "failed to rollback transaction")
			}
			return errors.Wrap(err, "failed to update news_categories")
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}
