package storages

import (
	"context"
	"fmt"

	"github.com/kurojs/ipinfo/ent"
	_ "github.com/lib/pq"
)

type (
	Storage interface {
		CreateHistory(ctx context.Context, history *ent.History) (*ent.History, error)
		GetHistories(ctx context.Context) (ent.Histories, error)
		Close()
	}

	DBStorage struct {
		db *ent.Client
	}
)

func NewStorage(host, port, user, dbname, pwd string) (Storage, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, pwd))
	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return &DBStorage{
		db: client,
	}, nil
}

func (s *DBStorage) CreateHistory(ctx context.Context, history *ent.History) (*ent.History, error) {
	return s.db.History.Create().
		SetCity(history.City).
		SetIP(history.IP).
		SetRegion(history.Region).
		SetLoginTime(history.LoginTime).
		Save(ctx)
}

func (s *DBStorage) GetHistories(ctx context.Context) (ent.Histories, error) {
	return s.db.History.Query().Limit(100).All(ctx)
}

func (s *DBStorage) Close() {
	if s != nil {
		s.db.Close()
	}
}
