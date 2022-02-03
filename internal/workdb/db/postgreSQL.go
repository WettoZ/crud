package workdb

import (
	"context"
	"crud/internal/pkg/client/postgresql"
	wdb "crud/internal/workdb"
	"fmt"
)

type PoolPGX struct {
	client postgresql.Connect
}

func (p *PoolPGX) InsertRow(ctx context.Context, us *wdb.UserData) error {
	sql := `
		INSERT INTO users(userid, name, password) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`
	if err := p.client.QueryRow(ctx, sql, us.Uid, us.Name, us.Passwd).Scan(&us.Num); err != nil {
		fmt.Println("[ERROR] [DATABASE] INSERT: ", err)
		return err
	}
	return nil
}

func (p *PoolPGX) DeleteRow(ctx context.Context, element int64) error {
	_, err := p.client.Exec(ctx, "DELETE FROM users WHERE id = $1", element)
	if err != nil {
		fmt.Println("[ERROR] [DATABASE] DELETE: ", err)
		return err
	}

	return nil
}

func (p *PoolPGX) AllRows(ctx context.Context) ([]*wdb.UserData, error) {
	var us wdb.UserData
	var list = []*wdb.UserData{}

	rows, err := p.client.Query(ctx, "SELECT id, userid, name, password FROM users")
	if err != nil {
		fmt.Println("[ERROR] [DATABASE] SELECT: ", err)
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&us.Num, &us.Uid, &us.Name, &us.Passwd); err != nil {
			fmt.Println("[ERROR] Scan")
			return nil, err
		}
		fmt.Println(us)
		list = append(list, &us)
	}

	return list, nil
}

func NewpoolPGX(conn postgresql.Connect) *PoolPGX {
	return &PoolPGX{
		client: conn,
	}
}
