# Notes - static site generator [![tests](https://github.com/nkawaller/notes/actions/workflows/01-test.yml/badge.svg)](https://github.com/nkawaller/notes/actions/workflows/01-test.yml) [![build](https://github.com/nkawaller/notes/actions/workflows/02-build-site.yml/badge.svg?branch=main)](https://github.com/nkawaller/notes/actions/workflows/02-build-site.yml)

## Development

### Live Server

In addition to being a static-site generator, there's also a live
server you can run to get instant feedback while you develop. To set
this up, start the tailwind listener:

```
./runtailwind.sh
```

and then start the server:

```
./bin/runserver
```

### Compile Binaries and Build Static Site

Run the following command to generate the static site, and copy the css 
file over to the `deploy/static` directory. The runserver and staticgen
binaries are also created when running this command.

```
go run buildSite.go
```

### Test

From the root directory run all tests with:

```
go test ./...
```
