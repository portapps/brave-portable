//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "brave-portable"
	Papp.Name = "Brave"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = AppPathJoin("data")

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))
	roamingPath := CreateFolder(PathJoin(Papp.DataPath, "AppData", "Roaming", "Brave"))
	Log.Infof("Roaming path: %s", roamingPath)

	Papp.Process = PathJoin(Papp.AppPath, "Brave.exe")
	Papp.Args = nil
	Papp.WorkingDir = electronBinPath

	OverrideEnv("USERPROFILE", Papp.DataPath)
	Launch()
}
