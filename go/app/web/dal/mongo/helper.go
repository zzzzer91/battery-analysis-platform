package mongo

import (
	"battery-analysis-platform/app/web/constant"
	"context"
)

func newTimeoutCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), constant.MongoCtxTimeout)
	return ctx
}
