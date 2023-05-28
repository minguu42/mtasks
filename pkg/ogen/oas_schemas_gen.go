// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"time"
)

type CreateTasksBadRequest Error

func (*CreateTasksBadRequest) createTasksRes() {}

type CreateTasksReq struct {
	Title string `json:"title"`
}

// GetTitle returns the value of Title.
func (s *CreateTasksReq) GetTitle() string {
	return s.Title
}

// SetTitle sets the value of Title.
func (s *CreateTasksReq) SetTitle(val string) {
	s.Title = val
}

type CreateTasksUnauthorized Error

func (*CreateTasksUnauthorized) createTasksRes() {}

type DeleteTaskBadRequest Error

func (*DeleteTaskBadRequest) deleteTaskRes() {}

// DeleteTaskNoContent is response for DeleteTask operation.
type DeleteTaskNoContent struct{}

func (*DeleteTaskNoContent) deleteTaskRes() {}

type DeleteTaskNotFound Error

func (*DeleteTaskNotFound) deleteTaskRes() {}

type DeleteTaskUnauthorized Error

func (*DeleteTaskUnauthorized) deleteTaskRes() {}

// Ref: #/components/schemas/error
type Error struct {
	// ユーザ向けの大まかなエラーの説明.
	Message string `json:"message"`
	// 開発者向けの詳細なエラーの説明.
	Debug string `json:"debug"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// GetDebug returns the value of Debug.
func (s *Error) GetDebug() string {
	return s.Debug
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// SetDebug sets the value of Debug.
func (s *Error) SetDebug(val string) {
	s.Debug = val
}

func (*Error) getHealthRes() {}

// GetHealthOK is response for GetHealth operation.
type GetHealthOK struct{}

func (*GetHealthOK) getHealthRes() {}

type GetTasksBadRequest Error

func (*GetTasksBadRequest) getTasksRes() {}

type GetTasksUnauthorized Error

func (*GetTasksUnauthorized) getTasksRes() {}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type PatchTaskBadRequest Error

func (*PatchTaskBadRequest) patchTaskRes() {}

type PatchTaskNotFound Error

func (*PatchTaskNotFound) patchTaskRes() {}

type PatchTaskReq struct {
	IsCompleted OptBool `json:"isCompleted"`
}

// GetIsCompleted returns the value of IsCompleted.
func (s *PatchTaskReq) GetIsCompleted() OptBool {
	return s.IsCompleted
}

// SetIsCompleted sets the value of IsCompleted.
func (s *PatchTaskReq) SetIsCompleted(val OptBool) {
	s.IsCompleted = val
}

type PatchTaskUnauthorized Error

func (*PatchTaskUnauthorized) patchTaskRes() {}

// Ref: #/components/schemas/task
type Task struct {
	// タスクID.
	ID int64 `json:"id"`
	// タイトル.
	Title string `json:"title"`
	// 完了日時.
	CompletedAt time.Time `json:"completedAt"`
	// 作成日時.
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時.
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *Task) GetID() int64 {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Task) GetTitle() string {
	return s.Title
}

// GetCompletedAt returns the value of CompletedAt.
func (s *Task) GetCompletedAt() time.Time {
	return s.CompletedAt
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Task) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Task) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Task) SetID(val int64) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Task) SetTitle(val string) {
	s.Title = val
}

// SetCompletedAt sets the value of CompletedAt.
func (s *Task) SetCompletedAt(val time.Time) {
	s.CompletedAt = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Task) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Task) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

func (*Task) patchTaskRes() {}

// TaskHeaders wraps Task with response headers.
type TaskHeaders struct {
	Location OptString
	Response Task
}

// GetLocation returns the value of Location.
func (s *TaskHeaders) GetLocation() OptString {
	return s.Location
}

// GetResponse returns the value of Response.
func (s *TaskHeaders) GetResponse() Task {
	return s.Response
}

// SetLocation sets the value of Location.
func (s *TaskHeaders) SetLocation(val OptString) {
	s.Location = val
}

// SetResponse sets the value of Response.
func (s *TaskHeaders) SetResponse(val Task) {
	s.Response = val
}

func (*TaskHeaders) createTasksRes() {}

// Ref: #/components/schemas/tasks
type Tasks struct {
	// タスク一覧.
	Tasks []Task `json:"tasks"`
}

// GetTasks returns the value of Tasks.
func (s *Tasks) GetTasks() []Task {
	return s.Tasks
}

// SetTasks sets the value of Tasks.
func (s *Tasks) SetTasks(val []Task) {
	s.Tasks = val
}

func (*Tasks) getTasksRes() {}