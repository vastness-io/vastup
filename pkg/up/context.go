package up

// BuildContext contains the context for each component that is created from
// the component repository.
type BuildContext struct {
	RepoPath     string `json:"repo_path,omitempty"`
	BinPath      string `json:"bin_path,omitempty"`
	LogMountPath string `json:"log_mount_path,omitempty"`
}

// Config holds the configuration of the application and loads each components specific
// context. If a component context is not available, then the component container image
// will be pulled from the directory.
type Config struct {
	CoordinatorContext *BuildContext `json:"coordinator_context,omitempty"`
	VCSWebHookContext  *BuildContext `json:"vcs_webhook_context,omitempty"`
	LinguistContext    *BuildContext `json:"linguist_context,omitempty"`
}
