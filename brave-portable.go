//go:generate goversioninfo -icon=brave-portable.ico
package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "os/exec"
  "path"
  "path/filepath"
  "strings"
  "syscall"

  "github.com/op/go-logging"
)

var log = logging.MustGetLogger("brave-portable")
var logFormat = logging.MustStringFormatter(`%{time:2006-01-02 15:04:05} %{level:.4s} - %{message}`)

func main() {
  // Current path
  currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Error("Current path:", err)
  }

  // Logs folder
  var logsPath = path.Join(currentPath, "logs")
  if _, err := os.Stat(logsPath); os.IsNotExist(err) {
    log.Info("Create logs folder", logsPath)
    err = os.Mkdir(logsPath, 777)
    if err != nil {
      log.Error("Create logs folder:", err)
    }
  }

  // Log file
  logfile, err := os.OpenFile(path.Join(logsPath, "brave-portable.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
    log.Error("Log file:", err)
  }

  // Init logger
  logBackendStdout := logging.NewBackendFormatter(logging.NewLogBackend(os.Stdout, "", 0), logFormat)
  logBackendFile := logging.NewBackendFormatter(logging.NewLogBackend(logfile, "", 0), logFormat)
  logging.SetBackend(logBackendStdout, logBackendFile)
  log.Info("--------")
  log.Info("Starting brave-portable...")

  // Purge logs
  logsFolder, err := os.Open(logsPath)
  if err != nil {
    log.Error("Open logs folder:", err)
  }
  defer logsFolder.Close()
  logsFiles, err := logsFolder.Readdir(-1)
  if err != nil {
    log.Error("Read logs folder:", err)
  }
  log.Info("Reading", logsPath)
  for _, logsFile := range logsFiles {
    if ! strings.HasPrefix(logsFile.Name(), "brave-portable") {
      os.Remove(path.Join(logsPath, logsFile.Name()))
      log.Info("Deleted", path.Join(logsPath, logsFile.Name()))
    }
  }

  // Convert backslashes
  currentPath = path.Join(strings.Replace(string(currentPath), string(filepath.Separator), "/", -1))
  log.Info("Current path:" + currentPath)

  // Find app folder
  log.Info("Lookup app folder in", currentPath)
  var appPath = ""
  rootFiles, _ := ioutil.ReadDir(currentPath)
  for _, f := range rootFiles {
    if strings.HasPrefix(f.Name(), "app-") && f.IsDir() {
      log.Info("App folder found:", f.Name())
      appPath = path.Join(currentPath, f.Name())
      break
    }
  }
  if _, err := os.Stat(appPath); err == nil {
    log.Info("App path:", appPath)
  } else {
    log.Error("App path does not exist");
  }

  // Init vars
  var braveExe = path.Join(appPath, "Brave.exe")
  var dataPath = path.Join(currentPath, "data")
  var symlinkPath = path.Clean(path.Join(os.Getenv("APPDATA"), "brave"))
  log.Info("Brave executable:", braveExe)
  log.Info("Data path:", dataPath)
  log.Info("Symlink path:", symlinkPath)

  // Create data folder
  if _, err := os.Stat(dataPath); os.IsNotExist(err) {
    log.Info("Create data folder...", dataPath)
    err = os.Mkdir(dataPath, 777)
    if err != nil {
      log.Error("Create data folder:", err)
    }
  }

  // Check old data folder
  if _, err := os.Stat(symlinkPath); err == nil {
    fi, err := os.Lstat(symlinkPath)
    if err != nil {
      log.Error("Symlink lstat:", err)
    }
    if fi.Mode() & os.ModeSymlink != os.ModeSymlink {
      // Copy old data folder
      log.Info("Copy old data from", symlinkPath)
      err = copyDir(symlinkPath, dataPath)
      if err != nil {
        log.Error("Copying old data folder:", err)
      }

      // Rename old data folder
      log.Info("Chmod old data folder...")
      err = os.Chmod(symlinkPath, 0777)
      if err != nil {
        log.Error("Chmod old data folder:", err)
      }

      log.Info("Rename old data folder to", symlinkPath + "_old")
      err = os.Rename(symlinkPath, symlinkPath + "_old")
      if err != nil {
        log.Error("Renaming old data folder:", err)
      }
    }
  }

  // Create symlink
  log.Info("Create symlink", symlinkPath)
  os.Remove(symlinkPath)
  cmd := exec.Command("cmd", "/c", "mklink", "/J", strings.Replace(symlinkPath, "/", "\\", -1), strings.Replace(dataPath, "/", "\\", -1))
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
  if err := cmd.Run(); err != nil {
    log.Error("Symlink:", err)
  }
  /*err = os.Symlink(dataPath, symlinkPath)
  if err != nil {
    log.Error(err)
  }*/

  // Launch brave
  log.Info("Launch Brave...")
  cmd = exec.Command(braveExe, "--log-file", "./")
  cmd.Dir = logsPath

  defer logfile.Close()
  cmd.Stdout = logfile
  cmd.Stderr  = logfile

  if err := cmd.Start(); err != nil {
    log.Error("Cmd Start:", err)
  }

  cmd.Wait()
}

// src: https://gist.github.com/m4ng0squ4sh/92462b38df26839a3ca324697c8cba04
func copyFile(src, dst string) (err error) {
  in, err := os.Open(src)
  if err != nil {
    return
  }
  defer in.Close()

  out, err := os.Create(dst)
  if err != nil {
    return
  }
  defer func() {
    if e := out.Close(); e != nil {
      err = e
    }
  }()

  _, err = io.Copy(out, in)
  if err != nil {
    return
  }

  err = out.Sync()
  if err != nil {
    return
  }

  si, err := os.Stat(src)
  if err != nil {
    return
  }

  err = os.Chmod(dst, si.Mode())
  if err != nil {
    return
  }

  return
}

// src: https://gist.github.com/m4ng0squ4sh/92462b38df26839a3ca324697c8cba04
func copyDir(src string, dst string) (err error) {
  src = filepath.Clean(src)
  dst = filepath.Clean(dst)

  si, err := os.Stat(src)
  if err != nil {
    return err
  }
  if !si.IsDir() {
    return fmt.Errorf("source is not a directory")
  }

  _, err = os.Stat(dst)
  if err != nil && !os.IsNotExist(err) {
    return
  }

  err = os.MkdirAll(dst, si.Mode())
  if err != nil {
    return
  }

  entries, err := ioutil.ReadDir(src)
  if err != nil {
    return
  }

  for _, entry := range entries {
    srcPath := filepath.Join(src, entry.Name())
    dstPath := filepath.Join(dst, entry.Name())
    if entry.IsDir() {
      err = copyDir(srcPath, dstPath)
      if err != nil {
        return
      }
    } else {
      err = copyFile(srcPath, dstPath)
      if err != nil {
        return
      }
    }
  }

  return
}
