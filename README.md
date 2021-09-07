# cli

[![Project Status: Concept – Minimal or no implementation has been done yet, or the repository is only intended to be a limited example, demo, or proof-of-concept.](https://www.repostatus.org/badges/latest/concept.svg)](https://www.repostatus.org/#concept)

[devexlabs cli](https://github.com/devexlabs/cli) is a tool designed to facilitate the installation and use of command line tools.

## Install

Download the latest release:

```bash
curl -sL https://github.com/devexlabs/cli/releases/download/v0.1-alpha/cli -o ~/cli
```

```bash
sudo mv ./cli /usr/local/bin
```

Run `init` command to choose and build docker with tools:

```bash
cli init
```

## Development

Dependencies:

- [golang](https://golang.org/doc/install)

Clone this repository:

```bash
git clone https://github.com/devexlabs/cli.git
```

Install dependencies:

```bash
go mod download
```

To build the binary run the command:

```bash
go build
```
