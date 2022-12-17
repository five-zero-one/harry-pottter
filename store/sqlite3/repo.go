package sqlite3

import (
	"database/sql"
	hp "harry-potter"
	"harry-potter/store"
	"harry-potter/store/option"

	s3 "github.com/hyphengolang/prelude/sql"
)

// NOTE - in production I really should be defining a type that is not the same as teh application type for `Character`
// this is just for simplicity of the example
type repo struct {
	conn *sql.DB
}

// Filter implements store.Repo
func (r *repo) Filter(f option.FilterOption) ([]hp.Character, error) {
	blood := sql.Named("blood", f.Blood)
	month := sql.Named("month", f.Month)
	name := sql.Named("name", f.Name)

	qry := `
	SELECT * FROM characters
	WHERE
		(CASE 
			WHEN ifnull(lower(@blood),'') = '' THEN 1
			WHEN instr(lower(blood),lower(@blood)) > 0 THEN 1
			ELSE 0 
		END) = 1
	AND
		(CASE 
			WHEN ifnull(lower(@month),'') = '' THEN 1
			WHEN instr(lower(born),lower(@month)) > 0 THEN 1
			ELSE 0 
		END) = 1
	AND
		(CASE 
			WHEN ifnull(lower(@name),'') = '' THEN 1
			WHEN instr(lower(name),lower(@name)) > 0 THEN 1
			ELSE 0 
		END) = 1
	`

	//

	return s3.Query(r.conn, func(rows *sql.Rows, c *hp.Character) error {
		return rows.Scan(&c.ID, &c.Name, &c.Blood, &c.Species, &c.Patronus, &c.Born, &c.Quote, &c.ImgURL)
	}, qry, blood, month, name)
}

// Search implements store.Repo
func (r *repo) Search(key any) (*hp.Character, error) {
	var c hp.Character
	err := s3.QueryRow(r.conn, func(row *sql.Row) error {
		return row.Scan(&c.ID, &c.Name, &c.Blood, &c.Species, &c.Patronus, &c.Born, &c.Quote, &c.ImgURL)
	}, "SELECT * FROM characters WHERE id = ?;", key)
	return &c, err
}

func New(conn *sql.DB) store.Repo {
	return &repo{conn}
}
