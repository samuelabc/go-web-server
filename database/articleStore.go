package database

import (
	"context"
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
func (s *ArticleStore) Get(id string) (*models.Article, error) {
	var err error
	ctx := context.Background()
	row := s.db.QueryRow(ctx, `select *
		from articles
		where articles.id = $1
		limit 1`, id)
	var res models.Article
	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Get an article by filter.
func (s *ArticleStore) List(a *articleController.ListArticleRequest) (*[]models.Article, error) {
	// res := []models.Article{}
	// err := s.db.Model(&res).
	// 	Where("article.id = ?", a.ID, a.ID).
	// 	Where("article.title = ?", a.Title, a.Title).
	// 	Select()

	// fmt.Println("res", a.ID, a.Title, res, err)
	// return &res, err
	return nil, nil
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
// func (s *ArticleStore) Update(a *models.Article) error {
// 	_, err := s.db.Model(a).
// 		WherePK().
// 		Update()
// 	return err
// }

// Delete an account.
// func (s *ArticleStore) Delete(a *pwdless.Account) error {
// 	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
// 		if _, err := tx.Model(&jwt.Token{}).
// 			Where("account_id = ?", a.ID).
// 			Delete(); err != nil {
// 			return err
// 		}
// 		if _, err := tx.Model(&models.Profile{}).
// 			Where("account_id = ?", a.ID).
// 			Delete(); err != nil {
// 			return err
// 		}
// 		return tx.Delete(a)
// 	})
// 	return err
// }
