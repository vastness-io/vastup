// Package cri contains common container client interfaces.
package cri

import (
	"context"
)

// ContainerRuntime is the common interface that all container engines
// must implement.
type ContainerRuntime interface {
	PullImage(context.Context) error
	BuildImage(context.Context) error
	RunImage(context.Context) error
	RebuildImage(context.Context) error
}

// ContainerImage holds metadata about the image used by the resource.
type ContainerImage struct {
	Name     string `json:"name"`
	Registry string `json:"registry"`
	BaseSHA  string `json:"base_sha"`
	Tag      string `json:"tag"`
}
