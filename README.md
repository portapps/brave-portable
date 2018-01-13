<p align="center"><a href="https://github.com/portapps/brave-portable" target="_blank"><img width="100" src="https://github.com/portapps/brave-portable/blob/master/res/papp.png"></a></p>

<p align="center">
  <a href="https://github.com/portapps/brave-portable/releases/latest"><img src="https://img.shields.io/github/release/portapps/brave-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/portapps/brave-portable/releases/latest"><img src="https://img.shields.io/github/downloads/portapps/brave-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/brave-portable"><img src="https://img.shields.io/appveyor/ci/crazy-max/brave-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/portapps/brave-portable"><img src="https://goreportcard.com/badge/github.com/portapps/brave-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/brave-portable"><img src="https://img.shields.io/codacy/grade/a416cd778ef743de91623aca7a622a8e.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://saythanks.io/to/crazymax"><img src="https://img.shields.io/badge/thank-crazymax-426aa5.svg?style=flat-square" alt="Say Thanks"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

[Brave](https://brave.com) portable app made with 🚀 [Portapps](https://github.com/portapps).<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are different kinds of artifacts :

* `brave-portable-{ia32,x64}-x.x.x-x-setup.exe` : Full portable release of Brave as a setup. **Recommended way**!
* `brave-portable-{ia32,x64}-x.x.x-x.7z` : Full portable release of Brave as a 7z archive.
* `brave-portable-{ia32,x64}.exe` : Only the portable binary (must be renamed `brave-portable.exe`)
* `BraveSetup-{ia32,x64}-x.x.x.exe` : The original setup from the [official website](https://brave.com/downloads.html).
* `brave-{ia32,x64}-x.x.x-full.nupkg` : The original NUPKG file extracted from the original setup.

### Fresh installation

Install `brave-portable-{ia32,x64}-x.x.x-x-setup.exe` where you want then run `brave-portable.exe`.

### App already installed

If you have already installed Brave from the original setup, do the same thing as a fresh installation and :

* Move data located in `%APPDATA%\brave\*` to `data` folder.

Run `brave-portable.exe` and then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) Brave from your computer.

### Upgrade

For an upgrade, simply download and install the [latest setup](https://github.com/portapps/brave-portable/releases/latest).

## How can i help ?

We welcome all kinds of contributions :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
Any funds donated will be used to help further development on this project! :gift_heart:

[![Donate Paypal](https://raw.githubusercontent.com/portapps/portapps/master/res/paypal.png)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG)

## License

MIT. See `LICENSE` for more details.<br />
Rocket icon credit to [Squid Ink](http://thesquid.ink).
