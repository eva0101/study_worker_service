package sql

import (
	"context"
	"study_docker/employee"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertEmployee(
	ctx context.Context,
	pool *pgxpool.Pool,
	employee employee.Employee,
) {
	sqlQuery := `
		INSERT INTO employees (fullName, position)
		VALUES ($1, $2);
	`
	_, err := pool.Exec(
		ctx,
		sqlQuery,
		employee.FullName,
		employee.Position,
	)
	if err != nil {
		panic(err)
	}
}
