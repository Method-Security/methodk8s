# Node

The `methodk8s node` family of commands provide information about an cluster's nodes and containers.

## Enumerate

The enumerate command will gather information about all of the nodes that the provided context have access to.

### Usage

```bash
methodk8s node enumerate
```

### Help Text

```bash
$ methodk8s node enumerate -h
Enumerate Node objects

Usage:
  methodk8s node enumerate [flags]

Flags:
  -h, --help       help for enumerate

Global Flags:
  -c, --context string       cluster context
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -p, --path string          config path
  -q, --quiet                Suppress output
  -u, --url string           cluster url
  -v, --verbose              Verbose output
```
