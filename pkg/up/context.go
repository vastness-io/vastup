package up

// BuildContext contains the context for each component that is created from
// the component repository.
type BuildContext struct {
	RepoPath     string
	BinPath      string
	LogMountPath string
}
