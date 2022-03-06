package migrate

import (
	"fmt"

	"github.com/go-pg/migrations"
)

const articleTable = `
CREATE TABLE articles (
id uuid NOT NULL,
title text NOT NULL,
content text NOT NULL,
created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
updated_at timestamp with time zone DEFAULT current_timestamp,
PRIMARY KEY (id)
)`

func init() {
	up := []string{
		articleTable,
	}

	down := []string{
		`DROP TABLE articles`,
	}

	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating initial tables")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping initial tables")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
