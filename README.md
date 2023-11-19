# GIST Tool

GIST Tool - is a minimalistic CLI utility for fast publishing GitHub gists via terminal.

## Installing:

----
Clone this repository: 
```bash
git clone git@github.com:jus1d/gist-tool.git
```

---
Download all dependencies:
```bash
go mod download
```

--- 
#### Go to **GitHub > Setting > Developer Settings > Personal Access Tokens** and generate new access token with `gist: create` scope.

---

#### Paste your access token to TOKEN const in `./cmd/gist/main.go`:

---
#### Build the application:

```bash
make build
```
