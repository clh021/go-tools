package compat

import "context"

type ICompatV1 interface {
	DetectRecordSearch(ctx context.Context, req *DetectRecordSearchReq) (res *DetectRecordSearchRes, err error)
}