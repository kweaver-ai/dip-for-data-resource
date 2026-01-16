package trace_util

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/kweaver-ai/TelemetrySDK-Go/exporter/v2/ar_trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func StartSpan(ctx context.Context) (newCtx context.Context, span trace.Span) {
	pc, file, linkNo, ok := runtime.Caller(1)
	if ok {
		funcPaths := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		spanName := funcPaths[len(funcPaths)-1]
		newCtx, span = ar_trace.Tracer.Start(ctx, spanName)
		span.SetAttributes(attribute.String("func path", fmt.Sprintf("%s:%v", file, linkNo)))
	} else {
		newCtx, span = ar_trace.Tracer.Start(ctx, "runtime.Caller failed")
	}

	return
}

func TranceMiddleware(ctx context.Context, spanName string, fn func(ctx context.Context)) {
	ctx, span := ar_trace.Tracer.Start(ctx, spanName)
	defer span.End()

	fn(ctx)

}

//valid, errs := TranceValidator(context.back(), "",1, form_validator.BindQueryAndValid)

func TranceValidator(ctx context.Context, spanName string, p1 interface{}, fn func(ctx context.Context, v interface{}) (bool, error)) (bool, error) {
	_, span := ar_trace.Tracer.Start(ctx, spanName)
	defer span.End()
	return fn(ctx, p1)
}

func TranceP1R2[P1 any, R1 any, R2 any](ctx context.Context, spanName string, p1 P1, fn func(context.Context, P1) (R1, R2)) (r1 R1, r2 R2) {
	ctx, span := ar_trace.Tracer.Start(ctx, spanName)
	defer span.End()
	return fn(ctx, p1)
}

func TranceP1R1(ctx context.Context, spanName string, fn func(ctx context.Context)) {

}
