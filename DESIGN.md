# Design

Companion to README.md. README tracks scope/status/log; this tracks the
system's shape — requirements and architecture — so implementation has
something concrete to check against.

## Functional Requirements

| # | Command | Behavior |
|---|---|---|
| FR1 | `vcs init` | Creates `.vcs/objects/` in cwd. Errors if `.vcs/` already exists. |
| FR2 | `vcs hash-object [-w] <file>` | Computes SHA-1 over `"blob <size>\0<content>"`. Always prints the hash. With `-w`, writes it to the object store; without, dry-run only (matches real git — useful for sanity-checking against `git hash-object`). |
| FR3 | `vcs commit -m "<msg>"` | Snapshots the entire current working directory (no staging area — see D1) into tree object(s), builds a commit object (tree hash + parent hash + message + timestamp), updates `.vcs/HEAD` to the new commit hash, prints the new hash. |
| FR4 | `vcs log` | Starts at `.vcs/HEAD`, walks parent links backward, prints each commit until reaching one with no parent (DAG root). |
| FR5 | `vcs diff <commit-a> <commit-b>` | Compares the two commits' trees; reports added/removed/changed files. Line-level diff can stay basic (report "changed", not a full unified diff) unless extended later. |

## Non-Functional Requirements

- **Stdlib only** — no external deps. The point is to go through `crypto/sha1`,
  `os`, `flag`, `path/filepath`, `bufio`/`io` directly.
- **Git-hash-compatible for blobs** — `hash-object` output must match real
  `git hash-object` byte-for-byte on identical content. This is the
  verification oracle for Step 2.
- **No compression, no concurrency, no staging area** — single-threaded CLI,
  plain files on disk.
- **Local-only, single-user, small repos** — no performance requirements.
  Optimize for readable code over speed.
- **Linux/WSL target only** (per SETUP.md) — no Windows path handling needed.

## Design Decisions

Same convention as README's "Object model decisions" — what was considered,
what was picked, and why.

### D1 — No staging area

Considered: mirroring git's working-dir → index → commit two-step flow.
Rejected: adds a whole object (the index) and command (`add`) not in scope.
Chosen: `commit` snapshots the entire current working directory as-is, every
time. Consistent with cutting branching/merging/remotes for the same reason —
scope control, not because staging isn't a real concept worth learning later.

### D2 — HEAD as a single flat pointer file

Considered: git's actual model, `HEAD` → `refs/heads/<branch>` → commit hash
(a layer of indirection to support multiple branches).
Rejected: branches are explicitly out of scope for v1, so the indirection has
no purpose here.
Chosen: `.vcs/HEAD` holds exactly one commit hash — the tip of the (single,
linear) history. Absent/empty file means no commits yet, i.e. the DAG-root
case `commit` and `log` both need to handle.

### D3 — Tree objects: text format, not git's binary format

Considered: matching git's actual tree encoding (mode + filename + raw
20-byte hash, packed binary) for full wire-compatibility.
Rejected: harder to debug (can't `cat` it), and the README's git-compatibility
requirement (FR2) only commits to *blob* hashing, not the whole object model.
Chosen: plain text, one line per entry — e.g. `blob <hash> <name>` /
`tree <hash> <name>` for a subdirectory. Trades git-compatibility for
readability while a beginner is still debugging their own encoder.

### D4 — `commit` excludes `.vcs/` when walking the working directory

Stated explicitly so it's a requirement, not a bug discovered later: the tree
walk for a commit must skip the `.vcs/` directory itself, or the repo starts
hashing its own object store into every commit.

### D5 — Object storage layout: flat, not sharded

(Carried over from README, restated here since it's part of the object
model.) `.vcs/objects/<full-hash>`, no `<first-2>/<remaining-38>` split.
Sharding is a filesystem-performance workaround for repos with huge object
counts — irrelevant at this scale, and the path-construction function is a
one-line change later if it's ever needed. See README's "Storage layout" note.

## Architecture

```
                 ┌─────────────┐
   CLI args ───▶ │   main.go   │  dispatch only — no logic
                 └──────┬──────┘
                        │
                 ┌──────▼───────────┐
                 │ internal/commands│  per-subcommand arg parsing,
                 │  init / hash-obj │  orchestration, output
                 │  commit/log/diff │
                 └──────┬───────────┘
                        │
                 ┌──────▼───────────┐
                 │ internal/objects │  hashing; encode/decode for
                 │                  │  blob·tree·commit; object-store
                 │                  │  read/write; HEAD read/write
                 └──────┬───────────┘
                        │
                 ┌──────▼──────┐
                 │ filesystem  │  .vcs/objects/<hash>, .vcs/HEAD
                 └─────────────┘
```

`internal/commands` never touches the filesystem directly — only
`internal/objects` does. Same "one seam" principle as D5's path-construction
function, one layer up: change the storage details in one place without
touching every command.

## Object Model Reference

Quick summary of what each object type looks like on disk, once built (fully
defined step-by-step as Steps 1-4 land — see README).

- **blob**: `"blob <size>\0<content>"`, SHA-1 hashed, raw bytes stored as-is
  (no compression — see README).
- **tree**: text, one line per entry: `<type> <hash> <name>`, where `<type>`
  is `blob` or `tree` (recursive — subdirectories are nested tree objects).
- **commit**: text, fields for `tree <hash>`, `parent <hash>` (omitted for the
  root commit), `message`, `timestamp`.
