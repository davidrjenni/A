[![Build Status](https://drone.io/github.com/davidrjenni/A/status.png)](https://drone.io/github.com/davidrjenni/A/latest)
[![GoDoc](https://godoc.org/github.com/davidrjenni/A?status.svg)](https://godoc.org/github.com/davidrjenni/A)

# A

A - Go tools for Acme.

## Installation

```
% go get github.com/davidrjenni/A

% go get -u github.com/zmb3/gogetdoc
% go get github.com/godoctor/godoctor
% go get github.com/josharian/impl
% go get golang.org/x/tools/cmd/gorename
```

## Usage

### Documentation

```
A doc
```
Shows the documentation for the entity under the cursor.
This command uses `github.com/zmb3/gogetdoc`.

### Extract to Function/Method

```
A ex <name>
```
Extracts the selected statements to a new function/method with name `<name>`.
This command uses `github.com/godoctor/godoctor`.

### Generate Method Stubs

```
A impl <recv> <iface>
A impl 'f *File' io.ReadWriteCloser
```
Generates method stubs with receiver `<recv>` for implementing the interface `<iface>` and inserts them at the location of the cursor.
This command uses `github.com/josharian/impl`.

### Renaming

```
A rn <name>
```
Renames the entity under the cursor with `<name>`.
This commands uses `golang.org/x/tools/cmd/gorename`.

### Share

```
A share
```
Uploads the selected snippet to play.golang.org and prints the URL.
