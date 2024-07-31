# Service

The `methodk8s service` family of commands provide information about an cluster's services and containers.

## Enumerate

The enumerate command will gather information about all of the services that the provided context have access to.

### Usage

```bash
methodk8s service enumerate
```

### Help Text

```bash
$ methodk8s service enumerate -h
Enumerate Service objects

Usage:
  methodk8s service enumerate [flags]

Flags:
  -h, --help       help for enumerate

Global Flags:
  -c, --context string       cluster url
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -p, --path string          config path
  -q, --quiet                Suppress output
  -u, --url string           cluster url
  -v, --verbose              Verbose output
```
