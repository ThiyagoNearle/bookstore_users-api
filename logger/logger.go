package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger // logger is a library placed in zap package & it is private variable because it has small letter l, we need all this logs for only our purposess
)

func init() {
	logConfig := zap.Config{ // configure
		OutputPaths: []string{"stdout"}, // in this way we are not looging in any log files instaed just logging into this ("standardout")
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",                       // "level": "info about "
			TimeKey:      "time",                        // "time": 2019-11-10T23:00:23-0700"
			MessageKey:   "msg",                         // "msg": "This is the logging line"
			EncodeTime:   zapcore.ISO8601TimeEncoder,    // for formatting the time
			EncodeLevel:  zapcore.LowercaseLevelEncoder, // formatting for level:,
			EncodeCaller: zapcore.ShortCallerEncoder,    // something related to directory
		},
	}

	var err error // when programm calls the init function, either it gives *zap.Looger or error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// later we discuss about the tags

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...) // here Info is a buit in function already defined
	log.Sync()             // By using our logger object (Log) we are calling the buitin Info function & we pass the message & tags
} // finally we sync our Log object , so it holds the message and tags

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}
