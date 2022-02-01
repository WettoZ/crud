package main

import (
	"context"
	"fmt"

	pb "crud/guser"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertRow(ctx context.Context, pool *pgxpool.Pool, us *pb.User) error {
	if _, err := pool.Exec(ctx, "INSERT INTO users(userid, name, password) VALUES ($1, $2, $3);", us.Uid, us.Name, us.Passwd); err != nil {
		fmt.Println("[ERROR] [DATABASE] INSERT: ", err)
		return err
	}
	return nil
}

func DeleteRow(ctx context.Context, pool *pgxpool.Pool, element int64) error {
	_, err := pool.Exec(ctx, "DELETE FROM users WHERE id = $1", element)
	if err != nil {
		fmt.Println("[ERROR] [DATABASE] DELETE: ", err)
		return err
	}

	return nil
}

func AllRows(ctx context.Context, pool *pgxpool.Pool) ([]*pb.User, error) {
	var us pb.User
	var list = []*pb.User{}

	rows, err := pool.Query(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Println("[ERROR] [DATABASE] SELECT: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.FieldDescriptions()
		if err := rows.Scan(&us.Num, &us.Uid, &us.Name, &us.Passwd); err != nil {
			fmt.Println("[ERROR] Scan")
			return nil, err
		}
		list = append(list, &us)
	}

	return list, nil
}
