package utils

import (
	"fmt"
	"time"
)

var maps = map[string]time.Time{}

func StartTrack(name string) {
	maps[name] = time.Now()
}

func StopTrack(name string) {
	Info(fmt.Sprintf(`"%s": operation took %v.`, name, time.Now().Sub(maps[name])))
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	Info(fmt.Sprintf(`"%s": operation took %v`, name, elapsed))
	delete(maps, name)
}
