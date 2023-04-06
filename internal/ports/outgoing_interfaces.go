package ports

import (
	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/domain"
)

type Settings interface {
	Get() (cfg *domain.Settings, err error)
}

type PersistentSettings interface {
	Load() (err error)
	Save() (err error)
}

type LogFactory interface {
	GetLogger() log.Logger
}

type Process interface {
	Filtered(filenames []string) (processes []*domain.Process, err error)
	Start(executablePath string, args []string, workingDir string) (err error)
	Stop(id int) (err error)
}
