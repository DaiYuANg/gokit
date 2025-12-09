package bblot_wrapper

import (
	"fmt"
	"log/slog"

	"github.com/samber/oops"
)

type Slogger struct {
	logger *slog.Logger
}

func NewSlogger(l *slog.Logger) *Slogger {
	if l == nil {
		l = slog.Default()
	}
	return &Slogger{logger: l}
}

// formatArgs: concatenate all args into one string.
func formatArgs(v ...interface{}) string {
	if len(v) == 0 {
		return ""
	}
	if len(v) == 1 {
		return fmt.Sprint(v[0])
	}
	return fmt.Sprint(v...)
}

// wrapErrorIfNeeded: if v[0] is error -> wrap with oops and insert context
func wrapErrorIfNeeded(v ...interface{}) (error, string) {
	if len(v) == 0 {
		return nil, ""
	}

	// error first style: logger.Error(err, "msg")
	if err, ok := v[0].(error); ok {
		msg := formatArgs(v[1:]...)
		return oops.With("msg", msg).Wrap(err), msg
	}

	// normal style: logger.Error("something failed")
	return nil, formatArgs(v...)
}

func (s *Slogger) Debug(v ...interface{}) {
	s.logger.Debug(formatArgs(v...))
}

func (s *Slogger) Debugf(format string, v ...interface{}) {
	s.logger.Debug(fmt.Sprintf(format, v...))
}

func (s *Slogger) Info(v ...interface{}) {
	s.logger.Info(formatArgs(v...))
}

func (s *Slogger) Infof(format string, v ...interface{}) {
	s.logger.Info(fmt.Sprintf(format, v...))
}

func (s *Slogger) Warning(v ...interface{}) {
	s.logger.Warn(formatArgs(v...))
}

func (s *Slogger) Warningf(format string, v ...interface{}) {
	s.logger.Warn(fmt.Sprintf(format, v...))
}

func (s *Slogger) Error(v ...interface{}) {
	if err, msg := wrapErrorIfNeeded(v...); err != nil {
		s.logger.Error(msg, slog.String("error", err.Error()))
		return
	}
	s.logger.Error(formatArgs(v...))
}

func (s *Slogger) Errorf(format string, v ...interface{}) {
	// explicit error in v?
	if len(v) > 0 {
		if err, ok := v[0].(error); ok {
			msg := fmt.Sprintf(format, v[1:]...)
			wrapped := oops.With("msg", msg).Wrap(err)
			s.logger.Error(msg, slog.String("error", wrapped.Error()))
			return
		}
	}
	s.logger.Error(fmt.Sprintf(format, v...))
}

func (s *Slogger) Fatal(v ...interface{}) {
	if err, msg := wrapErrorIfNeeded(v...); err != nil {
		s.logger.Error(msg, slog.String("error", err.Error()))
		panic("fatal: " + err.Error())
	}
	msg := formatArgs(v...)
	s.logger.Error(msg)
	panic("fatal: " + msg)
}

func (s *Slogger) Fatalf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	s.logger.Error(msg)
	panic("fatal: " + msg)
}

func (s *Slogger) Panic(v ...interface{}) {
	if err, msg := wrapErrorIfNeeded(v...); err != nil {
		s.logger.Error(msg, slog.String("error", err.Error()))
		panic(err)
	}
	msg := formatArgs(v...)
	s.logger.Error(msg)
	panic(msg)
}

func (s *Slogger) Panicf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	s.logger.Error(msg)
	panic(msg)
}
