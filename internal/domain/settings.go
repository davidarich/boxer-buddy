package domain

type Settings struct {
	MultiboxGroups []MultiboxGroup
	InteropOptions InteropOptions
	UiOptions      UiOptions
}

func (cfg *Settings) GetGroupByName(name string) *MultiboxGroup {
	for i := range cfg.MultiboxGroups {
		if cfg.MultiboxGroups[i].Name == name {
			return &cfg.MultiboxGroups[i]
		}
	}
	return nil
}

func NewSettings() *Settings {
	return &Settings{}
}

// settings for a Game
type GameProfile struct {
	Name string
	Password string
	// client
	Path        string // location of the game client's main directory
	BinPath     string // subpath of binary relative to Path
	BinFileName string // filename of the game's main process
	StartCmd    string
	StartArgs   []string
}

// returns the complete filepath to a game client binary
// todo: this is written in a windows specific way and needs updated for cross platform support
func (gp *GameProfile) GetFullBinPath() string {
	return gp.Path + "\\" + gp.BinFileName
}

// settings of multiple GameProfiles
type MultiboxGroup struct {
	Name         string
	GameProfiles []GameProfile
}

// gets a unique list of game client binary filenames; useful for filtering processes
func (mg *MultiboxGroup) GetUniqueFilenames() (filenames []string) {
	present := map[string]bool{}
	for i := range mg.GameProfiles {
		if mg.GameProfiles[i].BinFileName == "" {
			continue
		}
		if !present[mg.GameProfiles[i].BinFileName] {
			present[mg.GameProfiles[i].BinFileName] = true
			filenames = append(filenames, mg.GameProfiles[i].BinFileName)
		}
	}

	return
}

// settings options for Interop server
type InteropOptions struct {
	host string
	port string
}

func (opts *InteropOptions) GetHost() string {
	host := opts.host
	if host == "" {
		host = "localhost"
	}
	return host
}

func (opts *InteropOptions) GetPort() string {
	port := opts.port
	if port == "" {
		port = "7778" // default
	}
	return port
}

func (opts *InteropOptions) GetAddress() string {
	return opts.GetHost() + ":" + opts.GetPort()
}

type UiOptions struct {
	host string
	port string
}

func (opts *UiOptions) GetAddress() string {
	host := opts.host
	if host == "" {
		host = "localhost"
	}
	port := opts.port
	if port == "" {
		port = "7777" // default
	}
	return host + ":" + port
}
