//go:generate go get -v github.com/josephspurrier/goversioninfo/...
//go:generate goversioninfo -icon=res/app-portable.ico
package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/op/go-logging"
)

const (
	NAME            = "brave-portable"
	APP_NAME        = "Brave"
	APP_DATA_FOLDER = "brave"
	APP_PROCESS     = "Brave.exe"
)

var (
	log       = logging.MustGetLogger(NAME)
	logFormat = logging.MustStringFormatter(`%{time:2006-01-02 15:04:05} %{level:.4s} - %{message}`)
)

func main() {
	// Current path
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("Current path:", err)
	}

	// Log file
	logfile, err := os.OpenFile(pathJoin(currentPath, NAME+".log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Error("Log file:", err)
	}

	// Init logger
	logBackendStdout := logging.NewBackendFormatter(logging.NewLogBackend(os.Stdout, "", 0), logFormat)
	logBackendFile := logging.NewBackendFormatter(logging.NewLogBackend(logfile, "", 0), logFormat)
	logging.SetBackend(logBackendStdout, logBackendFile)
	log.Info("--------")
	log.Info("Starting " + NAME + "...")
	log.Info("Current path:", currentPath)

	// Find app folder
	log.Info("Lookup app folder in", currentPath)
	var appPath = ""
	rootFiles, _ := ioutil.ReadDir(currentPath)
	for _, f := range rootFiles {
		if strings.HasPrefix(f.Name(), "app-") && f.IsDir() {
			log.Info("App folder found:", f.Name())
			appPath = pathJoin(currentPath, f.Name())
			break
		}
	}
	if _, err := os.Stat(appPath); err == nil {
		log.Info("App path:", appPath)
	} else {
		log.Error("App path does not exist")
	}

	// Init vars
	appExe := pathJoin(appPath, APP_PROCESS)
	dataPath := pathJoin(currentPath, "data")
	dataAppPath := pathJoin(dataPath, "AppData", "Roaming", APP_DATA_FOLDER)
	log.Info("App executable:", appExe)
	log.Info("Data path:", dataPath)

	// Create data folder
	log.Info("Create data folder...", dataAppPath)
	err = os.MkdirAll(dataAppPath, 777)
	if err != nil {
		log.Error("Create data folder:", err)
	}

	// Override USERPROFILE env var
	if err := os.Setenv("USERPROFILE", dataPath); err != nil {
		log.Error(err)
	}

	// Launch
	log.Infof("Launch %s...", APP_NAME)
	execApp := exec.Command(appExe)
	execApp.Dir = appPath

	defer logfile.Close()
	execApp.Stdout = logfile
	execApp.Stderr = logfile

	if err := execApp.Start(); err != nil {
		log.Error("Cmd Start:", err)
	}

	execApp.Wait()
}

func pathJoin(elem ...string) string {
	for i, e := range elem {
		if e != "" {
			return strings.Join(elem[i:], `\`)
		}
	}
	return ""
}
