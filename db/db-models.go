package database

import (
    "fmt"
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
    "server/models"
)

type BlogPost struct {
    ID       int64
    Title    string
    Body    string
}

func (s BlogPost) String() string {
    return fmt.Sprintf("BlogPost<%d %s >", s.ID, s.Title)
}

// createSchema creates database schema for User and Story models.
func CreateSchema(db *pg.DB) error {
    models := []interface{}{
        (*models.User)(nil),
        (*models.Project)(nil),
        (*BlogPost)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            IfNotExists: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}
