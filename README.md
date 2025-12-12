# goreleaser-starter
A minimal Go project demonstrating how to build and publish cross-platform binaries using GoReleaser and GitHub Releases.

## Features

- Simple Go binary  
- Version/commit metadata injected via `ldflags`  
- Multi-platform builds (Linux, macOS, Windows)  
- Automatic archives and checksums  
- Automated GitHub Release publishing 

## Key Concepts

### 1. Git Tags = Releases

GoReleaser builds are triggered by **Git tags**, not branches.

```bash
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
```

### 2. GitHub Token

GoReleaser needs a GitHub Personal Access Token to publish releases.

Set it locally:

```
export GITHUB_TOKEN="your_token_here"
```

### 3. Running GoReleaser

Dry-run (no GitHub upload):

```
goreleaser release --snapshot --clean

```

Real release:

```
goreleaser release --clean

```


## Useful Commands

### Create a tag

```
git tag -a v0.1.1 -m "Release message"
```

### Push a tag

```
git push origin v0.1.1
```

### Delete a tag

```
git tag -d v0.1.0
git push origin :refs/tags/v0.1.0
```

### Run GoReleaser
```
goreleaser release --clean
```

---

## What GoReleaser Generates

1. **Archives** : Compressed .tar.gz or .zip files for each OS/architecture.
2. **Checksums** : SHA256 fingerprints for verifying binaries.
3. **Changelog** : Auto-generated from commits between tags.
4. **GitHub Release**: Includes binaries, archives, checksums, and the changelog.

---

### Summary

This project demonstrates the full GoReleaser flow:

- Write and commit your code
- Ensure the repo is clean
- Create and push a version tag
- Run GoReleaser to build and publish
- View the final artifacts in GitHub Releases

---