package db

import "context"

func AddMessage(ctx context.Context, message *sql.Message) {
	return DB.WithContext(ctx).Create(message).Error
}
