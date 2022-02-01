package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func connection(countRepeat int) (*pgxpool.Pool, error) {
	var db *pgxpool.Pool
	var err error

	param := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.PgUser, conf.PgPasswd, conf.PgHost, conf.PgPort, conf.PgDB)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < countRepeat; i++ {

		fmt.Print("Connection attempt: ")
		db, err = pgxpool.Connect(ctx, param)

		if err != nil {
			time.Sleep(time.Second)
			fmt.Println("Fail!")
			continue
		}

		fmt.Println("Success!")
		return db, nil
	}

	return nil, err
}
