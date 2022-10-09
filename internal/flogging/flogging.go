package flogging

import (
	logging "github.com/ipfs/go-log/v2"
)

var appVersion = "v1"
var log = logging.Logger("crawler").With("app_version", appVersion)
