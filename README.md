# lablog
### The simple little homelab notebook
![screenshot](lablog-screenshot.png)
### Motivation
 - Find yourself forgetting usefull information about your home network devices? Need another CRUD app in your life? Store info here for easy access fom inside your terminal. See device online status at a glance.
### Quick Start
 - Required: sqlite3
 - Clone this repository and use the go toolchain to install or build an executable. `go install .` or `go build .` in the root of the repo.
### Usage
 - Note: Help section/key is in progress. See below for keybinds:
  - "+" Register a device
  - "e/u" Update device details
  - "del" Remove device
  - "enter" Detail view
  - "esc" Back out to main view/update online status
  - "q" Quit
### Contributing
 - If you would like to contribute, please fork and open a pull request to the main branch.
### TODO
 - Refine formatting (especially the forms)
 - Style the detail view
 - Add port register functionality to keep track of useful ports/APIs
