package account

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq" // what is this for? ---> to register the postgres driver with database/sql package
)

type Repository interface {
	Close()
	PutAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip, take int) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db: db}, nil
}


func (r *postgresRepository) Close() {
	r.db.Close()
}


func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}	

func (r *postgresRepository) PutAccount(ctx  context.Context , a Account) error {
	_ , err := r.db.ExecContext(ctx, "INSERT INTO accounts (id, name) VALUES ($1, $2)", a.ID, a.Name)
	if err != nil {
		return err
	}
	return nil
}



func ( r *postgresRepository ) GetAccountByID(ctx context.Context, id string) (*Account, error) {

	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM accounts WHERE id = $1", id)
	var a Account
	err := row.Scan(&a.ID, &a.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}




func ( r *postgresRepository ) ListAccounts(ctx context.Context, skip, take int) ([]Account, error) {
	rows , err := r.db.QqueryContext(ctx, "SELECT id, name FROM accounts ORDER BY id OFFSET $1 LIMIT $2" , skip, take)

	if err != nil {
		return nil , err
	}

	defer rows.Close()
	var accounts []Account
	for rows.Next() {
		var a Account
		err := rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil , err
		}
		accounts = append(accounts, a)
	}
	return accounts , nil
}




