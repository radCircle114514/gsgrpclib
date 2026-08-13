# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`gsgrpclib` is a shared Go library that acts as a **gRPC client** for four internal Gaosheng backend services: **booking**, **business**, **track**, and **wecom**. It is consumed by other projects in the `Gaosheng-Service-Projects` workspace via a local `replace` directive:

```
replace github.com/iznilul/gsgrpclib => /Users/mac/Gaosheng-Service-Projects/gsgrpclib
```

There are no servers, `main` packages, or tests in this module — only proto definitions, generated code, low-level gRPC connection helpers, and higher-level "invoke" wrappers.

## Common commands

```bash
# Build everything
go build ./...

# Regenerate Go gRPC code after editing a .proto (run from repo root)
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/<service>/<service>-rpc.proto

# e.g. for wecom
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/wecom/wecom-rpc.proto
```

`protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` are expected to be on `PATH`. No Makefile or test suite exists.

## Architecture

### The three layers

1. **`proto/<service>/`** — gRPC service definitions (`.proto`) plus the generated `*.pb.go` and `*_grpc.pb.go`. Each service defines a universal `RequestAO` / `ResponseVO` pair where all payload fields are `google.protobuf.Any` (see `proto/any.proto`). This lets RPC methods pass arbitrary JSON-serializable Go values without defining typed messages per method. A few methods additionally use dedicated typed `*AO` messages (e.g. `FindCustomerAO`, `GetApprovalDetailAO`).

2. **`client/`** — one `<service>_rpc_client.go` per service plus `config.go`. Each file exposes:
   - `Init<Service>RpcClientConn()` — dials the host from config (TLS or insecure) and returns a `*grpc.ClientConn`. **The caller is responsible for `conn.Close()`** (all wrappers `defer` it).
   - `Invoke<Service>RPCMethod(ctx, methodName, ao)` — a **reflection-based dispatcher**: it opens a connection, constructs the generated gRPC client, then calls `reflect.ValueOf(client).MethodByName(methodName)`. This is why method-name strings appear throughout the invoke layer — they must match the RPC method name in the `.proto` exactly.
   - `SetTimeout(ctx)` (in `config.go`) — wraps the context with a timeout. In `dev` mode it uses an absurdly long `114514 * time.Second`; in `test`/`prod` it uses 10 seconds.

   `config.go`'s `init()` reads `conf/config.<mode>.toml` via viper, where `<mode>` is selected by the **first CLI argument** (`prod`, `test`, or default `dev`). Because the consuming application's `os.Args` drives this, running host apps with an unexpected first arg can silently fall back to `dev`.

3. **`invoke/<service>/`** — typed, ergonomic Go wrappers around each RPC. This is the layer consuming projects actually import. Functions follow a consistent shape: marshal Go values into `RequestAO` via the `utils.Parse*ToAny` helpers, call `client.Invoke<Service>RPCMethod`, then unmarshal the `Any` fields back into Go maps/slices. Some wrappers for methods with dedicated `*AO` messages bypass the reflection helper and call the generated gRPC client directly (e.g. `InvokeRpcFindCustomerList`, `InvokeRpcAddCustomer`).

   Functions within each service are split into category files (e.g. `invoke/wecom/customer.go`, `invoke/wecom/approval.go`, `invoke/wecom/check_in.go`). The root `invoke` package only contains `doc.go` and `aliases.go`; `aliases.go` re-exports every subpackage function as a package-level `var` for **backwards compatibility** with code that imported the old monolithic `invoke` package. New code should import the subpackages directly (e.g. `github.com/iznilul/gsgrpclib/invoke/wecom`).

### `utils/`

`parse_tools.go` is the bridge between Go values and `google.protobuf.Any`. It uses `encoding/json` (via `github.com/mumushuiding/util`) rather than protobuf's native `Any` packing: values are JSON-marshaled into the `bytes value` field. Key pairs:
- `ParseDataToAny` / `ParseMapToAny` / `ParseMapIntToAny` — Go → `*any.Any`
- `ParseAnyToMap` / `ParseAnyToMapList` / `ParseAnyToData` — `*any.Any` → Go

**These helpers `panic` on unmarshal errors**, so invoke functions typically only return errors from the gRPC call itself; malformed payloads crash rather than surface as errors.

### `constant/`

Small string constants (customer states, report genres) shared across services.

## Conventions when adding an RPC

1. Add the `rpc` to the relevant `proto/<service>/<service>-rpc.proto`. Prefer `RequestAO`/`ResponseVO` unless a dedicated typed message already exists or is warranted.
2. Run `protoc` to regenerate (command above).
3. Add an invoke wrapper in the appropriate `invoke/<service>/<category>.go` file. Use `client.Invoke<Service>RPCMethod(ctx, "MethodName", ao)` for `RequestAO`-based methods; for typed `*AO` messages, follow the direct-connection pattern (`Init...ClientConn` → `New...Client` → `SetTimeout` → call method).
4. If using a new dedicated message type, mirror the pattern in `invoke/wecom/customer.go` rather than the reflection helper.
5. Regenerate `invoke/aliases.go` if a new exported function is added (the re-export list is not automatic).
6. Run `go build ./...` to verify.

## Naming quirks

- The codebase mixes `Rpc`/`RPC` casing and contains intentional misspellings that must be preserved to match wire/proto names (e.g. `GetUserIdentitfy`, `FindCustomerCouldBeSelectedList`). Do not "fix" these without updating the proto and all callers in lockstep.
- WeCom user-facing terms: `openID` (external contact), `userID` (internal employee), `spNo` (approval/order number), `contentNo` / `serialNumber` (tracking identifiers).
