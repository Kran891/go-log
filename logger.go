package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func init() {

}
func setDefualtLogLevel() {
	if os.Getenv("ENV") == "prod" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

type ZeroLogger interface {
	Error() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Debug() *zerolog.Event
	With() zerolog.Context
}

var loggerInstance ZeroLogger

func SetLogger(baseLogger ZeroLogger) {
	loggerInstance = baseLogger
}

func logger() ZeroLogger {
	if loggerInstance != nil {
		return loggerInstance
	}
	zerolog.TimeFieldFormat = time.RFC1123Z
	baseLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	SetLogger(&baseLogger)
	return loggerInstance
}

var (
	invalidArgMessage      = "Invalid arg: %s"
	invalidArgValueMessage = "Invalid value for agrument :- %s : %v"
	missingArgMessage      = "Missing arg :- %s"
)

func InvaligArg(argName string) {
	logger().Error().Msgf(invalidArgMessage, argName)
}

func InvaligArgValue(argName, argValue string) {
	logger().Error().Msgf(invalidArgValueMessage, argName, argValue)
}

func MissingArg(argName string) {
	logger().Error().Msgf(missingArgMessage, missingArgMessage)
}

func Error(err error, msg string) {
	logger().Error().Msgf("%v : %v", msg, err)
}

func Errorf(msgFormat string, v ...interface{}) {
	logger().Error().Msgf(msgFormat, v...)
}

func ErrorWithFields(err error, msg string, fields map[string]interface{}) {
	log := logger().With().Fields(fields).Logger()
	log.Error().Msgf("%v : %v ", msg, err)
}

func Infof(msg string, v ...interface{}) {
	logger().Info().Msgf(msg, v...)
}

func InfoWithFields(msg string, fields map[string]interface{}) {
	log := logger().With().Fields(fields).Logger()
	log.Info().Msg(msg)
}

func Warnf(msg string, v ...interface{}) {
	logger().Warn().Msgf(msg, v...)
}

func WarnWithFields(msg string, fields map[string]interface{}) {
	log := logger().With().Fields(fields).Logger()
	log.Warn().Msg(msg)
}

func Debugf(msg string, v ...interface{}) {
	logger().Debug().Msgf(msg, v...)
}

func DebugWithFields(msg string, fields map[string]interface{}) {
	log := logger().With().Fields(fields).Logger()
	log.Debug().Msg(msg)
}
