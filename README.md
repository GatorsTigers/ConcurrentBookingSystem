# ConcurrentBookingSystem
Concurrent ticket booking system.


# Project setup
### Go requires you to organize your code within a specific workspace directory. The workspace contains three subdirectories: bin, pkg, and src. The source code for your projects should be placed in the src directory.

### Create a new directory for your workspace (choose a suitable name) and set the GOPATH environment variable to point to that directory.
> export GOPATH="/Users/anmol/GoProjects/"

### Create a new project directory:
> mkdir -pv $GOPATH/src/github.com/GatorsTigers/ConcurrentBookingSystem
> cd $GOPATH/src/github.com/GatorsTigers/ConcurrentBookingSystem

### Initialize a Go module:
### Go modules are a dependency management system introduced in Go 1.11 to help manage external dependencies. Initialize a Go module in your project directory.

> go mod init github.com/GatorsTigers/ConcurrentBookingSystem

### Then clone your repository like:
> cd $GOPATH/src/github.com/GatorsTigers/
> git clone github.com/GatorsTigers/ConcurrentBookingSystem

### Install dependencies like so:
> go get -u github.com/gin-gonic/gin
> go get -u gorm.io/gorm
> go get -u gorm.io/driver/mysql

