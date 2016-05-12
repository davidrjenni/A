[![Build Status](https://drone.io/github.com/davidrjenni/A/status.png)](https://drone.io/github.com/davidrjenni/A/latest)
[![GoDoc](https://godoc.org/github.com/davidrjenni/A?status.svg)](https://godoc.org/github.com/davidrjenni/A)

# A

A - Go tools for Acme.

## Installation

```
% go get github.com/davidrjenni/A

% go get golang.org/x/tools/cmd/guru
% go get github.com/zmb3/gogetdoc
% go get github.com/godoctor/godoctor
% go get github.com/josharian/impl
% go get golang.org/x/tools/cmd/gorename
```

## Usage

### Goto Definition

```
A def
```
Shows the declaration for the identifier under the cursor.
This command uses `golang.org/x/tools/cmd/guru`.

### Describe

```
A desc
```
Describes the declaration for the syntax under the cursor.
This command uses `golang.org/x/tools/cmd/guru`.

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

### Freevars

```
A fv
```
Shows the free variables of the selected snippet.
This command uses `golang.org/x/tools/cmd/guru`.

### Generate Method Stubs

```
A impl <recv> <iface>
A impl 'f *File' io.ReadWriteCloser
```
Generates method stubs with receiver `<recv>` for implementing the interface `<iface>` and inserts them at the location of the cursor.
This command uses `github.com/josharian/impl`.

### Referrers

```
A refs
```
Shows all refs to the entity denoted by identifier under the cursor.
This command uses `golang.org/x/tools/cmd/guru`.

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
