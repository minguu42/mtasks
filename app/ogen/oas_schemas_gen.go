// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/url"
	"time"
)

type DeleteProjectBadRequest Error

func (*DeleteProjectBadRequest) deleteProjectRes() {}

type DeleteProjectInternalServerError Error

func (*DeleteProjectInternalServerError) deleteProjectRes() {}

// DeleteProjectNoContent is response for DeleteProject operation.
type DeleteProjectNoContent struct{}

func (*DeleteProjectNoContent) deleteProjectRes() {}

type DeleteProjectNotFound Error

func (*DeleteProjectNotFound) deleteProjectRes() {}

type DeleteProjectNotImplemented Error

func (*DeleteProjectNotImplemented) deleteProjectRes() {}

type DeleteProjectUnauthorized Error

func (*DeleteProjectUnauthorized) deleteProjectRes() {}

type DeleteTaskBadRequest Error

func (*DeleteTaskBadRequest) deleteTaskRes() {}

type DeleteTaskInternalServerError Error

func (*DeleteTaskInternalServerError) deleteTaskRes() {}

// DeleteTaskNoContent is response for DeleteTask operation.
type DeleteTaskNoContent struct{}

func (*DeleteTaskNoContent) deleteTaskRes() {}

type DeleteTaskNotFound Error

func (*DeleteTaskNotFound) deleteTaskRes() {}

type DeleteTaskNotImplemented Error

func (*DeleteTaskNotImplemented) deleteTaskRes() {}

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

type GetHealthNotImplemented Error

func (*GetHealthNotImplemented) getHealthRes() {}

type GetHealthOK struct {
	// MTasks APIのバージョン.
	Version string `json:"version"`
	// MTasks APIのリビジョン.
	Revision string `json:"revision"`
}

// GetVersion returns the value of Version.
func (s *GetHealthOK) GetVersion() string {
	return s.Version
}

// GetRevision returns the value of Revision.
func (s *GetHealthOK) GetRevision() string {
	return s.Revision
}

// SetVersion sets the value of Version.
func (s *GetHealthOK) SetVersion(val string) {
	s.Version = val
}

// SetRevision sets the value of Revision.
func (s *GetHealthOK) SetRevision(val string) {
	s.Revision = val
}

func (*GetHealthOK) getHealthRes() {}

type GetHealthServiceUnavailable Error

func (*GetHealthServiceUnavailable) getHealthRes() {}

type GetProjectsBadRequest Error

func (*GetProjectsBadRequest) getProjectsRes() {}

type GetProjectsInternalServerError Error

func (*GetProjectsInternalServerError) getProjectsRes() {}

type GetProjectsNotImplemented Error

func (*GetProjectsNotImplemented) getProjectsRes() {}

type GetProjectsUnauthorized Error

func (*GetProjectsUnauthorized) getProjectsRes() {}

type GetTasksBadRequest Error

func (*GetTasksBadRequest) getTasksRes() {}

type GetTasksInternalServerError Error

func (*GetTasksInternalServerError) getTasksRes() {}

type GetTasksNotFound Error

func (*GetTasksNotFound) getTasksRes() {}

type GetTasksNotImplemented Error

func (*GetTasksNotImplemented) getTasksRes() {}

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

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
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

type PatchProjectBadRequest Error

func (*PatchProjectBadRequest) patchProjectRes() {}

type PatchProjectInternalServerError Error

func (*PatchProjectInternalServerError) patchProjectRes() {}

type PatchProjectNotFound Error

func (*PatchProjectNotFound) patchProjectRes() {}

type PatchProjectNotImplemented Error

func (*PatchProjectNotImplemented) patchProjectRes() {}

type PatchProjectReq struct {
	Name OptString `json:"name"`
}

// GetName returns the value of Name.
func (s *PatchProjectReq) GetName() OptString {
	return s.Name
}

// SetName sets the value of Name.
func (s *PatchProjectReq) SetName(val OptString) {
	s.Name = val
}

type PatchProjectUnauthorized Error

func (*PatchProjectUnauthorized) patchProjectRes() {}

type PatchTaskBadRequest Error

func (*PatchTaskBadRequest) patchTaskRes() {}

type PatchTaskInternalServerError Error

func (*PatchTaskInternalServerError) patchTaskRes() {}

type PatchTaskNotFound Error

func (*PatchTaskNotFound) patchTaskRes() {}

type PatchTaskNotImplemented Error

func (*PatchTaskNotImplemented) patchTaskRes() {}

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

type PostProjectsBadRequest Error

func (*PostProjectsBadRequest) postProjectsRes() {}

type PostProjectsInternalServerError Error

