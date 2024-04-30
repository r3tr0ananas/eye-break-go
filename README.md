<div align="center">

  # 👀 eye break
  
  <sub>[Goldy's eye break program](https://github.com/thegoldenpro/eye_break) but in go<sub>
  
  <img width="320px" src="https://github.com/r3tr0ananas/eye-break-go/assets/132799819/51f436f9-3a99-4bfc-b649-a42511183b7d">
  <img width="320px" src="https://github.com/r3tr0ananas/eye-break-go/assets/132799819/ecd921f2-4d87-43e8-b723-0ce2742a746b">

</div>

## Installation 🛠️
Install from source

Prerequisites: **[``git``](https://git-scm.com/downloads), [``golang``](https://go.dev/), [``make``](https://www.gnu.org/software/make/) (recommended), [``systemd``](https://systemd.io/)**

### Linux 🐧
```sh
git clone https://github.com/r3tr0ananas/eye-break-go
cd eye-break-go
```
After cloning do:
```sh
make # build
make install # install to bin, install assets and add service to systemd
```

Enable the systemd service.
```sh
systemctl --user enable eye-break-go
```
Start it!
```sh
systemctl --user start eye-break-go
```

## Why did i do this?

i don't know, was bored.