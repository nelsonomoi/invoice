package mysql

import (
	"context"
	"database/sql"

	"github.com/nelsonomoi/invoice/domain"
)

type mysqlClientRepository struct {
	Conn *sql.DB
}


func NewMysqlClientRepository(Conn *sql.DB) domain.ClientRepository  {
	return &mysqlClientRepository{Conn}
}



func (m *mysqlClientRepository) Store(ctx context.Context, a *domain.Client) (err error) {
	query := `INSERT  client SET client_name=? , email=? , phone=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Client_name, a.Email, a.Phone)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}
