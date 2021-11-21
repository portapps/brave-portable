//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -prefix res/ -pkg assets -o assets/assets.go res/Brave.lnk
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest
package main

import (
	"os"
	"path"

	"github.com/portapps/brave-portable/assets"
	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/registry"
	"github.com/portapps/portapps/v3/pkg/shortcut"
	"github.com/portapps/portapps/v3/pkg/utl"
)

type config struct {
	Cleanup bool `yaml:"cleanup" mapstructure:"cleanup"`
}

var (
	app *portapps.App
	cfg *config
)

func init() {
	var err error

	// Default config
	cfg = &config{
		Cleanup: false,
	}

	// Init app
	if app, err = portapps.NewWithCfg("brave-portable", "Brave", cfg); err != nil {
		log.Fatal().Err(err).Msg("Cannot initialize application. See log file for more info.")
	}
}

func main() {
	utl.CreateFolder(app.DataPath)
	app.Process = utl.PathJoin(app.AppPath, "brave.exe")
	app.Args = []string{
		"--user-data-dir=" + app.DataPath,
		"--disable-brave-update",
		"--no-default-browser-check",
		"--disable-logging",
		"--disable-breakpad",
		"--disable-machine-id",
		"--disable-encryption-win",
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			utl.Cleanup([]string{
				path.Join(os.Getenv("APPDATA"), "BraveSoftware"),
				path.Join(os.Getenv("LOCALAPPDATA"), "BraveSoftware"),
			})
		}()
	}

	// Copy default shortcut
	shortcutPath := path.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Brave Portable.lnk")
	defaultShortcut, err := assets.Asset("Brave.lnk")
	if err != nil {
		log.Error().Err(err).Msg("Cannot load asset Brave.lnk")
	}
	err = os.WriteFile(shortcutPath, defaultShortcut, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Cannot write default shortcut")
	}

	// Update default shortcut
	err = shortcut.Create(shortcut.Shortcut{
		ShortcutPath:     shortcutPath,
		TargetPath:       app.Process,
		Arguments:        shortcut.Property{Clear: true},
		Description:      shortcut.Property{Value: "Brave Portable by Portapps"},
		IconLocation:     shortcut.Property{Value: app.Process},
		WorkingDirectory: shortcut.Property{Value: app.AppPath},
	})
	if err != nil {
		log.Error().Err(err).Msg("Cannot create shortcut")
	}
	defer func() {
		if err := os.Remove(shortcutPath); err != nil {
			log.Error().Err(err).Msg("Cannot remove shortcut")
		}
	}()

	// Registry keys
	regsPath := utl.CreateFolder(app.RootPath, "reg")
	bsRegKey := registry.Key{
		Key:  `HKCU\SOFTWARE\BraveSoftware`,
		Arch: "32",
	}
	bbdRegKey := registry.Key{
		Key:  `HKCU\SOFTWARE\Brave-Browser-Development`,
		Arch: "32",
	}

	if err := bsRegKey.Import(utl.PathJoin(regsPath, "BraveSoftware.reg")); err != nil {
		log.Error().Err(err).Msg("Cannot import registry key")
	}
	if err := bbdRegKey.Import(utl.PathJoin(regsPath, "Brave-Browser-Development.reg")); err != nil {
		log.Error().Err(err).Msg("Cannot import registry key")
	}

	defer func() {
		if err := bsRegKey.Export(utl.PathJoin(regsPath, "BraveSoftware.reg")); err != nil {
			log.Error().Err(err).Msg("Cannot export registry key")
		}
		if err := bbdRegKey.Export(utl.PathJoin(regsPath, "Brave-Browser-Development.reg")); err != nil {
			log.Error().Err(err).Msg("Cannot export registry key")
		}
		if cfg.Cleanup {
			if err := bsRegKey.Delete(true); err != nil {
				log.Error().Err(err).Msg("Cannot remove registry key")
			}
			if err := bbdRegKey.Delete(true); err != nil {
				log.Error().Err(err).Msg("Cannot remove registry key")
			}
		}
	}()

	defer app.Close()
	app.Launch(os.Args[1:])
}
