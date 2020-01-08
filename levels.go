package onelogplus

const (
	_ = iota

	FINEST uint32 = 1 << iota
	FINE
	FINER
	DEBUG
	CONFIG
	INFO
	WARN
	ERROR
	SEVERE
	FATAL
)

// Levels is the mapping between int log levels and their string value
var Levels = make([]string, 2048)
var levelsJSON = make([][]byte, 2048)
var levelKey = "level"

func init() {
	Levels[FINEST] = "finest"
	Levels[FINER] = "finer"
	Levels[FINE] = "fine"
	Levels[CONFIG] = "config"
	Levels[INFO] = "info"
	Levels[DEBUG] = "debug"
	Levels[WARN] = "warn"
	Levels[ERROR] = "error"
	Levels[FATAL] = "fatal"
	Levels[SEVERE] = "severe"
	genLevelSlices()
}

func genLevelSlices() {
	levelsJSON[FINEST] = []byte(`{"` + levelKey + `":"` + Levels[FINEST] + `","` + msgKey + `":`)
	levelsJSON[FINER] = []byte(`{"` + levelKey + `":"` + Levels[FINER] + `","` + msgKey + `":`)
	levelsJSON[FINE] = []byte(`{"` + levelKey + `":"` + Levels[FINE] + `","` + msgKey + `":`)
	levelsJSON[INFO] = []byte(`{"` + levelKey + `":"` + Levels[INFO] + `","` + msgKey + `":`)
	levelsJSON[DEBUG] = []byte(`{"` + levelKey + `":"` + Levels[DEBUG] + `","` + msgKey + `":`)
	levelsJSON[WARN] = []byte(`{"` + levelKey + `":"` + Levels[WARN] + `","` + msgKey + `":`)
	levelsJSON[ERROR] = []byte(`{"` + levelKey + `":"` + Levels[ERROR] + `","` + msgKey + `":`)
	levelsJSON[FATAL] = []byte(`{"` + levelKey + `":"` + Levels[FATAL] + `","` + msgKey + `":`)
	levelsJSON[SEVERE] = []byte(`{"` + levelKey + `":"` + Levels[SEVERE] + `","` + msgKey + `":`)
}
