package database

import "context"

type Database interface {
	// AddChatID(ctx context.Context, chat_id int) error
	// AddWaitMessage(ctx context.Context, chat_id int) error
	// DeleteWaitMessage(ctx context.Context, chat_id int) error
	// WaitMessage(ctx context.Context, chat_id int) (bool, error)
	AddChatIDFilters(ctx context.Context, chat_id int) (int, error)
	AddMonitoringFilter(ctx context.Context, id int, url string) error
	MonitoringByIDFilter(ctx context.Context, id int) (string, error)
	AddCityFilter(ctx context.Context, id int, city string) error
	AddRadiusFilter(ctx context.Context, id int, radius string) error
	AddCategoryFilter(ctx context.Context, id int, category string) error
	DeleteFilter(ctx context.Context, id int) error
	SelectAllFilter(ctx context.Context, chat_id int) ([]Filter, error)
	AddFilterFile(ctx context.Context, id int, filterfile string) error
	SelectFilterFile(ctx context.Context, id int) (string, error)
	SelectFilterToFilterFile(ctx context.Context, filterFile string) (Filter, error)
}
