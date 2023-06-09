package database

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/minguu42/mtasks/gen/mock"
	"github.com/minguu42/mtasks/pkg/entity"
)

var taskCmpOpt = cmpopts.IgnoreFields(entity.Task{}, "ID", "CreatedAt", "UpdatedAt")

func TestDB_CreateTask(t *testing.T) {
	type args struct {
		ctx       context.Context
		projectID string
		title     string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(g *mock.MockIDGenerator)
		want          *entity.Task
		wantErr       error
	}{
		{
			name: "新タスクを作成する",
			args: args{
				ctx:       context.Background(),
				projectID: "01DXF6DT000000000000000000",
				title:     "新タスク",
			},
			prepareMockFn: func(g *mock.MockIDGenerator) {
				g.EXPECT().Generate().Return("01DXF6DT000000000000000002")
			},
			want: &entity.Task{
				ID:          "01DXF6DT000000000000000002",
				ProjectID:   "01DXF6DT000000000000000000",
				Title:       "新タスク",
				CompletedAt: nil,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := mock.NewMockIDGenerator(gomock.NewController(t))
			tt.prepareMockFn(g)
			testDB.SetIDGenerator(g)

			if err := testDB.Begin(); err != nil {
				t.Fatalf("testDB.Begin failed: %s", err)
			}
			defer testDB.Rollback()

			got, err := testDB.CreateTask(tt.args.ctx, tt.args.projectID, tt.args.title)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("testDB.CreateTask error want '%v', but '%v'", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got, taskCmpOpt); diff != "" {
				t.Errorf("testDB.CreateTask mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDB_GetTaskByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Task
		wantErr error
	}{
		{
			name: "タスク1を取得する",
			args: args{
				ctx: context.Background(),
				id:  "01DXF6DT000000000000000000",
			},
			want: &entity.Task{
				ID:          "01DXF6DT000000000000000000",
				ProjectID:   "01DXF6DT000000000000000000",
				Title:       "タスク1",
				CompletedAt: nil,
				CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "存在しないタスクを指定した場合はエラーを返す",
			args: args{
				ctx: context.Background(),
				id:  "01DXF6DT000000000000000002",
			},
			want:    nil,
			wantErr: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDB.GetTaskByID(tt.args.ctx, tt.args.id)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("testDB.GetTaskByID error want '%v', but '%v'", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("testDB.GetTaskByID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDB_GetTasksByProjectID(t *testing.T) {
	completedAt := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	type args struct {
		ctx       context.Context
		projectID string
		sort      string
		limit     int
		offset    int
	}
	tests := []struct {
		name    string
		args    args
		want    []*entity.Task
		wantErr error
	}{
		{
			name: "タスク一覧を取得する",
			args: args{
				ctx:       context.Background(),
				projectID: "01DXF6DT000000000000000000",
				sort:      "title",
				limit:     11,
				offset:    0,
			},
			want: []*entity.Task{
				{
					ID:          "01DXF6DT000000000000000000",
					ProjectID:   "01DXF6DT000000000000000000",
					Title:       "タスク1",
					CompletedAt: nil,
					CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          "01DXF6DT000000000000000001",
					ProjectID:   "01DXF6DT000000000000000000",
					Title:       "タスク2",
					CompletedAt: &completedAt,
					CreatedAt:   time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "タスクを1つだけ取得する",
			args: args{
				ctx:       context.Background(),
				projectID: "01DXF6DT000000000000000000",
				sort:      "title",
				limit:     1,
				offset:    0,
			},
			want: []*entity.Task{
				{
					ID:          "01DXF6DT000000000000000000",
					ProjectID:   "01DXF6DT000000000000000000",
					Title:       "タスク1",
					CompletedAt: nil,
					CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDB.GetTasksByProjectID(tt.args.ctx, tt.args.projectID, tt.args.sort, tt.args.limit, tt.args.offset)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("testDB.GetTasksByProjectID error want '%v', but '%v'", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("testDB.GetTasksByProjectID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDB_UpdateTask(t *testing.T) {
	completedAt := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	type args struct {
		ctx         context.Context
		id          string
		completedAt *time.Time
		updatedAt   time.Time
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "タスク1を更新する",
			args: args{
				ctx:         context.Background(),
				id:          "01DXF6DT000000000000000000",
				completedAt: &completedAt,
				updatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDB.Begin(); err != nil {
				t.Fatalf("testDB.Begin failed: %s", err)
			}
			defer testDB.Rollback()

			if err := testDB.UpdateTask(tt.args.ctx, tt.args.id, tt.args.completedAt, tt.args.updatedAt); (tt.want == nil) != (err == nil) {
				t.Errorf("testDB.UpdateTask want '%v', but '%v'", tt.want, err)
			}
		})
	}
}

func TestDB_DeleteTask(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "タスク1を削除する",
			args: args{
				ctx: context.Background(),
				id:  "01DXF6DT000000000000000000",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDB.Begin(); err != nil {
				t.Fatalf("testDB.Begin failed: %s", err)
			}
			defer testDB.Rollback()

			if err := testDB.DeleteTask(tt.args.ctx, tt.args.id); (tt.want == nil) != (err == nil) {
				t.Errorf("testDB.DeleteTask want '%v', but '%v'", tt.want, err)
			}
		})
	}
}
