package main

//go:generate go tool goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/files"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/registry"
	"github.com/portapps/portapps/v3/pkg/shortcut"
)

//go:embed res/Brave.lnk
var defaultShortcut []byte

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
	if err := os.MkdirAll(app.DataPath, 0o755); err != nil {
		log.Fatal().Err(err).Msg("Cannot create data path")
	}
	app.Process = filepath.Join(app.AppPath, "brave.exe")
	app.Args = []string{
		"--user-data-dir=" + app.DataPath,
		"--disable-brave-update",
		"--no-default-browser-check",
		"--disable-logging",
		"--disable-breakpad",
		"--disable-machine-id",
		"--disable-encryption-win",
		"--update-feed-url=" + fmt.Sprintf("https://raw.githubusercontent.com/portapps/brave-portable/refs/tags/%s-%s/res/appcast.xml", app.Info.Version, app.Info.Release),
	}

	// Cleanup on exit
	if cfg.Cleanup {
		defer func() {
			files.Cleanup(
				filepath.Join(os.Getenv("APPDATA"), "BraveSoftware"),
				filepath.Join(os.Getenv("LOCALAPPDATA"), "BraveSoftware"),
			)
		}()
	}

	// Copy default shortcut
	shortcutPath := filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "Brave Portable.lnk")
	err := os.WriteFile(shortcutPath, defaultShortcut, 0644)
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
	regsPath := filepath.Join(app.RootPath, "reg")
	if err := os.MkdirAll(regsPath, 0o755); err != nil {
		log.Fatal().Err(err).Msg("Cannot create registry path")
	}
	bsRegKey := registry.Key{
		Key:  `HKCU\SOFTWARE\BraveSoftware`,
		Arch: "32",
	}
	bbdRegKey := registry.Key{
		Key:  `HKCU\SOFTWARE\Brave-Browser-Development`,
		Arch: "32",
	}

	if err := bsRegKey.Import(filepath.Join(regsPath, "BraveSoftware.reg")); err != nil {
		log.Error().Err(err).Msg("Cannot import registry key")
	}
	if err := bbdRegKey.Import(filepath.Join(regsPath, "Brave-Browser-Development.reg")); err != nil {
		log.Error().Err(err).Msg("Cannot import registry key")
	}

	defer func() {
		if err := bsRegKey.Export(filepath.Join(regsPath, "BraveSoftware.reg")); err != nil {
			log.Error().Err(err).Msg("Cannot export registry key")
		}
		if err := bbdRegKey.Export(filepath.Join(regsPath, "Brave-Browser-Development.reg")); err != nil {
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
