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

### Generate New Static Site

Run the following command to generate the static site, and copy the css file over to the `deploy/static` directory.

```
go run buildSite.go
```

### Test

From the root directory run all tests with:

```
go test ./...
```
