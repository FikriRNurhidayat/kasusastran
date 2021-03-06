// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package query

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.countSeratsStmt, err = db.PrepareContext(ctx, countSerats); err != nil {
		return nil, fmt.Errorf("error preparing query CountSerats: %w", err)
	}
	if q.countWulangansStmt, err = db.PrepareContext(ctx, countWulangans); err != nil {
		return nil, fmt.Errorf("error preparing query CountWulangans: %w", err)
	}
	if q.createSeratStmt, err = db.PrepareContext(ctx, createSerat); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSerat: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.createWulanganStmt, err = db.PrepareContext(ctx, createWulangan); err != nil {
		return nil, fmt.Errorf("error preparing query CreateWulangan: %w", err)
	}
	if q.deleteSeratStmt, err = db.PrepareContext(ctx, deleteSerat); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSerat: %w", err)
	}
	if q.deleteWulanganStmt, err = db.PrepareContext(ctx, deleteWulangan); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteWulangan: %w", err)
	}
	if q.doesUserEmailExistStmt, err = db.PrepareContext(ctx, doesUserEmailExist); err != nil {
		return nil, fmt.Errorf("error preparing query DoesUserEmailExist: %w", err)
	}
	if q.getSeratStmt, err = db.PrepareContext(ctx, getSerat); err != nil {
		return nil, fmt.Errorf("error preparing query GetSerat: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.getWulanganStmt, err = db.PrepareContext(ctx, getWulangan); err != nil {
		return nil, fmt.Errorf("error preparing query GetWulangan: %w", err)
	}
	if q.listSeratsStmt, err = db.PrepareContext(ctx, listSerats); err != nil {
		return nil, fmt.Errorf("error preparing query ListSerats: %w", err)
	}
	if q.listWulangansStmt, err = db.PrepareContext(ctx, listWulangans); err != nil {
		return nil, fmt.Errorf("error preparing query ListWulangans: %w", err)
	}
	if q.updateSeratStmt, err = db.PrepareContext(ctx, updateSerat); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSerat: %w", err)
	}
	if q.updateWulanganStmt, err = db.PrepareContext(ctx, updateWulangan); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateWulangan: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countSeratsStmt != nil {
		if cerr := q.countSeratsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countSeratsStmt: %w", cerr)
		}
	}
	if q.countWulangansStmt != nil {
		if cerr := q.countWulangansStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countWulangansStmt: %w", cerr)
		}
	}
	if q.createSeratStmt != nil {
		if cerr := q.createSeratStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSeratStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.createWulanganStmt != nil {
		if cerr := q.createWulanganStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createWulanganStmt: %w", cerr)
		}
	}
	if q.deleteSeratStmt != nil {
		if cerr := q.deleteSeratStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSeratStmt: %w", cerr)
		}
	}
	if q.deleteWulanganStmt != nil {
		if cerr := q.deleteWulanganStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteWulanganStmt: %w", cerr)
		}
	}
	if q.doesUserEmailExistStmt != nil {
		if cerr := q.doesUserEmailExistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing doesUserEmailExistStmt: %w", cerr)
		}
	}
	if q.getSeratStmt != nil {
		if cerr := q.getSeratStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSeratStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.getWulanganStmt != nil {
		if cerr := q.getWulanganStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getWulanganStmt: %w", cerr)
		}
	}
	if q.listSeratsStmt != nil {
		if cerr := q.listSeratsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listSeratsStmt: %w", cerr)
		}
	}
	if q.listWulangansStmt != nil {
		if cerr := q.listWulangansStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listWulangansStmt: %w", cerr)
		}
	}
	if q.updateSeratStmt != nil {
		if cerr := q.updateSeratStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSeratStmt: %w", cerr)
		}
	}
	if q.updateWulanganStmt != nil {
		if cerr := q.updateWulanganStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateWulanganStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	countSeratsStmt        *sql.Stmt
	countWulangansStmt     *sql.Stmt
	createSeratStmt        *sql.Stmt
	createUserStmt         *sql.Stmt
	createWulanganStmt     *sql.Stmt
	deleteSeratStmt        *sql.Stmt
	deleteWulanganStmt     *sql.Stmt
	doesUserEmailExistStmt *sql.Stmt
	getSeratStmt           *sql.Stmt
	getUserByEmailStmt     *sql.Stmt
	getWulanganStmt        *sql.Stmt
	listSeratsStmt         *sql.Stmt
	listWulangansStmt      *sql.Stmt
	updateSeratStmt        *sql.Stmt
	updateWulanganStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		countSeratsStmt:        q.countSeratsStmt,
		countWulangansStmt:     q.countWulangansStmt,
		createSeratStmt:        q.createSeratStmt,
		createUserStmt:         q.createUserStmt,
		createWulanganStmt:     q.createWulanganStmt,
		deleteSeratStmt:        q.deleteSeratStmt,
		deleteWulanganStmt:     q.deleteWulanganStmt,
		doesUserEmailExistStmt: q.doesUserEmailExistStmt,
		getSeratStmt:           q.getSeratStmt,
		getUserByEmailStmt:     q.getUserByEmailStmt,
		getWulanganStmt:        q.getWulanganStmt,
		listSeratsStmt:         q.listSeratsStmt,
		listWulangansStmt:      q.listWulangansStmt,
		updateSeratStmt:        q.updateSeratStmt,
		updateWulanganStmt:     q.updateWulanganStmt,
	}
}
