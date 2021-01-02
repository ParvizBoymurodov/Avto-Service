package psql

import "github.com/jackc/pgx/v4/pgxpool"

type DB struct {
	conn *pgxpool.Pool
}

func NewDB(conn *pgxpool.Pool) *DB {
	if conn == nil {
		panic("conn can't be nii")
	}
	return &DB{conn: conn}
}