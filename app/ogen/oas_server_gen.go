// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DeleteProject implements deleteProject operation.
	//
	// プロジェクトを削除する.
	//
	// DELETE /projects/{projectID}
	DeleteProject(ctx context.Context, params DeleteProjectParams) (DeleteProjectRes, error)
	// DeleteTask implements deleteTask operation.
	//
	// タスクを削除する.
	//
	// DELETE /projects/{projectID}/tasks/{taskID}
	DeleteTask(ctx context.Context, params DeleteTaskParams) (DeleteTaskRes, error)
	// GetHealth implements getHealth operation.
	//
	// サーバの状態を取得する.
	//
	// GET /health
	GetHealth(ctx context.Context) (GetHealthRes, error)
	// GetProjects implements getProjects operation.
	//
	// 作成日時の降順で取得する。.
	//
	// GET /projects
	GetProjects(ctx context.Context, params GetProjectsParams) (GetProjectsRes, error)
	// GetTasks implements getTasks operation.
	//
	// 作成日時の降順で取得する。.
	//
	// GET /projects/{projectID}/tasks
	GetTasks(ctx context.Context, params GetTasksParams) (GetTasksRes, error)
	// PatchProject implements patchProject operation.
	//
	// プロジェクトを更新する.
	//
	// PATCH /projects/{projectID}
	PatchProject(ctx context.Context, req *PatchProjectReq, params PatchProjectParams) (PatchProjectRes, error)
	// PatchTask implements patchTask operation.
	//
	// タスクを更新する.
	//
	// PATCH /projects/{projectID}/tasks/{taskID}
	PatchTask(ctx context.Context, req *PatchTaskReq, params PatchTaskParams) (PatchTaskRes, error)
	// PostProjects implements PostProjects operation.
	//
	// 新しいプロジェクトを作成する.
	//
	// POST /projects
	PostProjects(ctx context.Context, req *PostProjectsReq, params PostProjectsParams) (PostProjectsRes, error)
	// PostTasks implements PostTasks operation.
	//
	// 新しいタスクを作成する.
	//
	// POST /projects/{projectID}/tasks
	PostTasks(ctx context.Context, req *PostTasksReq, params PostTasksParams) (PostTasksRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}