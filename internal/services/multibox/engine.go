package multibox

import (
	"encoding/json"
	"time"

	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/adapters/os/path"
	"github.com/davidarich/boxer-buddy/internal/domain"
	"github.com/davidarich/boxer-buddy/internal/ports"
)

// core functionality of the application
type Engine struct {
	IsRunning       bool
	activeProfiles  []*ActiveProfile
	clientFilenames []string
	ticker          *time.Ticker

	logger                  log.Logger
	process                 ports.Process
	window                  ports.WindowManager
	activeProfileDispatcher ports.InteropWriter
}

// syncs Engine.activeProfiles state with UI
func (e *Engine) dispatchActiveProfileUiUpdate() {
	message, err := json.Marshal(e.activeProfiles)
	if err != nil {
		e.logger.Error("error preparing activeProfile state message")
		return
	}
	if e.activeProfileDispatcher == nil {
		e.logger.Error("no interop dispatcher is set")
		return
	}
	e.activeProfileDispatcher.Write(message)
}

// focuses a game window to the foreground
func (e *Engine) FocusGame(profileName string) {
	// todo: implement FocusGame

	// check if the game is running

	// get pid from profile name

	// get all windows

	// find window with matching pid

	// focus the window
}

func (e *Engine) getActiveProfileByName(profileName string) *ActiveProfile {
	for i := range e.activeProfiles {
		if e.activeProfiles[i] == nil {
			continue
		}
		if e.activeProfiles[i].GameProfile.Name != profileName {
			continue
		}
		return e.activeProfiles[i]
	}
	return nil
}

// one cycle for engine
func (e *Engine) run() {
	e.logger.Info("started run")
	processes, err := e.process.Filtered(e.clientFilenames)
	if err != nil {
		e.logger.Error(err)
		return
	}

	// match ActiveProfiles to running processes
	needsUpdate := false
	for i := range e.activeProfiles {
		isMissing := true
		for j := range processes {
			// check if Process filename matches current ActiveProfile
			if processes[j].FileName != e.activeProfiles[i].GameProfile.BinFileName {
				continue
			}
			expandedPath, err := path.Expand(e.activeProfiles[i].GameProfile.GetFullBinPath())
			if err != nil {
				expandedPath = e.activeProfiles[i].GameProfile.GetFullBinPath()
			}
			// check if Process path matches current ActiveProfile
			if processes[j].Path != expandedPath {
				continue
			}

			e.logger.Info("found game process ", processes[j].FileName, " with id ", processes[j].Id, " at path ", processes[j].Path)
			isMissing = false
			if e.activeProfiles[i].GameClient.Process == nil {
				e.activeProfiles[i].GameClient.Process = processes[j]
				needsUpdate = true
			}
		}
		// remove stale process references
		if isMissing && e.activeProfiles[i].GameClient.Process != nil {
			e.activeProfiles[i].GameClient.Process = nil
			needsUpdate = true
		}
	}

	// sync changes with UI
	if needsUpdate {
		e.dispatchActiveProfileUiUpdate()
	}
	e.logger.Info("finished run")
}

// starts the main engine loop
func (e *Engine) Start(mbGroup *domain.MultiboxGroup) {
	if e.IsRunning {
		e.logger.Warn("engine is already running and will not be started")
		return
	}
	e.IsRunning = true
	e.clientFilenames = mbGroup.GetUniqueFilenames()
	e.logger.Info("engine started")

	// convert MultiboxProfile into ActiveProfiles
	for i := range mbGroup.GameProfiles {
		activeProfile := NewActiveProfile(mbGroup.GameProfiles[i])
		e.activeProfiles = append(e.activeProfiles, activeProfile)
	}
	e.dispatchActiveProfileUiUpdate()

	// setup ticker
	interval := time.Second * 10
	e.ticker = time.NewTicker(interval)
	defer e.ticker.Stop()

	for {
		<-e.ticker.C
		go e.run()
	}
}

// starts a game client process
func (e *Engine) StartGame(profileName string) {
	activeProfile := e.getActiveProfileByName(profileName)
	if activeProfile.GameClient == nil {
		e.logger.Error("did not find game client for profile: ", profileName)
		return
	}
	// check if game is running
	if activeProfile.GameClient.IsRunning() {
		e.logger.Warn("game for profile ", profileName, " is already running, nothing to do")
		return
	}

	// game isn't running, so start process
	cmd := activeProfile.GameProfile.GetFullBinPath()
	if activeProfile.GameProfile.StartCmd != "" {
		cmd = activeProfile.GameProfile.Path + "\\" + activeProfile.GameProfile.StartCmd
	}

	e.logger.Info("launching process")
	e.process.Start(cmd, activeProfile.GameProfile.StartArgs, activeProfile.GameProfile.Path)
}

// gets the state of the engine
func (e *Engine) Status() {
	e.dispatchActiveProfileUiUpdate()
}

// shuts down the Engine and performs cleanup
func (e *Engine) Stop() (err error) {
	e.activeProfiles = []*ActiveProfile{}
	e.clientFilenames = []string{}
	e.ticker.Stop()
	e.IsRunning = false
	return
}

// stops a game client process
func (e *Engine) StopGame(profileName string) {
	activeProfile := e.getActiveProfileByName(profileName)
	if activeProfile.GameClient == nil {
		e.logger.Error("did not find game client for profile: ", profileName)
		return
	}
	// check if game is running
	if !activeProfile.GameClient.IsRunning() {
		e.logger.Warn("game for profile ", profileName, " is not running, nothing to do")
		return
	}

	// stop process
	e.process.Stop(activeProfile.GameClient.Process.Id)
}

// set the dispatcher after Engine instantiation to resolve circular dependancy
func (e *Engine) SetActiveProfileDispatcher(w ports.InteropWriter) {
	e.activeProfileDispatcher = w
}

func NewEngine(logger log.Logger, process ports.Process, window ports.WindowManager) *Engine {
	return &Engine{
		logger:  logger,
		process: process,
		window:  window,
	}
}
