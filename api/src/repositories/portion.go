package repositories

import (
	"database/sql"
	"fmt"
	"go/src/model"

	_ "github.com/go-sql-driver/mysql"
)

type portions struct {
	db *sql.DB
}

//NewRepoPortion create a new portion repository
func NewRepoPortion(db *sql.DB) *portions {
	return &portions{db}

}

//Create insert a new portion on database
func (p portions) Create(portion model.Portion) (uint64, error) {
	statement, erro := p.db.Prepare(
		"insert into portion (quantity) values (?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(portion.Quantity)
	if erro != nil {
		return 0, erro
	}

	lastIDEntered, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDEntered), nil
}

// Find search for portions by quantity
func (p portions) Find(quantity string) ([]model.Portion, error) {
	quantity = fmt.Sprintf("%%%s%%", quantity)

	lines, erro := p.db.Query("select id, quantity from portion where quantity LIKE ?", quantity)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var portions []model.Portion

	for lines.Next() {
		var portion model.Portion

		if erro = lines.Scan(
			&portion.ID,
			&portion.Quantity,
		); erro != nil {
			return nil, erro
		}

		portions = append(portions, portion)
	}
	return portions, nil
}

func (p portions) FindById(ID uint64) (model.Portion, error) {
	lines, erro := p.db.Query("select id, quantity from portion where id = ?", ID)
	if erro != nil {
		return model.Portion{}, erro
	}
	defer lines.Close()

	var portion model.Portion

	if lines.Next() {
		if erro = lines.Scan(
			&portion.ID,
			&portion.Quantity,
		); erro != nil {
			return model.Portion{}, erro
		}
	}
	return portion, erro
}

func (p portions) Update(ID uint64, portion model.Portion) error {
	statement, erro := p.db.Prepare("update portion set quantity = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(portion.Quantity, ID); erro != nil {
		return erro
	}
	return nil
}

func (p portions) Delete(ID uint64) error {
	statement, erro := p.db.Prepare("delete from portion where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
