package handler

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/minguu42/mtasks/gen/mock"
	"github.com/minguu42/mtasks/gen/ogen"
	"github.com/minguu42/mtasks/pkg/entity"
	"gorm.io/gorm"
)

func TestHandler_CreateProject(t *testing.T) {
	type args struct {
		ctx context.Context
		req *ogen.CreateProjectReq
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(r *mock.MockRepository)
		want          *ogen.Project
		wantErr       error
	}{
		{
			name: "プロジェクト1を作成する",
			args: args{
				ctx: mockCtx,
				req: &ogen.CreateProjectReq{Name: "プロジェクト1"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().CreateProject(mockCtx, "01DXF6DT000000000000000000", "プロジェクト1").
					Return(&entity.Project{
						ID:        "01DXF6DT000000000000000000",
						UserID:    "01DXF6DT000000000000000000",
						Name:      "プロジェクト1",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					}, nil)
			},
			want: &ogen.Project{
				ID:        "01DXF6DT000000000000000000",
				Name:      "プロジェクト1",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
		},
		{
			name:          "コンテキストからユーザを取得できない場合はエラーを返す",
			args:          args{ctx: context.Background()},
			prepareMockFn: func(r *mock.MockRepository) {},
			want:          nil,
			wantErr:       errUnauthorized,
		},
		{
			name: "データベースへの操作がエラーを返す場合はエラーを返す",
			args: args{
				ctx: mockCtx,
				req: &ogen.CreateProjectReq{Name: "プロジェクト1"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().CreateProject(mockCtx, "01DXF6DT000000000000000000", "プロジェクト1").
					Return(nil, errors.New("some error"))
			},
			want:    nil,
			wantErr: errInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := mock.NewMockRepository(c)
			tt.prepareMockFn(r)
			h := &Handler{Repository: r}

			got, err := h.CreateProject(tt.args.ctx, tt.args.req)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("h.CreateProject() error want '%v', but '%v'", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("h.CreateProject() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandler_ListProjects(t *testing.T) {
	type args struct {
		ctx    context.Context
		params ogen.ListProjectsParams
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(r *mock.MockRepository)
		want          *ogen.Projects
		wantErr       error
	}{
		{
			name: "プロジェクト一覧を取得する",
			args: args{ctx: mockCtx},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectsByUserID(mockCtx, "01DXF6DT000000000000000000", "-createdAt", 11, 0).
					Return([]*entity.Project{
						{
							ID:        "01DXF6DT000000000000000000",
							UserID:    "01DXF6DT000000000000000000",
							Name:      "プロジェクト1",
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
						{
							ID:        "01DXF6DT000000000000000001",
							UserID:    "01DXF6DT000000000000000000",
							Name:      "プロジェクト2",
							CreatedAt: time.Time{},
							UpdatedAt: time.Time{},
						},
					}, nil)
			},
			want: &ogen.Projects{
				Projects: []ogen.Project{
					{
						ID:        "01DXF6DT000000000000000000",
						Name:      "プロジェクト1",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
					{
						ID:        "01DXF6DT000000000000000001",
						Name:      "プロジェクト2",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
				},
				HasNext: false,
			},
			wantErr: nil,
		},
		{
			name:          "コンテキストからユーザを取得できない場合はエラーを返す",
			args:          args{ctx: context.Background()},
			prepareMockFn: func(r *mock.MockRepository) {},
			want:          nil,
			wantErr:       errUnauthorized,
		},
		{
			name: "repository.GetProjectsByUserIDがエラーを返す場合はエラーを返す",
			args: args{ctx: mockCtx},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectsByUserID(mockCtx, "01DXF6DT000000000000000000", "-createdAt", 11, 0).
					Return(nil, errors.New("some error"))
			},
			want:    nil,
			wantErr: errInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := mock.NewMockRepository(c)
			tt.prepareMockFn(r)
			h := &Handler{Repository: r}

			got, err := h.ListProjects(tt.args.ctx, tt.args.params)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("h.ListProjects() error want '%v', but '%v'", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("h.ListProjects() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandler_UpdateProject(t *testing.T) {
	type args struct {
		ctx    context.Context
		req    *ogen.UpdateProjectReq
		params ogen.UpdateProjectParams
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(r *mock.MockRepository)
		want          *ogen.Project
		wantErr       error
	}{
		{
			name: "プロジェクト1を変更する",
			args: args{
				ctx:    mockCtx,
				req:    &ogen.UpdateProjectReq{Name: ogen.OptString{Value: "新プロジェクト1", Set: true}},
				params: ogen.UpdateProjectParams{ProjectID: "01DXF6DT000000000000000000"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000000").Return(&entity.Project{
					ID:        "01DXF6DT000000000000000000",
					UserID:    "01DXF6DT000000000000000000",
					Name:      "プロジェクト1",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, nil)
				r.EXPECT().UpdateProject(mockCtx, "01DXF6DT000000000000000000", "新プロジェクト1", time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)).
					Return(nil)
			},
			want: &ogen.Project{
				ID:        "01DXF6DT000000000000000000",
				Name:      "新プロジェクト1",
				CreatedAt: time.Time{},
				UpdatedAt: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "コンテキストからユーザを取得できない場合はエラーを返す",
			args: args{
				ctx: context.Background(),
				req: &ogen.UpdateProjectReq{Name: ogen.OptString{Value: "新プロジェクト1", Set: true}},
			},
			prepareMockFn: func(r *mock.MockRepository) {},
			want:          nil,
			wantErr:       errUnauthorized,
		},
		{
			name: "指定したプロジェクトが見つからない場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				req:    &ogen.UpdateProjectReq{Name: ogen.OptString{Value: "新プロジェクト2", Set: true}},
				params: ogen.UpdateProjectParams{ProjectID: "01DXF6DT000000000000000001"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000001").Return(nil, gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: errProjectNotFound,
		},
		{
			name: "repository.GetProjectByIDが何らかのエラーを返す場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				req:    &ogen.UpdateProjectReq{Name: ogen.OptString{Value: "新プロジェクト1", Set: true}},
				params: ogen.UpdateProjectParams{ProjectID: "01DXF6DT000000000000000000"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000000").Return(nil, errors.New("some error"))
			},
			want:    nil,
			wantErr: errInternalServerError,
		},
		{
			name: "指定したプロジェクトをユーザが保持していない場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				req:    &ogen.UpdateProjectReq{Name: ogen.OptString{Value: "新プロジェクト2", Set: true}},
				params: ogen.UpdateProjectParams{ProjectID: "01DXF6DT000000000000000001"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000001").Return(&entity.Project{
					ID:        "01DXF6DT000000000000000001",
					UserID:    "01DXF6DT000000000000000001",
					Name:      "プロジェクト2",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, nil)
			},
			want:    nil,
			wantErr: errProjectNotFound,
		},
		{
			name: "repository.UpdateProjectが何らかのエラーを返す場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				req:    &ogen.UpdateProjectReq{Name: ogen.OptString{Value: "新プロジェクト1", Set: true}},
				params: ogen.UpdateProjectParams{ProjectID: "01DXF6DT000000000000000000"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000000").Return(&entity.Project{
					ID:        "01DXF6DT000000000000000000",
					UserID:    "01DXF6DT000000000000000000",
					Name:      "プロジェクト1",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, nil)
				r.EXPECT().UpdateProject(mockCtx, "01DXF6DT000000000000000000", "新プロジェクト1", time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)).
					Return(errors.New("some error"))
			},
			want:    nil,
			wantErr: errInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := mock.NewMockRepository(c)
			tt.prepareMockFn(r)
			h := &Handler{Repository: r}

			got, err := h.UpdateProject(tt.args.ctx, tt.args.req, tt.args.params)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("h.UpdateProject() error want '%v', but '%v'", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("h.UpdateProject() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandler_DeleteProject(t *testing.T) {
	type args struct {
		ctx    context.Context
		params ogen.DeleteProjectParams
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(r *mock.MockRepository)
		want          error
	}{
		{
			name: "プロジェクト1を削除する",
			args: args{
				ctx:    mockCtx,
				params: ogen.DeleteProjectParams{ProjectID: "01DXF6DT000000000000000000"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000000").Return(&entity.Project{
					ID:        "01DXF6DT000000000000000000",
					UserID:    "01DXF6DT000000000000000000",
					Name:      "プロジェクト1",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, nil)
				r.EXPECT().DeleteProject(mockCtx, "01DXF6DT000000000000000000").Return(nil)
			},
			want: nil,
		},
		{
			name:          "コンテキストからユーザを取得できない場合はエラーを返す",
			args:          args{ctx: context.Background()},
			prepareMockFn: func(r *mock.MockRepository) {},
			want:          errUnauthorized,
		},
		{
			name: "指定したプロジェクトが見つからない場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				params: ogen.DeleteProjectParams{ProjectID: "01DXF6DT000000000000000001"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000001").Return(nil, gorm.ErrRecordNotFound)
			},
			want: errProjectNotFound,
		},
		{
			name: "repository.GetProjectByIDが何らかのエラーを返す場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				params: ogen.DeleteProjectParams{ProjectID: "01DXF6DT000000000000000000"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000000").Return(nil, errors.New("some error"))
			},
			want: errInternalServerError,
		},
		{
			name: "指定したプロジェクトをユーザが保持していない場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				params: ogen.DeleteProjectParams{ProjectID: "01DXF6DT000000000000000001"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000001").Return(&entity.Project{
					ID:        "01DXF6DT000000000000000001",
					UserID:    "01DXF6DT000000000000000001",
					Name:      "プロジェクト2",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, nil)
			},
			want: errProjectNotFound,
		},
		{
			name: "repository.UpdateProjectが何らかのエラーを返す場合はエラーを返す",
			args: args{
				ctx:    mockCtx,
				params: ogen.DeleteProjectParams{ProjectID: "01DXF6DT000000000000000000"},
			},
			prepareMockFn: func(r *mock.MockRepository) {
				r.EXPECT().GetProjectByID(mockCtx, "01DXF6DT000000000000000000").Return(&entity.Project{
					ID:        "01DXF6DT000000000000000000",
					UserID:    "01DXF6DT000000000000000000",
					Name:      "プロジェクト1",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, nil)
				r.EXPECT().DeleteProject(mockCtx, "01DXF6DT000000000000000000").Return(errors.New("some error"))
			},
			want: errInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := mock.NewMockRepository(c)
			tt.prepareMockFn(r)
			h := &Handler{Repository: r}

			if err := h.DeleteProject(tt.args.ctx, tt.args.params); tt.want != err {
				t.Errorf("h.DeleteProject() want '%v', but '%v'", tt.want, err)
			}
		})
	}
}
