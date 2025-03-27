package log

import (
	"context"
	"reflect"

	"github.com/DOGTT/dm-api-server/internal/utils"
	"github.com/davecgh/go-spew/spew"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap/zapcore"
)

func log(ctx context.Context, lvl zapcore.Level, msg string, keysAndValues ...any) {
	l := otelzap.Ctx(ctx).Logger()
	for i := 1; i < len(keysAndValues); i += 2 {
		// 指针类型使用spew.Sdump()进行格式化
		if reflect.ValueOf(keysAndValues[i]).Kind() != reflect.Ptr {
			keysAndValues[i] = spew.Sdump(keysAndValues[i])
		}
	}
	if tc := utils.GetClaimFromContext(ctx); tc != nil {
		keysAndValues = append(keysAndValues, "uid", tc.UId)
	}
	l.Sugar().Logw(lvl, msg, keysAndValues...)
}

func D(ctx context.Context, msg string, keysAndValues ...any) {
	if otelzap.L().Level() > zapcore.DebugLevel {
		return
	}
	log(ctx, zapcore.DebugLevel, msg, keysAndValues...)
}

func E(ctx context.Context, msg string, err error, keysAndValues ...any) {
	keysAndValues = append(keysAndValues, "error", err.Error())
	log(ctx, zapcore.ErrorLevel, msg, keysAndValues...)
}
