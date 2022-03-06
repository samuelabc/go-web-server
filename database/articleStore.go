package database

import (
	"fmt"
	articleController "web-server/api/article"
	models "web-server/model"

	"github.com/go-pg/pg"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
)

// ProfileStore implements database operations for profile management.
type ArticleStore struct {
	db *pg.DB
}

// NewAccountStore returns an AccountStore.
func NewArticleStore(db *pg.DB) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

// Get an article by ID.
func (s *ArticleStore) Get(id string) (*models.Article, error) {
	res := models.Article{ID: id}
	err := s.db.Model(&res).
		Where("article.id = ?", id).
		First()
	return &res, err
}

// Get an article by filter.
func (s *ArticleStore) List(a *articleController.ListArticleRequest) (*[]models.Article, error) {
	res := []models.Article{}
	err := s.db.Model(&res).
		Where("article.id = ?", a.ID, a.ID).
		Where("article.title = ?", a.Title, a.Title).
		Select()

	fmt.Println("res", a.ID, a.Title, res, err)
	return &res, err
}

// Create create a new article.
func (s *ArticleStore) Create(a *models.Article) error {
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		u, err := uuid.NewV4()
		if err != nil {
			return err
		}
		a.ID = u.String()

		err = tx.Insert(a)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// Update an article.
func (s *ArticleStore) Update(a *models.Article) error {
	_, err := s.db.Model(a).
		WherePK().
		Update()
	return err
}

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

// UpdateToken updates a jwt refresh token.
func (s *ArticleStore) UpdateToken(t *jwt.Token) error {
	_, err := s.db.Model(t).
		Column("identifier").
		WherePK().
		Update()
	return err
}

// DeleteToken deletes a jwt refresh token.
func (s *ArticleStore) DeleteToken(t *jwt.Token) error {
	err := s.db.Delete(t)
	return err
}
