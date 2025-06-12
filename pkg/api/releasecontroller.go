package api

import (
	"time"
)

// Copied from release controller API: https://github.com/openshift/release-controller/blob/main/pkg/release-controller/types.go
// APIReleaseInfo encapsulates the release verification results and upgrade history for a release tag.
type APIReleaseInfo struct {
	// Name is the name of the release tag.
	Name string `json:"name"`
	// Phase is the phase of the release tag.
	Phase string `json:"phase"`
	// Results is the status of the release verification jobs for this release tag
	Results *VerificationJobsSummary `json:"results,omitempty"`
	// UpgradesTo is the list of UpgradeHistory "to" this release tag
	UpgradesTo []UpgradeHistory `json:"upgradesTo,omitempty"`
	//UpgradesFrom is the list of UpgradeHistory "from" this release tag
	UpgradesFrom []UpgradeHistory `json:"upgradesFrom,omitempty"`
	//ChangeLog is the html representation of the changes included in this release tag
	ChangeLog []byte `json:"changeLog,omitempty"`
	//ChangeLogJson is the json representation of the changes included in this release tag
	ChangeLogJson ChangeLog `json:"changeLogJson,omitempty"`
}

type VerificationStatus struct {
	State   string `json:"state"`
	URL     string `json:"url"`
	Retries int    `json:"retries,omitempty"`
	// TransitionTime *metav1.Time `json:"transitionTime,omitempty"`
}

type VerificationStatusMap map[string]*VerificationStatus

// VerificationJobsSummary an organized, by job type, collection of VerificationStatusMap objects
type VerificationJobsSummary struct {
	BlockingJobs  VerificationStatusMap `json:"blockingJobs,omitempty"`
	InformingJobs VerificationStatusMap `json:"informingJobs,omitempty"`
	PendingJobs   VerificationStatusMap `json:"pendingJobs,omitempty"`
}

type UpgradeResult struct {
	State string `json:"state"`
	URL   string `json:"url"`
}

type UpgradeHistory struct {
	From string
	To   string

	Success int
	Failure int
	Total   int

	History map[string]UpgradeResult
}

// ChangeLog represents the data structure that oc returns when providing a changelog in JSON format
// TODO: This is being carried from changes in openshift/oc.  These changes should be removed if/when we bump up our k8s dependencies up to the latest/greatest version.  We're currently pinned at: v0.24.2
type ChangeLog struct {
	From ChangeLogReleaseInfo `json:"from"`
	To   ChangeLogReleaseInfo `json:"to"`

	Components    []ChangeLogComponentInfo `json:"components,omitempty"`
	NewImages     []ChangeLogImageInfo     `json:"newImages,omitempty"`
	RemovedImages []ChangeLogImageInfo     `json:"removedImages,omitempty"`
	RebuiltImages []ChangeLogImageInfo     `json:"rebuiltImages,omitempty"`
	UpdatedImages []ChangeLogImageInfo     `json:"updatedImages,omitempty"`
}

type ChangeLogReleaseInfo struct {
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	//	Digest       digest.Digest `json:"digest"`
	PromotedFrom string `json:"promotedFrom,omitempty"`
}

type ChangeLogComponentInfo struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	VersionUrl string `json:"versionUrl,omitempty"`
	From       string `json:"from,omitempty"`
	FromUrl    string `json:"fromUrl,omitempty"`
	DiffUrl    string `json:"diffUrl,omitempty"`
}

type ChangeLogImageInfo struct {
	Name          string       `json:"name"`
	Path          string       `json:"path"`
	ShortCommit   string       `json:"shortCommit,omitempty"`
	Commit        string       `json:"commit,omitempty"`
	ImageRef      string       `json:"imageRef,omitempty"`
	Commits       []CommitInfo `json:"commits,omitempty"`
	FullChangeLog string       `json:"fullChangeLog,omitempty"`
}

type CommitInfo struct {
	Bugs      map[string]string `json:"bugs,omitempty"`
	Issues    map[string]string `json:"issues,omitempty"`
	Subject   string            `json:"subject,omitempty"`
	PullID    int               `json:"pullID,omitempty"`
	PullURL   string            `json:"pullURL,omitempty"`
	CommitID  string            `json:"commitID,omitempty"`
	CommitURL string            `json:"commitURL,omitempty"`
}

// End of releaseController API
