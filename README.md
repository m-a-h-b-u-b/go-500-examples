# go-500-examples

A curated collection of **500+ small, focused Go (Golang) examples** demonstrating practical language features,
design patterns, and real‑world utilities. Inspired by the Node‑500‑Examples project.

## Repository Structure

```
go-500-examples/
│
├── README.md
├── LICENSE
├── CONTRIBUTING.md
├── basics/
│   ├── 001_hello_world.go
│   ├── 002_variables.go
│   └── ...
├── concurrency/
│   ├── 101_goroutines.go
│   ├── 102_channels.go
│   └── ...
├── networking/
│   ├── 201_http_server.go
│   ├── 202_tcp_client.go
│   └── ...
├── databases/
│   ├── 301_postgres.go
│   ├── 302_mongodb.go
│   └── ...
├── web/
│   ├── 401_rest_api.go
│   └── ...
├── algorithms/
│   ├── 501_sorting.go
│   ├── 502_graph_bfs.go
│   └── ...
└── etc...

```

## Getting Started

1. **Clone the repo**

   ```bash
   git clone https://github.com/<your-username>/go-500-examples.git
   cd go-500-examples
   ```

2. **Run an example**

   ```bash
   cd examples/001_hello_world
   go run main.go
   ```

3. **Add a new example**

   - Create a new folder under `examples/` with a sequential prefix (e.g., `003_new_feature`).
   - Include a `main.go` that compiles independently.
   - Update `docs/architecture.md` if necessary.

## Contribution Guidelines

- Follow idiomatic Go code style (`gofmt`, `go vet`).
- Keep each example minimal and self-explanatory.
- Use clear comments describing the concept being demonstrated.
- Submit a PR with a short description of the example and any dependencies.

## License

Apache 2.0 License. See [LICENSE](LICENSE) for details.

