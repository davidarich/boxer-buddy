package process

import (
	"os"

	"github.com/davidarich/boxer-buddy/internal/domain"
)

func filter(p *domain.Process, filenames []string) bool {
	// guard against null pointer
	if p == nil {
		return false
	}

	// skip the kernel process
	if p.Id == 0 {
		return false
	}

	// skip the current process
	if p.Id == os.Getpid() {
		return false
	}

	// skip process with a name not matching expected filename
	for i := range filenames {
		if p.FileName == filenames[i] {
			return true
		}
	}

	return false
}
