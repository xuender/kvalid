//go:build go_json

package json

import json "github.com/goccy/go-json"

// Marshal is exported by json package.
var Marshal = json.Marshal // nolint: gochecknoglobals
