package log

import "go.uber.org/zap"

type ZapLogFactory struct {
	logger *zap.SugaredLogger
}

func (lf *ZapLogFactory) Logger() Logger {
	return lf.logger
}

func NewZapLogFactory() *ZapLogFactory {
	logger, _ := zap.NewProduction()

	return &ZapLogFactory{
		// performance isn't a major concern, so always use the sugared logger
		logger: logger.Sugar(),
	}
}
