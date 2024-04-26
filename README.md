
# Notes - static site generator


## Development

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

  To recompile the live server, run:
  ```
  go run compileServer.go
  ```

  Run the tests with:
  ```
  go test ./...
  ```
