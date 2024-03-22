package ci

import "time"

type WebHookInjection struct {
	Id    int
	URL   string
	Owner string
	Repo  string
}
type GitlabCreateHook struct {
	ID                       int         `json:"id"`
	URL                      string      `json:"url"`
	CreatedAt                time.Time   `json:"created_at"`
	PushEvents               bool        `json:"push_events"`
	TagPushEvents            bool        `json:"tag_push_events"`
	MergeRequestsEvents      bool        `json:"merge_requests_events"`
	RepositoryUpdateEvents   bool        `json:"repository_update_events"`
	EnableSslVerification    bool        `json:"enable_ssl_verification"`
	ProjectID                int         `json:"project_id"`
	IssuesEvents             bool        `json:"issues_events"`
	ConfidentialIssuesEvents bool        `json:"confidential_issues_events"`
	NoteEvents               bool        `json:"note_events"`
	ConfidentialNoteEvents   interface{} `json:"confidential_note_events"`
	PipelineEvents           bool        `json:"pipeline_events"`
	WikiPageEvents           bool        `json:"wiki_page_events"`
	JobEvents                bool        `json:"job_events"`
	PushEventsBranchFilter   interface{} `json:"push_events_branch_filter"`
}
