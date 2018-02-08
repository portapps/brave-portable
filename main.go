//go:generate go get -v github.com/josephspurrier/goversioninfo/...
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
	Papp.DataPath = AppPathJoin("data")

	electronBinPath := PathJoin(Papp.AppPath, FindElectronAppFolder("app-", Papp.AppPath))

	Papp.Process = PathJoin(Papp.AppPath, "Brave.exe")
	Papp.Args = nil
	Papp.WorkingDir = electronBinPath

	Launch(os.Args[1:])
}
