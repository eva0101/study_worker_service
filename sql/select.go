package sql

import (
	"context"
	"study_docker/employee"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SelectEmployees(
	ctx context.Context,
	pool *pgxpool.Pool,
) []employee.Employee {
	sqlQuery := `
		SELECT id, fullName, position 
		FROM employees
	`

	rows, err := pool.Query(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var listEmployees []employee.Employee

	for rows.Next() {
		var list employee.Employee
		if err := rows.Scan(&list.ID, &list.FullName, &list.Position); err != nil {
			panic(err)
		}

		listEmployees = append(listEmployees, list)
	}

	return listEmployees
}
