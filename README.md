# gogens #

go generate command that can be used to embed files into const strings.
Useful in cases where one wants to embed templates or scripts into the binary
instead of having to load them from files.


## Installation ##

You can download the code via the usual go utilities:

```
go get github.com/datacratic/gogens
```

## Usage ##

To embed template files into a package `vis` when the templates are located in
the current folder, and output them to `templates.go` one can use:
```
//go:generate include_suffix -package=vis -folder=. -suffix=tmpl -output=templates.go
```

A handy helper function gets created Get`suffix in uppercase` (in this case `GetTMPL()`) which returns a slice of all the const strings generated.

## License ##

The source code is available under the Apache License. See the LICENSE file for
more details.
