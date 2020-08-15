<h1 align=center>
<img src="logo/512.svg" width=25%>
</h1>

# A ![CI](https://github.com/davidrjenni/A/workflows/CI/badge.svg?branch=master) [![GoDoc](https://godoc.org/github.com/davidrjenni/A?status.svg)](https://godoc.org/github.com/davidrjenni/A) [![Go Report Card](https://goreportcard.com/badge/github.com/davidrjenni/A)](https://goreportcard.com/report/github.com/davidrjenni/A)

A - Go tools for Acme.

## Installation

```
% go get github.com/davidrjenni/A

% go get -u golang.org/x/tools/cmd/guru
% go get -u github.com/zmb3/gogetdoc
% go get -u github.com/godoctor/godoctor
% go get -u github.com/josharian/impl
% go get -u golang.org/x/tools/cmd/gorename
% go get -u github.com/fatih/gomodifytags
% go get -u github.com/davidrjenni/reftools/cmd/fillstruct
% go get -u github.com/davidrjenni/reftools/cmd/fillswitch
```

## Usage

### Add Struct Tags

```
A addtags <tags> [options]
```
Adds struct tags and tag options to the selected struct fields.
`<tags>` is a comma-separated list of tags to add, e.g. `json,xml`.
`[options]` is an optional list of tag options to add, e.g. `'json=omitempty,xml=omitempty'`
This command uses `github.com/fatih/gomodifytags`.
See it in action [here](https://twitter.com/davidrjenni/status/893130797376516096).

### Callees

```
A cle <scope>
```
Shows possible targets of the function call under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Callers

```
A clr <scope>
```
Shows possible callers of the function under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Callstack

```
A cs <scope>
```
Shows the path from the callgraph root to the function under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

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

### Errors

```
A err <scope>
```
Shows possible values of the error variable under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Extract to Function/Method

```
A ex <name>
```
Extracts the selected statements to a new function/method with name `<name>`.
This command uses `github.com/godoctor/godoctor`.

### Fill a Struct Literal with Default Values

```
A fstruct
```
Fills the selected struct literal with default values.
This command uses `github.com/davidrjenni/reftools/cmd/fillstruct`.

### Fill a (Type) Switch Statement with Case Statements

```
A fswitch
```
Fills the selected (type) switch statement with case statements.
This command uses `github.com/davidrjenni/reftools/cmd/fillswitch`.

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

### Implements

```
A impls <scope>
```
Shows the `implements` relation for the type or method under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Peers

```
A peers <scope>
```
Shows send/receive corresponding to the selected channel op.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Points To

```
A pto <scope>
```
Shows variables the selected pointer may point to.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Referrers

```
A refs
```
Shows all refs to the entity denoted by identifier under the cursor.
This command uses `golang.org/x/tools/cmd/guru`.

### Remove Struct Tags

```
A rmtags <tags> [options]
```
Removes struct tags and tag options from the selected struct fields.
`<tags>` is a comma-separated list of tags to remove, e.g. `json,xml`.
`[options]` is an optional list of tag options to remove, e.g. `'json=omitempty,xml=omitempty'`
This command uses `github.com/fatih/gomodifytags`.
See it in action [here](https://twitter.com/davidrjenni/status/893130797376516096).

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

### What

```
A what
```
Shows basic information about the selected syntax node.
This command uses `golang.org/x/tools/cmd/guru`.
