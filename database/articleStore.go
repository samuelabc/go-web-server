package database

import (
	"context"
	"fmt"
	"strings"
	articleController "web-server/api/article"
	models "web-server/model"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// ProfileStore implements database operations for profile management.
type ArticleStore struct {
	db *pgxpool.Pool
}

// NewAccountStore returns an AccountStore.
func NewArticleStore(db *pgxpool.Pool) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

// Get an article by ID.
func (s *ArticleStore) Get(a *articleController.GetArticleRequest) (*models.Article, error) {
	var err error
	ctx := context.Background()
	row := s.db.QueryRow(ctx, `select *
		from articles
		where articles.id = $1
		limit 1`, *a.ID)
	var res models.Article
	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Get an article by filter.
func (s *ArticleStore) List(a *articleController.ListArticleRequest) (*[]models.Article, error) {
	var err error
	ctx := context.Background()

	var query string
	var filters []string
	query = "select * from articles"

	//filter
	if a.ID != nil || a.Title != nil || a.Content != nil {
		query = fmt.Sprint(query, " where")
	}
	if a.ID != nil {
		filters = append(filters, fmt.Sprint("articles.id = ", "'", *a.ID, "'"))
	}
	if a.Title != nil {
		filters = append(filters, fmt.Sprint("articles.title = ", "'", *a.Title, "'"))
	}
	if a.Content != nil {
		filters = append(filters, fmt.Sprint("articles.content = ", "'", *a.Content, "'"))
	}
	filterQuery := strings.Join(filters, " and ")
	query = fmt.Sprint(query, " ", filterQuery)

	//limit
	query = fmt.Sprint(query, " limit ", 100)

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := []models.Article{}
	for rows.Next() {
		var a models.Article
		err = rows.Scan(&a.ID, &a.Title, &a.Content, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &articles, err
}

// Create create a new article.
func (s *ArticleStore) Create(a *models.Article) (*models.Article, error) {
	var err error

	u, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	a.ID = u.String()

	tx, err := s.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}()

	var res models.Article
	row := tx.QueryRow(context.Background(), `
		INSERT INTO articles(id, title, content)
		VALUES ($1, $2, $3)
		RETURNING *;
	`, a.ID, a.Title, a.Content)
	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Update an article.
func (s *ArticleStore) Update(a *articleController.UpdateArticleRequest) (*models.Article, error) {
	var err error

	var query string
	var filters []string
	query = `UPDATE articles`

	//filter
	if a.Title != nil {
		filters = append(filters, fmt.Sprint("title = ", "'", *a.Title, "'"))
	}
	if a.Content != nil {
		filters = append(filters, fmt.Sprint("content = ", "'", *a.Content, "'"))
	}
	if len(filters) > 0 {
		filterQuery := strings.Join(filters, " , ")
		query = fmt.Sprint(query, " SET ", filterQuery)
	}
	query = fmt.Sprint(query, " WHERE articles.ID = ", "'", *a.ID, "'")

	//limit
	// query = fmt.Sprint(query, " limit ", 100)
	query = fmt.Sprint(query, " returning *")

	tx, err := s.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		} else {
			tx.Commit(context.Background())
		}
	}()

	var res models.Article
	row := tx.QueryRow(context.Background(), query)
	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Delete an account.
func (s *ArticleStore) Delete(a *articleController.DeleteArticleRequest) (*models.Article, error) {
	var err error
	ctx := context.Background()
	row := s.db.QueryRow(ctx, `
		DELETE FROM articles
		WHERE articles.id = $1
		RETURNING *
	`, *a.ID)
	var res models.Article
	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, err
}
