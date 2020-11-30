package db

import (
    "fmt"
    "os"
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
    Author   *User `pg:"rel:has-one"`
}

func (s BlogPost) String() string {
    return fmt.Sprintf("BlogPost<%d %s %s>", s.Id, s.Title, s.Author)
}

func Hello(name string) string {
   return "hello " + name
}

func DB_Model() {
    dbpass := os.Getenv("PASSWORD")
    fmt.Println(dbpass)
    db := pg.Connect(&pg.Options{
        Addr:     ":5432",
        User: "postgres",
        Password: dbpass,
    })
    defer db.Close()
    err := createSchema(db)
    if err != nil {
        panic(err)
    }
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
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
