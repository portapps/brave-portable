package main

//go:generate go tool goversioninfo -icon=res/papp.ico -manifest=res/papp.manifest

import (
	_ "embed"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/portapps/portapps/v3"
	"github.com/portapps/portapps/v3/pkg/files"
	"github.com/portapps/portapps/v3/pkg/log"
	"github.com/portapps/portapps/v3/pkg/mutex"
	"github.com/portapps/portapps/v3/pkg/registry"
	"github.com/portapps/portapps/v3/pkg/shortcut"
	"github.com/portapps/portapps/v3/pkg/win"
	winregistry "golang.org/x/sys/windows/registry"
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
	if len(os.Args) > 1 && os.Args[1] == "register-browser" {
		if len(os.Args) != 2 {
			app.FatalBox(fmt.Errorf("Usage: %s register-browser", filepath.Base(os.Args[0])))
			return
		}
		if err := registerBrowser(); err != nil {
			app.FatalBox(err)
			return
		}
		if _, err := win.MsgBox(
			"Brave Portable",
			"Browser registered for the current user. Select Brave Portable for HTTP and HTTPS in Windows Settings > Apps > Default apps.",
			win.MsgBoxBtnOk|win.MsgBoxIconInformation); err != nil {
			log.Error().Err(err).Msg("Cannot display registration confirmation")
		}
		return
	}

	args, err := launchArgs(os.Args[1:])
	if err != nil {
		app.FatalBox(err)
		return
	}

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

	// Only the first wrapper owns registry imports, exports and cleanup.
	mu, err := mutex.Create(app.ID)
	if err != nil {
		if len(args) > 0 {
			log.Info().Msg("Forwarding arguments to the running instance")
			app.Launch(args)
		} else {
			app.ErrorBoxLog("Other instance detected. Brave Portable is already running.")
		}
		return
	}
	defer mutex.Release(mu)

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
	app.Launch(args)
}

// launchArgs keeps Windows URL associations separate from Brave's CLI. Reject
// extra arguments so an external URL cannot inject browser command-line flags.
func launchArgs(args []string) ([]string, error) {
	if len(args) == 0 || args[0] != "--open-url" {
		return args, nil
	}
	if len(args) != 2 {
		return nil, fmt.Errorf("--open-url requires exactly one HTTP or HTTPS URL")
	}
	u, err := url.Parse(args[1])
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return nil, fmt.Errorf("--open-url requires an absolute HTTP or HTTPS URL")
	}
	return []string{args[1]}, nil
}

// registerBrowser makes this wrapper available in Windows Default apps without
// changing the user's defaults or starting Brave.
func registerBrowser() error {
	wrapper, err := os.Executable()
	if err != nil {
		return err
	}
	const progID = `Software\Classes\Portapps.Brave.Url`
	const capabilities = `Software\Portapps\Brave\Capabilities`
	entries := []struct {
		path  string
		name  string
		value string
	}{
		{progID, "", "Brave Portable URL"},
		{progID, "URL Protocol", ""},
		{progID + `\DefaultIcon`, "", fmt.Sprintf(`"%s",0`, wrapper)},
		{progID + `\shell\open\command`, "", fmt.Sprintf(`"%s" --open-url "%%1"`, wrapper)},
		{capabilities, "ApplicationName", "Brave Portable"},
		{capabilities, "ApplicationDescription", "Brave via the Portapps portable wrapper"},
		{capabilities + `\URLAssociations`, "http", "Portapps.Brave.Url"},
		{capabilities + `\URLAssociations`, "https", "Portapps.Brave.Url"},
		{`Software\RegisteredApplications`, "Brave Portable", capabilities},
	}
	for _, entry := range entries {
		key, _, err := winregistry.CreateKey(winregistry.CURRENT_USER, entry.path, winregistry.SET_VALUE)
		if err != nil {
			return fmt.Errorf("Cannot create registry key %s: %w", entry.path, err)
		}
		err = key.SetStringValue(entry.name, entry.value)
		key.Close()
		if err != nil {
			return fmt.Errorf("Cannot set registry value %s\\%s: %w", entry.path, entry.name, err)
		}
	}
	return nil
}
