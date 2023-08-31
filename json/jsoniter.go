//go:build jsoniter

package json

import jsoniter "github.com/json-iterator/go"

// Marshal is exported by json package.
var Marshal = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal // nolint: gochecknoglobals
