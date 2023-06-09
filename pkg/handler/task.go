package handler

import (
	"context"

	"github.com/go-faster/errors"
	"github.com/minguu42/mtasks/gen/ogen"
	"github.com/minguu42/mtasks/pkg/entity"
	"github.com/minguu42/mtasks/pkg/logging"
	"github.com/minguu42/mtasks/pkg/ttime"
	"gorm.io/gorm"
)

// CreateTask は POST /projects/{projectID}/tasks に対応するハンドラ
func (h *Handler) CreateTask(ctx context.Context, req *ogen.CreateTaskReq, params ogen.CreateTaskParams) (*ogen.Task, error) {
	u, ok := ctx.Value(userKey{}).(*entity.User)
	if !ok {
		return nil, errUnauthorized
	}

	p, err := h.Repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errProjectNotFound
		}
		logging.Errorf(ctx, "repository.GetProjectByID failed: %v", err)
		return nil, errInternalServerError
	}
	if !u.HasProject(p) {
		logging.Errorf(ctx, "user does not have the project")
		return nil, errProjectNotFound
	}

	t, err := h.Repository.CreateTask(ctx, params.ProjectID, req.Title)
	if err != nil {
		logging.Errorf(ctx, "repository.Create failed: %v", err)
		return nil, errInternalServerError
	}

	return newTaskResponse(t), nil
}

// ListTasks は GET /projects/{projectID}/tasks に対応するハンドラ
func (h *Handler) ListTasks(ctx context.Context, params ogen.ListTasksParams) (*ogen.Tasks, error) {
	u, ok := ctx.Value(userKey{}).(*entity.User)
	if !ok {
		return nil, errUnauthorized
	}

	p, err := h.Repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errProjectNotFound
		}
		logging.Errorf(ctx, "repository.GetProjectByID failed: %v", err)
		return nil, errInternalServerError
	}
	if !u.HasProject(p) {
		logging.Errorf(ctx, "user does not have the project")
		return nil, errProjectNotFound
	}

	ts, err := h.Repository.GetTasksByProjectID(ctx, p.ID, string(params.Sort.Or(ogen.ListTasksSortMinusCreatedAt)), params.Limit.Or(10)+1, params.Offset.Or(0))
	if err != nil {
		logging.Errorf(ctx, "repository.GetTasksByProjectID failed: %v", err)
		return nil, errInternalServerError
	}

	hasNext := false
	if len(ts) == params.Limit.Or(10)+1 {
		hasNext = true
		ts = ts[:params.Limit.Or(10)]
	}

	return &ogen.Tasks{
		Tasks:   newTasksResponse(ts),
		HasNext: hasNext,
	}, nil
}

// UpdateTask は PATCH /projects/{projectID}/tasks/{taskID} に対応するハンドラ
func (h *Handler) UpdateTask(ctx context.Context, req *ogen.UpdateTaskReq, params ogen.UpdateTaskParams) (*ogen.Task, error) {
	if !req.IsCompleted.IsSet() {
		logging.Errorf(ctx, "value contains nothing")
		return nil, errBadRequest
	}

	u, ok := ctx.Value(userKey{}).(*entity.User)
	if !ok {
		return nil, errUnauthorized
	}

	p, err := h.Repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errProjectNotFound
		}
		logging.Errorf(ctx, "repository.GetProjectByID failed: %v", err)
		return nil, errInternalServerError
	}
	if !u.HasProject(p) {
		logging.Errorf(ctx, "user does not have the project")
		return nil, errProjectNotFound
	}
	t, err := h.Repository.GetTaskByID(ctx, params.TaskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errTaskNotFound
		}
		logging.Errorf(ctx, "repository.GetTaskByID failed: %v", err)
		return nil, errInternalServerError
	}
	if !p.ContainsTask(t) {
		logging.Errorf(ctx, "project does not contain the task")
		return nil, errTaskNotFound
	}

	now := ttime.Now(ctx)
	if req.IsCompleted.Value {
		t.CompletedAt = &now
	} else {
		t.CompletedAt = nil
	}
	t.UpdatedAt = now
	if err := h.Repository.UpdateTask(ctx, params.TaskID, t.CompletedAt, t.UpdatedAt); err != nil {
		logging.Errorf(ctx, "repository.UpdateTask failed: %v", err)
		return nil, errInternalServerError
	}

	return newTaskResponse(t), nil
}

// DeleteTask は DELETE /projects/{projectID}/tasks/{taskID} に対応するハンドラ
func (h *Handler) DeleteTask(ctx context.Context, params ogen.DeleteTaskParams) error {
	u, ok := ctx.Value(userKey{}).(*entity.User)
	if !ok {
		return errUnauthorized
	}

	p, err := h.Repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errProjectNotFound
		}
		logging.Errorf(ctx, "repository.GetProjectByID failed: %v", err)
		return errInternalServerError
	}
	if !u.HasProject(p) {
		logging.Errorf(ctx, "user does not have the project")
		return errProjectNotFound
	}
	t, err := h.Repository.GetTaskByID(ctx, params.TaskID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errTaskNotFound
		}
		logging.Errorf(ctx, "repository.GetTaskByID failed: %v", err)
		return errInternalServerError
	}
	if !p.ContainsTask(t) {
		logging.Errorf(ctx, "project does not contain the task")
		return errTaskNotFound
	}

	if err := h.Repository.DeleteTask(ctx, t.ID); err != nil {
		logging.Errorf(ctx, "repository.DeleteTask failed: %v", err)
		return errInternalServerError
	}

	return nil
}
