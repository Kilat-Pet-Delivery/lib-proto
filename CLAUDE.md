# Kilat Pet Delivery - lib-proto

The contract between services: CloudEvents schemas for every Kafka topic, and the cross-service DTOs (AddressDTO, PetSpecDTO and friends) that request and response bodies are built from.
Jira project **KPD** - GitHub `Kilat-Pet-Delivery/lib-proto` - stack **Go 1.24 - library**. Global rules live in `~/.claude/`;
this file only adds what is specific here.

## Orient here first

- `.claude/memory/project_state.md` - **resume here** (`/continue` reads it, `/recap` rewrites it).
- `README.md` - how to run it. `CHANGELOG.md` - what changed.
- The workspace map: `~/Documents/kilat-pet-delivery/CLAUDE.md`.

## Commands

| Task | Command |
|---|---|
| install | `go mod download` |
| test | `go test ./...` |
| integration tests | none in this repo |
| lint | `gofmt -l . && go vet ./...` |
| build | `go build ./...` |
| migrate | n/a - this repo owns no schema |

Needs no running stack: `go build ./...` and `go test ./...` work on a bare checkout.

## Conventions that differ from the global rules

- **Ticket branches and PRs** - company repo, never commit on `main` (`branch-guard` enforces it).
- This repo is imported by every service through a `replace` directive (`../lib-common`, `../lib-proto`), so CI has to check it out as a sibling - see the reusable workflow in lib-common.
- Protected paths (never edited in place, see `.claude/protected-paths.txt`): `migrations/*.sql`.

## Testing

`go test ./...` - 54 passing.

## Where things are

- `events/` one file per topic (booking_events.go, payment_events.go and so on) - `dto/` shared request and response shapes, with `dto/shop/` for the merchant surface

## Worth knowing

- First in the cross-repo change order: an event or DTO changes here before any service uses it. Builds alone.
