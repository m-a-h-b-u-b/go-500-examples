# go-500-examples

A curated collection of **500+ small, focused Go (Golang) examples** demonstrating practical language features,
design patterns, and real‑world utilities. Inspired by the Node‑500‑Examples project.

## Repository Structure

```
go-500-examples/
├── cmd/                 # CLI entry points
├── examples/            # Self-contained runnable examples
│   ├── 001_hello_world/
│   ├── 002_http_server/
│   └── ...
├── internal/            # Private reusable code
├── pkg/                 # Public reusable packages
├── docs/                # Architecture and contribution guides
├── go.mod
└── README.md
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

