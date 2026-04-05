package sql

import (
	"context"
	"study_docker/employee"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteEmployee(
	ctx context.Context,
	pool *pgxpool.Pool,
	id employee.Employee,
) string {
	sqlQuery := `
		DELETE FROM employees WHERE id=$1
	`
	tag, err := pool.Exec(
		ctx,
		sqlQuery,
		id.ID,
	)
	if err != nil {
		panic(err)
	}

	if tag.RowsAffected() == 0 {
		return "такого id нету"
	}

	return ""
}
