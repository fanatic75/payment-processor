package core

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbService struct {
	Client  *pgxpool.Pool
	Context context.Context
}

func (dbService *DbService) InitDbClient(connStr string, ctx context.Context) error {
	client, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return err
	}
	err = client.Ping(ctx)
	if err != nil {
		return err
	}
	dbService.Client = client
	dbService.Context = ctx
	return nil
}
