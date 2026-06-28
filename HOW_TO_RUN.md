# Journey of Go — How to Run Files

## Run a single Go file
```
go run 01_basics/hello_world.go
go run 01_basics/variables.go
go run 01_basics/data_types.go
go run 02_control_flow/conditionals.go
go run 02_control_flow/loops.go
go run 03_functions/functions.go
go run 04_data_structures/arrays_slices.go
go run 04_data_structures/maps.go
go run 05_structs_interfaces/structs.go
```

## Build an executable
```
go build -o hello.exe 01_basics/hello_world.go
./hello.exe
```

## Check Go version
```
go version
```

## Folder Structure
```
journey_of_go/
├── 01_basics/
│   ├── hello_world.go    ← Start here!
│   ├── variables.go
│   └── data_types.go
├── 02_control_flow/
│   ├── conditionals.go
│   └── loops.go
├── 03_functions/
│   └── functions.go
├── 04_data_structures/
│   ├── arrays_slices.go
│   └── maps.go
└── 05_structs_interfaces/
    └── structs.go
```

## Learning Order
1. hello_world.go   — Your first Go program
2. variables.go     — Variables, types, constants
3. data_types.go    — All built-in types
4. conditionals.go  — if/else, switch
5. loops.go         — for loops (only loop in Go!)
6. functions.go     — Functions, closures, multiple returns
7. arrays_slices.go — Arrays and slices
8. maps.go          — Key-value maps
9. structs.go       — Structs, methods, interfaces
