package badger_wrapper

import (
	"fmt"
	"log/slog"

	"github.com/samber/oops"
)

type BadgerSLogger struct {
	logger *slog.Logger
}

func NewBadgerSLogger(l *slog.Logger) *BadgerSLogger {
	return &BadgerSLogger{logger: l}
}

// ---- helper ----

func wrapIfError(v ...interface{}) error {
	if len(v) == 0 {
		return nil
	}
	last, ok := v[len(v)-1].(error)
	if !ok {
		return nil
	}
	return oops.Wrapf(last, "badger error")
}

func makeAttrs(v ...interface{}) []any {
	attrs := make([]any, 0, len(v)*2)
	for i, a := range v {
		attrs = append(attrs, fmt.Sprintf("arg_%d", i), a)
	}
	return attrs
}

// ---- Logger Implementation ----

func (b *BadgerSLogger) Error(v ...interface{}) {
	msg := fmt.Sprint(v...)
	if err := wrapIfError(v...); err != nil {
		b.logger.Error(msg, "error", err, slog.String("oops_stack", err.Error()))
		return
	}
	b.logger.Error(msg, makeAttrs(v...)...)
}

func (b *BadgerSLogger) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if err := wrapIfError(v...); err != nil {
		b.logger.Error(msg, "error", err, slog.String("oops_stack", err.Error()))
		return
	}
	b.logger.Error(msg)
}

func (b *BadgerSLogger) Warning(v ...interface{}) {
	msg := fmt.Sprint(v...)
	if err := wrapIfError(v...); err != nil {
		b.logger.Warn(msg, "error", err, slog.String("oops_stack", err.Error()))
		return
	}
	b.logger.Warn(msg, makeAttrs(v...)...)
}

func (b *BadgerSLogger) Warningf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	if err := wrapIfError(v...); err != nil {
		b.logger.Warn(msg, "error", err, slog.String("oops_stack", err.Error()))
		return
	}
	b.logger.Warn(msg)
}

func (b *BadgerSLogger) Info(v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Info(msg, makeAttrs(v...)...)
}

func (b *BadgerSLogger) Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.logger.Info(msg)
}

func (b *BadgerSLogger) Debug(v ...interface{}) {
	msg := fmt.Sprint(v...)
	b.logger.Debug(msg, makeAttrs(v...)...)
}

func (b *BadgerSLogger) Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	b.logger.Debug(msg)
}
