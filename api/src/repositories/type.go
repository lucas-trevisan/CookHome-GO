package repositories

import (
	"database/sql"
	"fmt"
	"go/src/model"

	_ "github.com/go-sql-driver/mysql"
)

type types struct {
	db *sql.DB
}

//NewRepoType create a new typ repository
func NewRepoType(db *sql.DB) *types {
	return &types{db}

}

//Create insert a new typ on database
func (t types) Create(typ model.Type) (uint64, error) {
	statement, erro := t.db.Prepare(
		"insert into type (type) values (?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(typ.Type)
	if erro != nil {
		return 0, erro
	}

	lastIDEntered, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDEntered), nil
}

// Find search for types by name
func (t types) Find(name string) ([]model.Type, error) {
	name = fmt.Sprintf("%%%s%%", name)

	lines, erro := t.db.Query("select id, name from type where type LIKE ?", name)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var types []model.Type

	for lines.Next() {
		var typ model.Type

		if erro = lines.Scan(
			&typ.ID,
			&typ.Type,
		); erro != nil {
			return nil, erro
		}

		types = append(types, typ)
	}
	return types, nil
}

func (t types) FindById(ID uint64) (model.Type, error) {
	lines, erro := t.db.Query("select id, type from type where id = ?", ID)
	if erro != nil {
		return model.Type{}, erro
	}
	defer lines.Close()

	var typ model.Type

	if lines.Next() {
		if erro = lines.Scan(
			&typ.ID,
			&typ.Type,
		); erro != nil {
			return model.Type{}, erro
		}
	}
	return typ, erro
}

func (t types) Update(ID uint64, typ model.Type) error {
	statement, erro := t.db.Prepare("update type set type = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(typ.Type, ID); erro != nil {
		return erro
	}
	return nil
}

func (t types) Delete(ID uint64) error {
	statement, erro := t.db.Prepare("delete from type where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
