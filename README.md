<h1 align="center">
Kalvium Test
</h1>

## Project Structure

```
├── configuration
│   └── config.go
├── controllers
│   ├── expressionController.go
│   └── historyController.go
├── Dockerfile
├── flake.lock
├── flake.nix
├── go.mod
├── go.sum
├── handlers
│   ├── examplesHandler.go
│   ├── expressionHandler.go
│   ├── handler.go
│   └── historyHandler.go
├── main.go
├── models
│   ├── history.go
│   └── kalvium.db
├── README.md
└── views
    ├── example.html
    └── history.html
```

The Project strictly follows the MVC acrhitecture where:

- `configuration` -> config related setup like databases, etc
- `controllers` -> controller logic
- `handlers` -> handler logic that controls what HTTP status code and body needs to be returned.
- `models` -> schema models along with other DB logic to insert and query tables.
- `views` -> html views that are shown in the browser with populated data from server.

## How to run

There are 2 ways to run this project.

### Docker

- If you have docker installed then you can simply run the server using:

```sh
docker build -t expression-server . # This will build the docker image
docker run -p 8080:8080 expression-server # This will run the image and expose it on host machine port 8080
```

### Go Toolchain

- If you have the go toolchain installed you can simple run `go run main.go`
- Otherwise you can get the go toolchain and all the neccessary development deps by using the `nix` package manager and running `nix develop` in root of project. This will install `go` and `sqlite` in the current shell, which can be used to run the server.
