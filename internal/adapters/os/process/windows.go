package process

import (
	"errors"
	"os"
	"os/exec"
	"unsafe"

	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/domain"
	"golang.org/x/sys/windows"
)

// adapter for interacting with processes on a Windows OS
type WindowsAdapter struct {
	logger log.Logger
}

// get a process slice filtered by filenames
func (w *WindowsAdapter) Filtered(filenames []string) (processes []*domain.Process, err error) {
	// get all processes
	allProcesses, err := w.get()
	if err != nil {
		return
	}

	// filter processes to only processes with a binary filename matching filenames slice
	for i := range allProcesses {
		if filter(allProcesses[i], filenames) {
			processes = append(processes, allProcesses[i])
		}
	}

	// add path data to each Process
	for i := range processes {
		w.setProcessPaths(processes[i])
	}

	return
}

// gets all processes
func (w *WindowsAdapter) get() (processes []*domain.Process, err error) {
	// creates a snapshot of running processes
	hSnapProc, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return
	}
	defer windows.CloseHandle(hSnapProc)

	// prepare ProcessEntry32, Size is required to be set for memory allocation
	pe := windows.ProcessEntry32{}
	pe.Size = uint32(unsafe.Sizeof(windows.ProcessEntry32{}))

	pcount := 0
	for {
		pcount++
		// safeguard against infinite loop
		if pcount == 1000 {
			err = errors.New("reached limit while getting processes")
			return
		}
		if pcount == 0 {
			err = windows.Process32First(hSnapProc, &pe)
			if err != nil {
				w.logger.Info("first process was not found")
				return
			}

		} else {
			err = windows.Process32Next(hSnapProc, &pe)
			if err != nil {
				w.logger.Info("no more processes were found")
				// clear err, as it is only used to indicate the end of the process list rather than a failure
				err = nil
				break
			}
		}

		// map ProcessEntry32 to *domain.Process
		process := domain.NewProcess()
		process.Id = int(pe.ProcessID)
		process.ParentId = int(pe.ParentProcessID)
		process.FileName = windows.UTF16ToString(pe.ExeFile[:])
		processes = append(processes, process)

		continue
	}
	return
}

// sets path of a process on the provided *domain.Pointer
func (w *WindowsAdapter) setProcessPaths(p *domain.Process) {
	defer func() {
		if r := recover(); r != nil {
			w.logger.Error("panic: ", r)
		}
	}()

	w.logger.Info("PID: ", uint32(p.Id))
	hProc, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION, false, uint32(p.Id))
	if err != nil {
		w.logger.Info("failed to open process handle")
		w.logger.Info(err)
		return
	}
	defer windows.CloseHandle(hProc)
	w.logger.Info("process opened")

	size := uint32(windows.MAX_LONG_PATH)
	var exeName [windows.MAX_LONG_PATH]uint16
	err = windows.QueryFullProcessImageName(hProc, 0, &exeName[0], &size)
	if err != nil {
		w.logger.Error(windows.GetLastError())
	}

	p.Path = windows.UTF16ToString(exeName[:])
	w.logger.Info(p.Path)
	w.logger.Info("OK")
}

// starts a process
func (w *WindowsAdapter) Start(executablePath string, args []string, workingDir string) (err error) {
	processArgs := []string{"/C", "start", "/D", workingDir, "/b", executablePath}
	processArgs = append(processArgs, args...)
	cmd := exec.Command("cmd.exe", processArgs...)
	err = cmd.Run()
	if err != nil {
		w.logger.Error(err)
	}
	return
}

// stops a process
func (w *WindowsAdapter) Stop(id int) (err error) {
	proc, _ := os.FindProcess(id)
	proc.Kill()
	return
}

func NewWindowsAdapter(logger log.Logger) *WindowsAdapter {
	return &WindowsAdapter{
		logger: logger,
	}
}
