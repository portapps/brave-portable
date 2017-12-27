# Changelog

## 0.19.122-18 (2017/12/27)

* New release of Brave : 0.19.122

## 0.19.116-17 (2017/12/14)

* New release of Brave : 0.19.116

## 0.19.105-16 (2017/12/09)

* New release of Brave : 0.19.105

## 0.19.95-15 (2017/11/21)

* Move app to a subfolder
* Reduce dependencies to avoid heuristic detection
* Add UPX compression

> :warning: **UPGRADE NOTES**
> * Delete all files and folders in root folder except `data` folder.

## 0.19.95-14 (2017/11/17)

* New release of Brave : 0.19.95
* Move repository to [Portapps](https://github.com/portapps) org

## 0.19.88-13 (2017/11/11)

* New release of Brave : 0.19.88

## 0.19.48-12 (2017/10/29)

* New release of Brave : 0.19.70
* Switch to [Golang Dep](https://github.com/golang/dep) as dependency manager
* Upgrade to Go 1.9.1

## 0.19.48-11 (2017/10/15)

* New release of Brave : 0.19.48

## 0.18.36-10 (2017/09/16)

* New release of Brave : 0.18.36
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
* New release of Brave : 0.18.23

## 0.18.14-7 (2017/08/09)

* New release of Brave : 0.18.14

## 0.17.19-6 (2017/07/23)

* New release of Brave : 0.17.19
* Admin privileges not required for Setup
* Small refactoring

## 0.17.13-5 (2017/07/09)

* New release of Brave : 0.17.13

## 0.16.9-4 (2017/05/31)

* New release of Brave : 0.16.9

## 0.15.310-3 (2017/05/31)

* New release of Brave : 0.15.310

## 0.15.2-2 (2017/05/14)

* New release of Brave : 0.15.2
* Provide the nupkg file in the release
* Add Go report card
* Move log file to the root folder
* Add Glide dependency manager for Go
* Standard code organization
* MIT license

## 0.15.1-1 (2017/05/03)

* Initial version
