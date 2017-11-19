//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	. "github.com/portapps/portapps"
)

func init() {
	App.ID = "brave-portable"
	App.Name = "Brave"
	Init()
}

func main() {
	App.MainPath = FindElectronMainFolder("app-")
	App.DataPath = CreateFolder(PathJoin(App.RootDataPath, "AppData", "Roaming", "Brave"))
	App.Process = RootPathJoin("Brave.exe")
	App.Args = nil
	App.WorkingDir = App.MainPath

	OverrideEnv("USERPROFILE", App.RootDataPath)
	Launch()
}
