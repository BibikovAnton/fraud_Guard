

package antifraud_v1

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
)


type Labeler struct {
	attrs []attribute.KeyValue
}


func (l *Labeler) Add(attrs ...attribute.KeyValue) {
	l.attrs = append(l.attrs, attrs...)
}


func (l *Labeler) AttributeSet() attribute.Set {
	return attribute.NewSet(l.attrs...)
}

type labelerContextKey struct{}






func LabelerFromContext(ctx context.Context) (*Labeler, bool) {
	if l, ok := ctx.Value(labelerContextKey{}).(*Labeler); ok {
		return l, true
	}
	return &Labeler{}, false
}

func contextWithLabeler(ctx context.Context, l *Labeler) context.Context {
	return context.WithValue(ctx, labelerContextKey{}, l)
}
