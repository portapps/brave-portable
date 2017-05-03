<p align="center"><a href="https://github.com/crazy-max/brave-portable" target="_blank"><img width="100" src="https://github.com/crazy-max/brave-portable/blob/master/res/brave-portable.png"></a></p>

<p align="center">
  <a href="https://github.com/crazy-max/brave-portable/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/brave-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/brave-portable/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/brave-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/brave-portable"><img src="https://img.shields.io/appveyor/ci/crazy-max/brave-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://www.codacy.com/app/crazy-max/brave-portable"><img src="https://img.shields.io/codacy/grade/a416cd778ef743de91623aca7a622a8e.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=SCMXA34TCLGMG"><img src="https://img.shields.io/badge/donate-paypal-blue.svg?style=flat-square" alt="Donate Paypal"></a>
  <a href="https://flattr.com/submit/auto?user_id=crazymax&url=https://github.com/crazy-max/brave-portable"><img src="https://img.shields.io/badge/flattr-this-green.svg?style=flat-square" alt="Flattr this!"></a>
</p>

## About

A single EXE created with [Golang](https://golang.org/) to make [Brave](https://brave.com) portable on Windows systems.<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are four different kinds of artifacts :

* `brave-portable-{ia32,x64}-x.x.x-x-setup.exe` : Full portable release of Brave as a setup. **Recommended way**!
* `brave-portable-{ia32,x64}-x.x.x-x.{7z,zip}` : Full portable release of Brave as an archive.
* `brave-portable-{ia32,x64}.exe` : Only the portable binary (must be renamed `brave-portable.exe`)
* `BraveSetup-{ia32,x64}-x.x.x.exe` : The original setup from the [official website](https://brave.com/downloads.html).

For a **fresh installation**, install `brave-portable-{ia32,x64}-x.x.x-x-setup.exe` where you want then run `brave-portable.exe`.

If **you have already installed Brave from the original setup**, do the same thing as a fresh installation and run `brave-portable.exe` a first time.<br />
The data located in `%APPDATA%\brave` will be moved in the `data` folder.<br />
Then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) Brave from your computer.

**For an upgrade**, simply download and install the [latest setup](https://github.com/crazy-max/brave-portable/releases/latest).

## How can i help ?

We welcome all kinds of contributions :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
Any funds donated will be used to help further development on this project! :gift_heart:

<p>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=SCMXA34TCLGMG">
    <img src="https://github.com/crazy-max/brave-portable/blob/master/res/paypal.png" alt="Donate Paypal">
  </a>
  <a href="https://flattr.com/submit/auto?user_id=crazymax&url=https://github.com/crazy-max/brave-portable">
    <img src="https://github.com/crazy-max/brave-portable/blob/master/res/flattr.png" alt="Flattr this!">
  </a>
</p>

## License

LGPL. See `LICENSE` for more details.<br />
USB icon credit to [Juliia Osadcha](https://www.iconfinder.com/Juliia_Os).
