// Package zk8s provides helpers for working with Kubernetes.
package zk8s

import (
	"github.com/acorn-io/baaah/pkg/name"
)

var (
	// SafeName generates a name suitable for use in a k8s object.
	SafeName = name.SafeConcatName
)
