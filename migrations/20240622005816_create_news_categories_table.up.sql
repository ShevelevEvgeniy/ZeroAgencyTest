CREATE TABLE news_categories (
                                 "news_id" BIGINT NOT NULL,
                                 "category_id" BIGINT NOT NULL,
                                 PRIMARY KEY ("news_id", "category_id"),
                                 CONSTRAINT "fk_news_id" FOREIGN KEY ("news_id") REFERENCES news ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX idx_news_id ON news_categories ("news_id");