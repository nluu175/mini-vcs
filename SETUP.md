# Setup

## Environment: WSL, not native Windows

Using WSL instead of native Windows Go. More consistent with how Go tooling,
file permissions, and CLI tools behave on Linux/real deployments, and it
matches what most backend/systems roles use anyway. Avoids path-separator
weirdness and lets standard Unix tooling just work.

## Steps

1. **Confirm WSL is set up.** Open the WSL terminal (Ubuntu) directly, or run
   `wsl --version` from PowerShell to check the version first.

2. **Install Go inside WSL — not the Windows `.msi` installer.** Grab the
   Linux tarball instead:

   ```bash
   wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
   sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
   ```

   Note: check https://go.dev/dl/ for the current version before running —
   this updates every few months.

3. **Add Go to PATH.** Append to `~/.bashrc` (or `~/.zshrc` if using zsh):

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```

   Then reload: `source ~/.bashrc`.

4. **Verify install:**

   ```bash
   go version
   ```

   Should print something like `go1.23.4 linux/amd64`.

5. **Keep project files inside the WSL filesystem** (e.g. `~/projects/vcs`),
   not under `/mnt/c/...`. Working across the Windows/WSL boundary is
   noticeably slower and can cause odd file-watching/permission issues.

## Status

- [ ] WSL confirmed
- [ ] Go installed in WSL
- [ ] PATH configured
- [ ] `go version` verified
- [ ] Project directory created inside WSL filesystem