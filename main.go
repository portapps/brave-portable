//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"os"

	. "github.com/portapps/portapps"
)

func init() {
	Papp.ID = "brave-portable"
	Papp.Name = "Brave"
	Init()
}

func main() {
	Papp.AppPath = AppPathJoin("app")
	Papp.DataPath = CreateFolder(AppPathJoin("data"))

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))

	Papp.Process = PathJoin(Papp.AppPath, "Brave.exe")
	Papp.Args = []string{
		"--user-data-dir=" + PathJoin(Papp.DataPath, "UserDataDir"),
		"--no-default-browser-check",
	}
	Papp.WorkingDir = electronBinPath

	Launch(os.Args[1:])
}