func (*PostProjectsInternalServerError) postProjectsRes() {}

type PostProjectsNotImplemented Error

func (*PostProjectsNotImplemented) postProjectsRes() {}

type PostProjectsReq struct {
	Name string `json:"name"`
}

// GetName returns the value of Name.
func (s *PostProjectsReq) GetName() string {
	return s.Name
}

// SetName sets the value of Name.
func (s *PostProjectsReq) SetName(val string) {
	s.Name = val
}

type PostProjectsUnauthorized Error

func (*PostProjectsUnauthorized) postProjectsRes() {}

type PostTasksBadRequest Error

func (*PostTasksBadRequest) postTasksRes() {}

type PostTasksInternalServerError Error

func (*PostTasksInternalServerError) postTasksRes() {}

type PostTasksNotFound Error

func (*PostTasksNotFound) postTasksRes() {}

type PostTasksNotImplemented Error

func (*PostTasksNotImplemented) postTasksRes() {}

type PostTasksReq struct {
	Title string `json:"title"`
}

// GetTitle returns the value of Title.
func (s *PostTasksReq) GetTitle() string {
	return s.Title
}

// SetTitle sets the value of Title.
func (s *PostTasksReq) SetTitle(val string) {
	s.Title = val
}

type PostTasksUnauthorized Error

func (*PostTasksUnauthorized) postTasksRes() {}

// Ref: #/components/schemas/project
type Project struct {
	// プロジェクトID.
	ID int64 `json:"id"`
	// プロジェクト名.
	Name string `json:"name"`
	// 作成日時.
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時.
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *Project) GetID() int64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *Project) GetName() string {
	return s.Name
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Project) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Project) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *Project) SetID(val int64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Project) SetName(val string) {
	s.Name = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Project) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Project) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

func (*Project) patchProjectRes() {}

// ProjectHeaders wraps Project with response headers.
type ProjectHeaders struct {
	Location url.URL
	Response Project
}

// GetLocation returns the value of Location.
func (s *ProjectHeaders) GetLocation() url.URL {
	return s.Location
}

// GetResponse returns the value of Response.
func (s *ProjectHeaders) GetResponse() Project {
	return s.Response
}

// SetLocation sets the value of Location.
func (s *ProjectHeaders) SetLocation(val url.URL) {
	s.Location = val
}

// SetResponse sets the value of Response.
func (s *ProjectHeaders) SetResponse(val Project) {
	s.Response = val
}

func (*ProjectHeaders) postProjectsRes() {}

// Ref: #/components/schemas/projects
type Projects struct {
	// プロジェクト一覧.
	Projects []Project `json:"projects"`
}

// GetProjects returns the value of Projects.
func (s *Projects) GetProjects() []Project {
	return s.Projects
}

// SetProjects sets the value of Projects.
func (s *Projects) SetProjects(val []Project) {
	s.Projects = val
}

func (*Projects) getProjectsRes() {}

// Ref: #/components/schemas/task
type Task struct {
	// タスクID.
	ID int64 `json:"id"`
	// 紐づくプロジェクトのID.
	ProjectID int64 `json:"project_id"`
	// タイトル.
	Title string `json:"title"`
	// 完了日時.
	CompletedAt OptDateTime `json:"completedAt"`
	// 作成日時.
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時.
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *Task) GetID() int64 {
	return s.ID
}

// GetProjectID returns the value of ProjectID.
func (s *Task) GetProjectID() int64 {
	return s.ProjectID
}

// GetTitle returns the value of Title.
func (s *Task) GetTitle() string {
	return s.Title
}

// GetCompletedAt returns the value of CompletedAt.
func (s *Task) GetCompletedAt() OptDateTime {
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

// SetProjectID sets the value of ProjectID.
func (s *Task) SetProjectID(val int64) {
	s.ProjectID = val
}

// SetTitle sets the value of Title.
func (s *Task) SetTitle(val string) {
	s.Title = val
}

// SetCompletedAt sets the value of CompletedAt.
func (s *Task) SetCompletedAt(val OptDateTime) {
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
	Location url.URL
	Response Task
}

// GetLocation returns the value of Location.
func (s *TaskHeaders) GetLocation() url.URL {
	return s.Location
}

// GetResponse returns the value of Response.
func (s *TaskHeaders) GetResponse() Task {
	return s.Response
}

// SetLocation sets the value of Location.
func (s *TaskHeaders) SetLocation(val url.URL) {
	s.Location = val
}

// SetResponse sets the value of Response.
func (s *TaskHeaders) SetResponse(val Task) {
	s.Response = val
}

func (*TaskHeaders) postTasksRes() {}

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
