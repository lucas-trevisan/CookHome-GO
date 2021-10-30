package repositories

import (
	"database/sql"
	"fmt"
	"go/src/model"

	_ "github.com/go-sql-driver/mysql"
)

type groups struct {
	db *sql.DB
}

//NewRepoGroup create a new group repository
func NewRepoGroup(db *sql.DB) *groups {
	return &groups{db}

}

//Create insert a new group on database
func (g groups) Create(group model.Group) (uint64, error) {
	statement, erro := g.db.Prepare(
		"insert into group1 (name) values (?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(group.Name)
	if erro != nil {
		return 0, erro
	}

	lastIDEntered, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDEntered), nil
}

// Find search for groups by name
func (g groups) Find(name string) ([]model.Group, error) {
	name = fmt.Sprintf("%%%s%%", name)

	lines, erro := g.db.Query("select id, name from group1 where name LIKE ?", name)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var groups []model.Group

	for lines.Next() {
		var group model.Group

		if erro = lines.Scan(
			&group.ID,
			&group.Name,
		); erro != nil {
			return nil, erro
		}

		groups = append(groups, group)
	}
	return groups, nil
}

func (g groups) FindById(ID uint64) (model.Group, error) {
	lines, erro := g.db.Query("select id, name from group1 where id = ?", ID)
	if erro != nil {
		return model.Group{}, erro
	}
	defer lines.Close()

	var group model.Group

	if lines.Next() {
		if erro = lines.Scan(
			&group.ID,
			&group.Name,
		); erro != nil {
			return model.Group{}, erro
		}
	}
	return group, erro
}

func (g groups) Update(ID uint64, group model.Group) error {
	statement, erro := g.db.Prepare("update group1 set name = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(group.Name, ID); erro != nil {
		return erro
	}
	return nil
}

func (g groups) Delete(ID uint64) error {
	statement, erro := g.db.Prepare("delete from group1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
