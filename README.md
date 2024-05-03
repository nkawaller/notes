# Notes - static site generator

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

### Compile App

To compile the live server and static site generator run:

```
go run compileBinaries.go
```
For right now, this command actually compiles both binaries, generates
the static site, and copies the css file over to the `deploy/static`
directory.

### Generate New Static Site
```
./bin/staticgen
```

### Test

Run the tests with:

```
go test ./...
```
