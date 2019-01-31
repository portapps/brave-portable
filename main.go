//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Brave.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico
package main

import (
	"io/ioutil"
	"os"
	"path"

	_ "github.com/kevinburke/go-bindata"
	"github.com/portapps/brave-portable/assets"
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
	Papp.Process = PathJoin(Papp.AppPath, "brave.exe")
	Papp.Args = []string{
		"--user-data-dir=" + Papp.DataPath,
		"--disable-brave-update",
		"--no-default-browser-check",
		"--allow-outdated-plugins",
		"--disable-logging",
		"--disable-breakpad",
		"--disable-machine-id",
		"--disable-encryption-win",
	}
	Papp.WorkingDir = Papp.AppPath

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Brave Portable.lnk")
	defaultShortcut, err := assets.Asset("Brave.lnk")
	if err != nil {
		Log.Error("Cannot load asset Brave.lnk:", err)
	}
	err = ioutil.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		Log.Error("Cannot write default shortcut:", err)
	}

	// Update default shortcut
	err = CreateShortcut(WindowsShortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       Papp.Process,
		Arguments:        WindowsShortcutProperty{Clear: true},
		Description:      WindowsShortcutProperty{Value: "Brave Portable by Portapps"},
		IconLocation:     WindowsShortcutProperty{Value: Papp.Process},
		WorkingDirectory: WindowsShortcutProperty{Value: Papp.AppPath},
	})
	if err != nil {
		Log.Error("Cannot create shortcut:", err)
	}

	Launch(os.Args[1:])

	_ = os.Remove(shortcutPath)
}
