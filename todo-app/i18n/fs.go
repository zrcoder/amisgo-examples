package i18n

import (
	_ "embed"
	"encoding/json"
)

//go:embed en-US.json
var EnUS json.RawMessage

//go:embed zh-CN.json
var ZhCN json.RawMessage
