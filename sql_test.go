package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "INSERT INTO users (name, email, password) VALUES ('test', 'tests@gmail.com', 'test')")

	if err != nil {
		panic(err)
	}
	fmt.Println("Executed")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, name, email FROM users")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string

		err = rows.Scan(&id, &name, &email)

		if err != nil {
			panic(err)
		}

		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("email", email)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, name, email, password, created_at FROM users")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		var password sql.NullString
		var createdAt sql.NullTime

		err = rows.Scan(&id, &name, &email, &password, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("email", email)
		if password.Valid {
			fmt.Println("password", password)
		}
		if createdAt.Valid {
			fmt.Println("createdAt", createdAt)
		}
	}
}
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "superadmin@energasindo.com"
	password := "$2y$10$mGUtPEJ45K21BQHT7W7HneW3D4vQxS0hLdB8TEVwXMEl0dCmlS2uy"

	rows, err := db.QueryContext(ctx, "SELECT email FROM users "+
		"WHERE email = '"+username+"' AND password = '"+password+"' LIMIT 1")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var email string
		err := rows.Scan(&email)

		if err != nil {
			panic(err)
		}
		fmt.Println("email", email)
	} else {
		fmt.Println("user not found")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "superadmin@energasindo.com"
	password := "$2y$10$mGUtPEJ45K21BQHT7W7HneW3D4vQxS0hLdB8TEVwXMEl0dCmlS2uy"

	script := "SELECT email FROM users WHERE email = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var email string
		err := rows.Scan(&email)

		if err != nil {
			panic(err)
		}
		fmt.Println("email", email)
	} else {
		fmt.Println("user not found")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "INSERT INTO users (name, email, password) "+
		"VALUES (?, ?, ?)", "test", "testmail@mail.com", "test")

	if err != nil {
		panic(err)
	}
	fmt.Println("Executed")
}

func TestAutoInc(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	name := "test s"
	description := "test description"

	result, err := db.ExecContext(ctx, "INSERT INTO categories (name, description) "+
		"VALUES (?, ?)", name, description)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert new category with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO categories (name, description) VALUES (?, ?)"

	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		name := "Kategori mantaps " + strconv.Itoa(i)
		description := "Kategori mantaps description " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, name, description)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert new category with id", id)
	}
}
