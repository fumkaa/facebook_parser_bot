package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) Storage {
	return Storage{
		db: db,
	}
}

var (
	ErrDublicateKey = errors.New("err Dublicate Key")
	ErrNoRows       = errors.New("error no rows")
)

func (s Storage) AddChatID(ctx context.Context, chat_id int) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddChatID]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `INSERT INTO data_user (chat_id) VALUES (?);`, chat_id)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return ErrDublicateKey
	}
	if err != nil {
		return fmt.Errorf("[AddChatID]set values chat_id in db error: %w", err)
	}

	return nil
}

func (s Storage) AddWaitMessage(ctx context.Context, chat_id int) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddWaitMessage]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_user SET wait_response = ? WHERE chat_id = ?;`, true, chat_id)
	if err != nil {
		return fmt.Errorf("[AddWaitMessage]set values wait_response in db error: %w", err)
	}
	return nil
}

func (s Storage) DeleteWaitMessage(ctx context.Context, chat_id int) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddWaitMessage]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_user SET wait_response = ? WHERE chat_id = ?;`, false, chat_id)
	if err != nil {
		return fmt.Errorf("[AddWaitMessage]set values wait_response in db error: %w", err)
	}
	return nil
}

func (s Storage) WaitMessage(ctx context.Context, chat_id int) (bool, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return false, fmt.Errorf("[ChatID]connection db error: %w", err)
	}
	defer conn.Close()

	var waitMessage bool
	if err := conn.GetContext(ctx, &waitMessage, `SELECT wait_response FROM data_user WHERE chat_id = ?`, chat_id); err != nil {
		return false, fmt.Errorf("[WaitMessage]get WaitMessage error: %w", err)
	}
	return waitMessage, nil
}

func (s Storage) AddCity(ctx context.Context, chat_id int, city string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddCity]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_user SET city = ? WHERE chat_id = ?;`, city, chat_id)
	if err != nil {
		return fmt.Errorf("[AddCity]set values city in db error: %w", err)
	}
	return nil
}

func (s Storage) City(ctx context.Context, chat_id int) (string, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return "", fmt.Errorf("[City]connection db error: %w", err)
	}
	defer conn.Close()

	var city string
	if err := conn.GetContext(ctx, &city, `SELECT city FROM data_user WHERE chat_id = ?`, chat_id); err != nil {
		return "", fmt.Errorf("[City]get city error: %w", err)
	}
	return city, nil
}

func (s Storage) DeleteCity(ctx context.Context, chat_id int) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddCity]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_user SET city = ? WHERE chat_id = ?;`, "", chat_id)
	if err != nil {
		return fmt.Errorf("[AddCity]set values city in db error: %w", err)
	}
	return nil
}

func (s Storage) AddChatIDFilters(ctx context.Context, chat_id int) (int, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, fmt.Errorf("[AddChatID]connection db error: %w", err)
	}
	defer conn.Close()

	data, err := conn.ExecContext(ctx, `INSERT INTO data_filters (chat_id) VALUES (?);`, chat_id)
	if err != nil {
		return 0, fmt.Errorf("[AddChatID]set values chat_id in db error: %w", err)
	}
	id, err := data.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get id error: %w", err)
	}
	return int(id), nil
}

func (s Storage) AddCityFilter(ctx context.Context, id int, city string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddCityFilter]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_filters SET city = ? WHERE id = ?;`, city, id)
	if err != nil {
		return fmt.Errorf("[AddCityFilter]set values city in db error: %w", err)
	}
	return nil
}
func (s Storage) AddRadiusFilter(ctx context.Context, chat_id int, radius string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddRadiusFilter]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_filters SET radius = ? WHERE chat_id = ?;`, radius, chat_id)
	if err != nil {
		return fmt.Errorf("[AddRadiusFilter]set values city in db error: %w", err)
	}
	return nil
}
func (s Storage) AddCategoryFilter(ctx context.Context, id int, category string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddCategoryFilter]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_filters SET category = ? WHERE id = ?;`, category, id)
	if err != nil {
		return fmt.Errorf("[AddCategoryFilter]set values city in db error: %w", err)
	}
	return nil
}

func (s Storage) AddMonitoringFilter(ctx context.Context, id int, url string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddMonitoringFilter]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_filters SET monitoring = ? WHERE id = ?;`, url, id)
	if err != nil {
		return fmt.Errorf("[AddMonitoringFilter]set values city in db error: %w", err)
	}
	return nil
}

func (s Storage) MonitoringByIDFilter(ctx context.Context, id int) (string, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return "", fmt.Errorf("[MonitoringByIDFilter]connection db error: %w", err)
	}
	defer conn.Close()
	var source string
	err = conn.GetContext(ctx, &source, `SELECT monitoring FROM data_filters WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return "", ErrNoRows
	}
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("[MonitoringByIDFilter]can't execute a request: %w", err)
	}

	return source, nil
}

func (s Storage) DeleteFilter(ctx context.Context, id int) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[DeleteFilter]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `DELETE FROM data_filters WHERE id=?;`, id)
	if err != nil {
		return fmt.Errorf("[DeleteFilter ]set values city in db error: %w", err)
	}
	return nil
}

type Filter struct {
	Id          string `db:"id"`
	Chat_id     int    `db:"chat_id"`
	Monitoring  string `db:"monitoring"`
	City        string `db:"city"`
	Radius      string `db:"radius"`
	Category    string `db:"category"`
	Filter_file string `db:"filter_file"`
}

func (s Storage) SelectAllFilter(ctx context.Context, chat_id int) ([]Filter, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, fmt.Errorf("[SelectAllFilter]connection db error: %w", err)
	}
	defer conn.Close()

	var sources []Filter
	if err := conn.SelectContext(ctx, &sources, `SELECT * FROM data_filters where chat_id = ?`, chat_id); err != nil {
		return nil, fmt.Errorf("[Sources]can't execute a request: %w", err)
	}

	return lo.Map(sources, func(source Filter, _ int) Filter { return Filter(source) }), nil
}

func (s Storage) AddFilterFile(ctx context.Context, id int, filterfile string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return fmt.Errorf("[AddFilterFile]connection db error: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `UPDATE data_filters SET filter_file = ? WHERE id = ?;`, filterfile, id)
	if err != nil {
		return fmt.Errorf("[AddFilterFile]set values city in db error: %w", err)
	}
	return nil
}

func (s Storage) SelectFilterFile(ctx context.Context, id int) (string, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return "", fmt.Errorf("[SelectFilterFile]connection db error: %w", err)
	}
	defer conn.Close()
	var source string
	err = conn.GetContext(ctx, &source, `SELECT filter_file FROM data_filters WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return "", ErrNoRows
	}
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("[SelectFilterFile]can't execute a request: %w", err)
	}

	return source, nil
}
