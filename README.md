# vcs (name TBD) — a simplified Git in Go

## Why this project

Git internals (content-addressable storage, DAGs, hashing, tree structures) are
real systems concepts, scoped small enough to build solo. Go is new to me —
building this doubles as evidence I can ramp quickly on unfamiliar tech.

This is background/parallel prep work, not a priority over active job search.
No fixed timeline — built in steps, whenever there's time for it.

## Scope (v1 — deliberately narrow)

1. `init` — object store directory structure
2. `hash-object` — content-addressable blob storage (the core concept)
3. `commit` — tree object + commit object construction
4. `log` — walk the commit DAG backward
5. `diff` — basic diff between two commits

**Explicitly out of scope for v1:** branching, merging, remotes. These are
where scope explodes and aren't needed for the core learning goal.

## Object model decisions

- **Hashing:** SHA-1, to match real git and allow direct comparison against
  `git hash-object` output on the same content.
- **Object format:** content is prefixed with a header before hashing
  (`"blob <size>\0<content>"`), same as git — lets me sanity-check my hashes
  against real git's.
- **Storage layout:** `.vcs/objects/<hash>` — TBD whether to split into
  `<first-2-chars>/<remaining-38-chars>` subdirectories like git does, or keep
  flat for v1 simplicity. Decision to be logged when made.
- **Compression:** skipping zlib compression for v1 (plain bytes on disk).
  Cut for scope, not forgotten.

## Steps

- [ ] Step 0 — Go orientation (just enough to start)
- [ ] Step 1 — `init`: create the object store
- [ ] Step 2 — `hash-object`: the core concept
- [ ] Step 3 — tree objects
- [ ] Step 4 — `commit`
- [ ] Step 5 — `log`
- [ ] Step 6 — `diff`

## Log

Entries are short (3-5 sentences), written same-day, problem/why before how.
Decisions get captured with what was considered and rejected, not just what
was chosen. Anything that breaks, and what changed as a result, gets logged —
that's the actual interview material.

### YYYY-MM-DD — Project started

Scoped the object model on paper before writing any Go. Chose SHA-1 over
SHA-256 to keep output comparable to real git. Deferred the flat-vs-nested
storage layout decision to Step 1, when it'll actually matter.
# mini-vcs
