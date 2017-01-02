# gobasic

This is a repository for learning Go programming language basics.

## Install Go
Xref [Installing Go](https://golang.org/doc/install)

### Example steps on Linux:
- Download Go from golang [downloads](https://golang.org/dl/)
```
$ wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
```
- Install Go
```
$ mkdir -p ~/Workspace/Go/
$ tar -C ~/Workspace/Go/ -xzf go1.7.4.linux-amd64.tar.gz
$ cd ~/Workspace/Go/
$ mkdir projects && cd projects
```
- Update ~/.bashrc to set envirnment variables
```
export GOROOT=$HOME/Workspace/Go/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/Workspace/Go/projects
```
