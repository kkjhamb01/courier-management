package tracing

import (
	"context"
	"encoding/hex"
	"github.com/getsentry/sentry-go"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"regexp"
	"time"
)

type TracingInterceptor struct {
	WaitForDelivery bool
	Timeout time.Duration
	Repanic bool
	ReportOn func(error) bool
}

type Tracing interface {
}

func (tr *TracingInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}

		md, _ := metadata.FromIncomingContext(ctx)
		span := sentry.StartSpan(ctx, "grpc.server", tr.ContinueFromGrpcMetadata(md))
		ctx = span.Context()
		defer span.Finish()

		hub.Scope().SetExtra("requestBody", req)
		hub.Scope().SetTransaction(info.FullMethod)
		defer tr.recoverWithSentry(hub, ctx)

		resp, err := handler(ctx, req)
		if err != nil && tr.ReportOn(err) {
			tags := grpc_tags.Extract(ctx)
			for k, v := range tags.Values() {
				hub.Scope().SetTag(k, v.(string))
			}

			hub.CaptureException(err)
		}

		return resp, err
	}
}

func (tr *TracingInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream,
		info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := stream.Context()

		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}

		md, _ := metadata.FromIncomingContext(ctx)
		span := sentry.StartSpan(ctx, "grpc.server", tr.ContinueFromGrpcMetadata(md))
		ctx = span.Context()
		defer span.Finish()

		hub.Scope().SetExtra("requestBody", stream)
		hub.Scope().SetTransaction(info.FullMethod)
		defer tr.recoverWithSentry(hub, ctx)

		err := handler(srv, stream)
		if err != nil && tr.ReportOn(err) {
			tags := grpc_tags.Extract(ctx)
			for k, v := range tags.Values() {
				hub.Scope().SetTag(k, v.(string))
			}

			hub.CaptureException(err)
		}

		return err
	}
}

func (tr *TracingInterceptor) recoverWithSentry(hub *sentry.Hub, ctx context.Context) {
	if err := recover(); err != nil {
		eventID := hub.RecoverWithContext(ctx, err)
		if eventID != nil && tr.WaitForDelivery {
			hub.Flush(tr.Timeout)
		}

		if tr.Repanic {
			panic(err)
		}
	}
}

func (tr *TracingInterceptor) ContinueFromGrpcMetadata(md metadata.MD) sentry.SpanOption {
	return func(s *sentry.Span) {
		if md == nil {
			return
		}
		trace, ok := md["sentry-trace"]
		if !ok {
			return
		}
		if len(trace) != 1 {
			return
		}
		if trace[0] == "" {
			return
		}
		tr.updateFromSentryTrace(s, []byte(trace[0]))
	}
}

var sentryTracePattern = regexp.MustCompile(`^([[:xdigit:]]{32})-([[:xdigit:]]{16})(?:-([01]))?$`)

func (tr *TracingInterceptor) updateFromSentryTrace(s *sentry.Span, header []byte) {
	m := sentryTracePattern.FindSubmatch(header)
	if m == nil {
		// no match
		return
	}
	_, _ = hex.Decode(s.TraceID[:], m[1])
	_, _ = hex.Decode(s.ParentSpanID[:], m[2])
	if len(m[3]) != 0 {
		switch m[3][0] {
		case '0':
			s.Sampled = sentry.SampledFalse
		case '1':
			s.Sampled = sentry.SampledTrue
		}
	}
}

type reporter func(error) bool

func ReportAlways(error) bool {
	return true
}

func ReportOnCodes(cc ...codes.Code) reporter {
	return func(err error) bool {
		for i := range cc {
			if status.Code(err) == cc[i] {
				return true
			}
		}
		return false
	}
}

func NewTracingInterceptor() *TracingInterceptor {
	return &TracingInterceptor{
		WaitForDelivery: true,
		Timeout: 1,
		Repanic: true,
		ReportOn: ReportAlways,
	}
}
