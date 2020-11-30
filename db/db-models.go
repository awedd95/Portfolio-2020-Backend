package database

import (
    "fmt"
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
    "server/models"
)

type User struct {
    ID     int64
    Name   string
    Password   string
    Emails []string
}

func (u User) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Emails)
}

type BlogPost struct {
    ID       int64
    Title    string
    Body    string
}

func (s BlogPost) String() string {
    return fmt.Sprintf("BlogPost<%d %s %s>", s.ID, s.Title)
}

// createSchema creates database schema for User and Story models.
func CreateSchema(db *pg.DB) error {
    models := []interface{}{
        (*models.User)(nil),
        (*BlogPost)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}
