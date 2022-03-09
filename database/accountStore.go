package database

import (
	"context"
	accountController "web-server/api/account"
	accountModel "web-server/model/account"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// ProfileStore implements database operations for profile management.
type AccountStore struct {
	db *pgxpool.Pool
}

// NewAccountStore returns an AccountStore.
func NewAccountStore(db *pgxpool.Pool) *AccountStore {
	return &AccountStore{
		db: db,
	}
}

// Get an account by ID.
func (s *AccountStore) Get(ctx context.Context, a *accountController.GetAccountRequest) (*accountModel.Account, error) {
	var err error
	var res accountModel.Account

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, param, err := psql.Select("*").From("accounts").Where(sq.Eq{"id": *a.ID}).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}
	err = s.db.QueryRow(ctx, query, param...).Scan(&res.ID, &res.Name, &res.PasswordHash, &res.Email, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &res, err
}

// Register a new account.
func (s *AccountStore) Register(ctx context.Context, a *accountController.RegisterAccountRequest) (*accountModel.Account, error) {
	var err error

	u, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	newAccount := map[string]interface{}{
		"id":            &u,
		"name":          a.Name,
		"password_hash": a.Password,
		"email":         a.Email,
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var res accountModel.Account

	var values []interface{}
	var columns []string
	for k, _ := range newAccount {
		if v, ok := newAccount[k]; ok {
			columns = append(columns, k)
			values = append(values, v)
		}
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query, param, err := psql.Insert("accounts").Columns(columns...).Values(values...).Suffix("Returning *").ToSql()
	if err != nil {
		return nil, err
	}
	// fmt.Println("query", query, param)
	err = s.db.QueryRow(ctx, query, param...).Scan(&res.ID, &res.Name, &res.PasswordHash, &res.Email, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get an article by filter.
func (s *AccountStore) List(ctx context.Context, a *accountController.ListAccountRequest) ([]*accountModel.Account, error) {
	var err error

	// account := map[string]interface{}{
	// 	"id":    a.ID,
	// 	"name":  a.Name,
	// 	"email": a.Email,
	// }

	// psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	// accounts := psql.Select("*").From("accounts")
	// for k, _ := range account {
	// 	if v, ok := account[k]; ok {
	// 		if !reflect.ValueOf(v).IsNil() {
	// 			T := reflect.TypeOf(v).Elem()
	// 			fmt.Printf("type name", T)
	// 			fmt.Println("value of k", v)
	// 			fmt.Printf("value of k, %v %T\n", v, v)

	// 			switch T.Name() {
	// 			case "stringvalue":
	// 				accounts = accounts.Where(fmt.Sprint(k, "=?"), v.(*string))

	// 			case "UUIDvalue":
	// 				accounts = accounts.Where(fmt.Sprint(k, "=?"), v.(*uuid.UUID))

	// 			}
	// 			// accounts = accounts.Where(fmt.Sprint(k, "=?"), v.(T))
	// 		}
	// 	}
	// }
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	accounts := psql.Select("*").From("accounts")
	if a.ID != nil {
		accounts = accounts.Where("id = ?", *a.ID)
	}
	if a.Name != nil {
		accounts = accounts.Where("name = ?", *a.Name)
	}
	if a.Email != nil {
		accounts = accounts.Where("email = ?", *a.Email)
	}
	query, param, err := accounts.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.Query(ctx, query, param...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := []*accountModel.Account{}
	pgxscan.ScanAll(&res, rows)
	if err != nil {
		return nil, err
	}
	return res, err
}

// Create create a new article.
// func (s *AccountStore) Create(a *userModel.User) (*userModel.User, error) {
// 	var err error

// 	u, err := uuid.NewV4()
// 	if err != nil {
// 		return nil, err
// 	}
// 	a.ID = u.String()

// 	tx, err := s.db.BeginTx(context.Background(), pgx.TxOptions{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback(context.Background())
// 		} else {
// 			tx.Commit(context.Background())
// 		}
// 	}()

// 	var res models.Article
// 	row := tx.QueryRow(context.Background(), `
// 		INSERT INTO articles(id, title, content)
// 		VALUES ($1, $2, $3)
// 		RETURNING *;
// 	`, a.ID, a.Title, a.Content)
// 	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }

// Update an article.
// func (s *AccountStore) Update(a *userController.UpdateUserRequest) (*userModel.User, error) {
// 	var err error

// 	var query string
// 	var filters []string
// 	query = `UPDATE articles`

// 	//filter
// 	if a.Title != nil {
// 		filters = append(filters, fmt.Sprint("title = ", "'", *a.Title, "'"))
// 	}
// 	if a.Content != nil {
// 		filters = append(filters, fmt.Sprint("content = ", "'", *a.Content, "'"))
// 	}
// 	if len(filters) > 0 {
// 		filterQuery := strings.Join(filters, " , ")
// 		query = fmt.Sprint(query, " SET ", filterQuery)
// 	}
// 	query = fmt.Sprint(query, " WHERE articles.ID = ", "'", *a.ID, "'")

// 	//limit
// 	// query = fmt.Sprint(query, " limit ", 100)
// 	query = fmt.Sprint(query, " returning *")

// 	tx, err := s.db.BeginTx(context.Background(), pgx.TxOptions{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback(context.Background())
// 		} else {
// 			tx.Commit(context.Background())
// 		}
// 	}()

// 	var res models.Article
// 	row := tx.QueryRow(context.Background(), query)
// 	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }

// Delete an account.
// func (s *AccountStore) Delete(a *articleController.DeleteArticleRequest) (*userModel.User, error) {
// 	var err error

// 	tx, err := s.db.BeginTx(context.Background(), pgx.TxOptions{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback(context.Background())
// 		} else {
// 			tx.Commit(context.Background())
// 		}
// 	}()

// 	row := tx.QueryRow(context.Background(), `
// 		DELETE FROM articles
// 		WHERE articles.id = $1
// 		RETURNING *
// 	`, *a.ID)
// 	var res userModel.User
// 	err = row.Scan(&res.ID, &res.Title, &res.Content, &res.CreatedAt, &res.UpdatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &res, err
// }
