package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"io"
)

func checkFormat(b io.ReadCloser, n io.ReadCloser, ctx reqcontext.RequestContext) err