package database

import (
	"context"
	"fmt"
	"time"

	"github.com/minguu42/mtasks/pkg/entity"
	"github.com/minguu42/mtasks/pkg/ttime"
)

func (db *DB) CreateTask(ctx context.Context, projectID string, title string) (*entity.Task, error) {
	now := ttime.Now(ctx)
	t := entity.Task{
		ID:          db.idGenerator.Generate(),
		ProjectID:   projectID,
		Title:       title,
		CompletedAt: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := db.conn(ctx).Create(&t).Error; err != nil {
		return nil, fmt.Errorf("gormDB.Create failed: %w", err)
	}
	return &t, nil
}

func (db *DB) GetTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	var t entity.Task
	if err := db.conn(ctx).First(&t, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("gormDB.First failed: %w", err)
	}
	return &t, nil
}

func (db *DB) GetTasksByProjectID(ctx context.Context, projectID string, sort string, limit, offset int) ([]*entity.Task, error) {
	ts := make([]*entity.Task, 0, limit)
	if err := db.conn(ctx).Where("project_id = ?", projectID).
		Order(generateOrderByClause(sort)).Limit(limit).Offset(offset).Find(&ts).Error; err != nil {
		return nil, fmt.Errorf("gormDB.Find failed: %w", err)
	}
	return ts, nil
}

func (db *DB) UpdateTask(ctx context.Context, id string, completedAt *time.Time, updatedAt time.Time) error {
	t := entity.Task{ID: id}
	if err := db.conn(ctx).Model(&t).Updates(entity.Task{CompletedAt: completedAt, UpdatedAt: updatedAt}).Error; err != nil {
		return fmt.Errorf("gormDB.Updates failed: %w", err)
	}
	return nil
}

func (db *DB) DeleteTask(ctx context.Context, id string) error {
	if err := db.conn(ctx).Delete(&entity.Task{ID: id}).Error; err != nil {
		return fmt.Errorf("gormDB.Delete failed: %w", err)
	}
	return nil
}
