package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"log"
	"sync"
)

type PostgresPlayerStore struct {
	storage  *postgres.PostgresContainer
	connOnce sync.Once
	conn     *pgx.Conn
}

func NewPostgresPlayerStore() *PostgresPlayerStore {
	postgresStore := &PostgresPlayerStore{storage: StartPostgresContainer()}
	postgresStore.createTable()
	return postgresStore
}
func (p *PostgresPlayerStore) RecordWin(name string) {
	_, err := p.getConnection().Exec(context.Background(), "INSERT INTO users(name) VALUES ($1)", name)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	var count int
	err := p.getConnection().QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE name = $1", name).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func (p *PostgresPlayerStore) createTable() {
	_, _, err := p.storage.Exec(context.Background(),
		[]string{"psql", "-U", "user", "-d", "users", "-c", "CREATE TABLE users (id SERIAL, name TEXT NOT NULL)"})
	if err != nil {
		log.Fatal(err)
	}
}

func (p *PostgresPlayerStore) getConnection() *pgx.Conn {
	p.connOnce.Do(func() {
		connectionString, err := p.storage.ConnectionString(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		conn, err := pgx.Connect(context.Background(), connectionString)
		if err != nil {
			log.Fatal(err)
		}
		p.conn = conn
	})
	return p.conn
}

func (p *PostgresPlayerStore) Close() {
	if p.conn != nil {
		p.conn.Close(context.Background())
	}
	if p.storage != nil {
		p.storage.Terminate(context.Background())
	}
}
