package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/reinhardjs/sayakaya/domain"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func NewMysqlUserRepository(conn *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{conn}
}

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}

		err = rows.Scan(
			&t.ID,
			&t.Email,
			&t.VerifiedStatus,
			&t.Birthday,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) Fetch(ctx context.Context) (res []domain.User, err error) {
	query := `SELECT * FROM user`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}
func (m *mysqlUserRepository) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	query := `SELECT * FROM user WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.User{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlUserRepository) Store(ctx context.Context, a *domain.User) (err error) {
	query := `INSERT user SET email=?, verifiedStatus=?, birthday=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Email, a.VerifiedStatus, a.Birthday)
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

func (m *mysqlUserRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM user WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}

func (m *mysqlUserRepository) Update(ctx context.Context, ar *domain.User) (err error) {
	query := `UPDATE user set SET email=?, verifiedStatus=?, birthday=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Email, ar.VerifiedStatus, ar.Birthday, ar.ID)
	if err != nil {
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
