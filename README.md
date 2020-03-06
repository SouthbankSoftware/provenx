# provenx-cli

`provenx-cli` is a simple CLI for ProvenX API Service (`provenx-api`)

## Usage

_**Currently**_, `provenx-cli` connects directly to the dev `provenx-api` with TLS, which can be configured via CLI options, and magic auth token

### Download dev binaries

- [mac](https://storage.googleapis.com/provendb-dev/provenx-cli/provenx-cli_darwin_amd64)
- [linux](https://storage.googleapis.com/provendb-dev/provenx-cli/provenx-cli_linux_amd64)
- [windows](https://storage.googleapis.com/provendb-dev/provenx-cli/provenx-cli_windows_amd64.exe)

### Build your own binary

```bash
# generate the `provenx-cli` binary
make
```

### Examples

```bash
# for help
./provenx-cli -h

# create a trie for a path
./provenx-cli create trie path/to/the/data

# create a trie for a path in a custom trie location
./provenx-cli create trie path/to/the/data -t path/to/output/the/trie.pxt

# create a trie for a path including metadata
./provenx-cli create trie path/to/the/data --include-metadata

# verify a trie for a path
./provenx-cli verify trie path/to/the/data

# verify a trie for a path and output the trie's Graphviz Dot Graph
./provenx-cli verify trie path/to/the/data -d path/to/output/the/dot/graph.dot

# verify a trie for a path from a custom trie location
./provenx-cli verify trie path/to/the/data -t path/to/the/trie.pxt
```

## FAQ

### Error: "provenx-cli_darwin_amd64" cannot be opened because the developer cannot be verified

![Mac Cannot Open Issue](docs/mac_cannot_open_issue.png)

Use the following command to fix:

```bash
xattr -d com.apple.quarantine path/to/provenx-cli_darwin_amd64
```
