package sqlite3_test

import (
	"database/sql"
	"flag"
	hp "harry-potter"
	"harry-potter/store/option"
	store "harry-potter/store/sqlite3"
	"testing"

	s3 "github.com/hyphengolang/prelude/sql"
	"github.com/hyphengolang/prelude/testing/is"
	_ "github.com/mattn/go-sqlite3"
)

var dbPath = flag.String("db", ":memory:", "path to sqlite3 database")

func TestRepo(t *testing.T) {
	is := is.New(t)

	conn, err := sql.Open("sqlite3", *dbPath)
	is.NoErr(err) // sqlite3 database connection

	t.Cleanup(func() {
		conn.Close()
	})

	t.Run("get all characters from database", func(t *testing.T) {
		var count int
		err := s3.QueryRow(conn, func(row *sql.Row) error { return row.Scan(&count) }, "SELECT count(*) FROM characters;")
		is.NoErr(err) // query count of characters

		is.Equal(count, 25) // there should be at least one character
	})

	t.Run("get one character from database", func(t *testing.T) {
		var c hp.Character
		scanner := func(row *sql.Row) error {
			return row.Scan(&c.ID, &c.Name, &c.Blood, &c.Species, &c.Patronus, &c.Born, &c.Quote, &c.ImgURL)
		}

		err := s3.QueryRow(conn, scanner, "SELECT * FROM characters WHERE id = 1;")
		is.NoErr(err) // query character name

		is.Equal(c.Name, "Ronald (Ron) Weasley") // character should be `Ronald (Ron) Weasley`
	})

	r := store.New(conn)

	t.Run("get all characters from database using store", func(t *testing.T) {
		t.Skip()
		var f option.FilterOption
		cs, err := r.Filter(f)
		is.NoErr(err) // query all characters

		is.Equal(len(cs), 25) // there should be at least one character
	})

	t.Run("get one character from database using store", func(t *testing.T) {
		c, err := r.Search(1)
		is.NoErr(err) // query character name

		is.Equal(c.Name, "Ronald (Ron) Weasley") // character should be `Ronald (Ron) Weasley`
	})

	t.Run("return all half-blood characters", func(t *testing.T) {
		var q = option.FilterOption{
			Blood: "half-blood",
		}
		cs, err := r.Filter(q)
		is.NoErr(err)         // filter by blood type
		is.Equal(len(cs), 11) // X characters in the list
	})

	t.Run("return all characters that born in February", func(t *testing.T) {
		var q = option.FilterOption{
			Month: "february",
		}
		cs, err := r.Filter(q)
		is.NoErr(err)        // filter by blood type
		is.Equal(len(cs), 2) // X characters in the list
	})

	t.Run("return all characters that are `half-blood` and born in `September`", func(t *testing.T) {
		var q = option.FilterOption{
			Blood: "half-blood",
			Month: "Sep",
		}
		cs, err := r.Filter(q)
		is.NoErr(err)        // filter by blood type
		is.Equal(len(cs), 4) // X characters in the list
	})

	t.Run("get one character from database using store", func(t *testing.T) {
		var q = option.FilterOption{
			Name: "ronald",
		}
		cs, err := r.Filter(q)
		is.NoErr(err)                                // filter by blood type
		is.Equal(cs[0].Name, "Ronald (Ron) Weasley") // character should be `Ronald (Ron) Weasley`
	})
}
