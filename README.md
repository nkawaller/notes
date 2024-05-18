# Notes - static site generator

[![tests](https://github.com/nkawaller/notes/actions/workflows/test.yml/badge.svg)](https://github.com/nkawaller/notes/actions/workflows/test.yml)
[![deploy](https://github.com/nkawaller/notes/actions/workflows/deploy.yml/badge.svg)](https://github.com/nkawaller/notes/actions/workflows/deploy.yml)

Welcome to the Notes project——a dynamic tool for generating static sites
with ease and efficiency. Harnessing the power of Go, this project not
only facilitates static site generation but also provides a live server
for instant feedback during development.

## Development

### Live Server

Apart from its primary function as a static-site generator, Notes also
offers a live server feature to enhance your development experience. 
Running the server involves two steps——first start the tailwind 
listener:

```
./runtailwind.sh
```

and then start the server:

```
./bin/runserver
```

### Compile Binaries



### Build Static Site

Execute the following command to generate the static site, and transfer
the CSS file to the `deploy/static` directory. This command also creates
the `runserver` and `staticgen` binaries:

```
go run buildSite.go
```

### Test

To run all tests from the root directory, utilize the following command.

```
go test ./...
```
