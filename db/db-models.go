package database

import (
    "fmt"
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
)

type User struct {
    Id     int64
    Name   string
    Password   string
    Emails []string
}

func (u User) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type BlogPost struct {
    Id       int64
    Title    string
    Body    string
}

func (s BlogPost) String() string {
    return fmt.Sprintf("BlogPost<%d %s %s>", s.Id, s.Title)
}

// createSchema creates database schema for User and Story models.
func CreateSchema(db *pg.DB) error {
    models := []interface{}{
        (*User)(nil),
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
