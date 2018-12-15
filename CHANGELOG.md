# Changelog

## 0.25.302-43 (2018/12/15)

* Upgrade to Brave 0.25.302

## 0.25.2-42 (2018/11/11)

* Chrome user-data-dir not saved (Issue #8)
* Upgrade to Go 1.11.2

## 0.25.2-40 (2018/10/12)

* Upgrade to Brave 0.25.2

## 0.24.0-39 (2018/09/27)

* Upgrade to Brave 0.24.0
* Upgrade to Go 1.11
* Use [go mod](https://golang.org/cmd/go/#hdr-Module_maintenance) instead of `dep`

## 0.23.105-38 (2018/09/01)

* Upgrade to Brave 0.23.105

## 0.23.79-37 (2018/08/15)

* Upgrade to Brave 0.23.79

## 0.23.77-36 (2018/08/07)

* Upgrade to Brave 0.23.77

## 0.23.73-35 (2018/08/03)

* Upgrade to Brave 0.23.73 

## 0.23.39-34 (2018/07/19)

* Upgrade to Brave 0.23.39 

## 0.23.31-33 (2018/07/12)

* Upgrade to Brave 0.23.31

## 0.23.19-32 (2018/07/03)

* Upgrade to Brave 0.23.19

## 0.22.810-31 (2018/06/21)

* Upgrade to Brave 0.22.810

## 0.22.727-30 (2018/06/05)

* Upgrade to Brave 0.22.727

## 0.22.721-29 (2018/06/02)

* Upgrade to Brave 0.22.721

## 0.22.714-28 (2018/05/23)

* Upgrade to Brave 0.22.714

## 0.22.667-27 (2018/04/26)

* Upgrade to Brave 0.22.667

## 0.22.22-26 (2018/04/20)

* Upgrade to Brave 0.22.22

## 0.22.13-25 (2018/04/05)

* Upgrade to Brave 0.22.13

## 0.21.18-24 (2018/03/02)

* Upgrade to Brave 0.21.18

## 0.20.30-23 (2018/02/11)

* Move ia32/x64 to win32/win64 for arch def
* Remove nupkg file

## 0.20.30-22 (2018/02/08)

* Upgrade to Brave 0.20.30
* Ability to pass custom args to the portable process

## 0.19.139-21 (2018/01/21)

* Upgrade to Brave 0.19.139

## 0.19.134-20 (2018/01/13)

* Upgrade to Brave 0.19.134
* Rebuild electron.asar to push data directly in `data` subfolder
* No need to override USERPROFILE environment variable anymore

> :warning: **UPGRADE NOTES**
> * Move everything in `data\AppData\Roaming\brave\*` to `data` folder and remove folder `data\AppData`.

## 0.19.131-19 (2018/01/12)

* Upgrade to Brave 0.19.131

## 0.19.122-18 (2017/12/27)

* Upgrade to Brave 0.19.122

## 0.19.116-17 (2017/12/14)

* Upgrade to Brave 0.19.116

## 0.19.105-16 (2017/12/09)

* Upgrade to Brave 0.19.105

## 0.19.95-15 (2017/11/21)

* Move app to a subfolder
* Reduce dependencies to avoid heuristic detection
* Add UPX compression

> :warning: **UPGRADE NOTES**
> * Delete all files and folders in root folder except `data` folder.

## 0.19.95-14 (2017/11/17)

* Upgrade to Brave 0.19.95
* Move repository to [Portapps](https://github.com/portapps) org

## 0.19.88-13 (2017/11/11)

* Upgrade to Brave 0.19.88

## 0.19.48-12 (2017/10/29)

* Upgrade to Brave 0.19.70
* Switch to [Golang Dep](https://github.com/golang/dep) as dependency manager
* Upgrade to Go 1.9.1

## 0.19.48-11 (2017/10/15)

* Upgrade to Brave 0.19.48

## 0.18.36-10 (2017/09/16)

* Upgrade to Brave 0.18.36
* Change logger

## 0.18.23-9 (2017/09/05)

* Override USERPROFILE env var instead of using symlink to APPDATA to store data
* Do not migrate old data folder from APPDATA
* Reduce dependencies and system calls to avoid heuristic detection

> :warning: **UPGRADE NOTES**
> * Move the content of `data\*` in `data\AppData\Roaming\brave\`
> * Remove symlink `%APPDATA%\brave`

## 0.18.23-8 (2017/08/26)

* Heuristic detection (Issue #2)
* Upgrade to Go 1.9
* Add support guidelines
* Upgrade to Brave 0.18.23

## 0.18.14-7 (2017/08/09)

* Upgrade to Brave 0.18.14

## 0.17.19-6 (2017/07/23)

* Upgrade to Brave 0.17.19
* Admin privileges not required for Setup
* Small refactoring

## 0.17.13-5 (2017/07/09)

* Upgrade to Brave 0.17.13

## 0.16.9-4 (2017/05/31)

* Upgrade to Brave 0.16.9

## 0.15.310-3 (2017/05/31)

* Upgrade to Brave 0.15.310

## 0.15.2-2 (2017/05/14)

* Upgrade to Brave 0.15.2
* Provide the nupkg file in the release
* Add Go report card
* Move log file to the root folder
* Add Glide dependency manager for Go
* Standard code organization
* MIT license

## 0.15.1-1 (2017/05/03)

* Initial version
