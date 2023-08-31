//go:build sonic && avx && (linux || windows || darwin) && amd64

package json

import "github.com/bytedance/sonic"

// Marshal is exported by json package.
var Marshal = sonic.ConfigStd.Marshal // nolint: gochecknoglobals
