package pkglogger

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	Success string = "success"
	Fail    string = "fail"
)

type PkgLogger struct {
	err          error
	logger       zerolog.Logger
	loggingEvent *zerolog.Event
	identifier   string
	result       string
	taxYear      int32
}

func NewPkgLogger(pkg, method string) *PkgLogger {
	p := pkg
	if p == "" {
		p = "unknown"
	}
	if method != "" {
		p = p + "." + method
	}
	return &PkgLogger{
		logger: log.With().Str("pkg", p).Logger(),
		result: Success,
	}
}

func (l *PkgLogger) Debug() *PkgLogger {

	l.loggingEvent = l.logger.Debug()
	return l
}

func (l *PkgLogger) Err(err error, msg string) {
	l.err = errors.Wrap(err, msg)
	l.Send()
}

func (l *PkgLogger) Errf(err error, msg string, args ...interface{}) {
	msgActual := fmt.Sprintf(msg, args...)
	l.err = errors.Wrap(err, msgActual)
	l.Send()
}

func (l *PkgLogger) Error() *PkgLogger {
	l.loggingEvent = l.logger.Error()
	l.Result(Fail)
	return l
}

func (l *PkgLogger) Fail() *PkgLogger {
	l.Result(Fail)
	return l
}

func (l *PkgLogger) Fatal() *PkgLogger {
	l.loggingEvent = l.logger.Fatal()
	l.Result(Fail)
	return l
}

func (l *PkgLogger) Identifier(identifier string) *PkgLogger {
	l.identifier = identifier
	return l
}

func (l *PkgLogger) Info() *PkgLogger {
	l.loggingEvent = l.logger.Info()
	return l
}

func (l *PkgLogger) Success() *PkgLogger {
	l.Result(Success)
	return l
}

func (l *PkgLogger) Result(result string) *PkgLogger {
	l.result = result
	return l
}

func (l *PkgLogger) Msg(msg string) {
	l.marshallOutcome().Msg(msg)
	l.reset()
}

func (l *PkgLogger) Msgf(msg string, args ...interface{}) {
	l.marshallOutcome().Msgf(msg, args...)
	l.reset()
}

func (l *PkgLogger) Panic() *PkgLogger {
	l.loggingEvent = l.logger.Panic()
	l.Result(Fail)
	return l
}

func (l *PkgLogger) Send() {
	l.marshallOutcome().Send()
	l.reset()
}

func (l *PkgLogger) TaxYear(taxYear int32) *PkgLogger {
	l.taxYear = taxYear
	return l
}

func (l *PkgLogger) Trace() *PkgLogger {
	l.loggingEvent = l.logger.Trace()
	return l
}

func (l *PkgLogger) Warning() *PkgLogger {
	l.loggingEvent = l.logger.Warn()
	return l
}

func (l *PkgLogger) marshallOutcome() *zerolog.Event {
	outcome := zerolog.Dict().Str("result", l.result)
	if l.err != nil {
		outcome.Stack()
		outcome = outcome.Err(l.err)
	}
	if l.identifier != "" {
		outcome = outcome.Str("identifier", l.identifier)
	}
	if l.taxYear > 0 {
		outcome = outcome.Int32("tax_year", l.taxYear)
	}
	return l.loggingEvent.Dict("outcome", outcome)
}

func (l *PkgLogger) reset() {
	l.err = nil
	l.loggingEvent = nil
	l.identifier = ""
	l.taxYear = 0
	l.result = Success
}
