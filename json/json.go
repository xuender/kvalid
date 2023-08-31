//go:build !jsoniter && !go_json && !(sonic && avx && (linux || windows || darwin) && amd64)

package json

import "encoding/json"

// Marshal is exported by json package.
var Marshal = json.Marshal // nolint: gochecknoglobals
