# Pod

The `methodk8s pod` family of commands provide information about an cluster's pods and containers.

## Enumerate

The enumerate command will gather information about all of the pods and their containers that the provided context have access to.

### Usage

```bash
methodk8s pod enumerate
```

### Help Text

```bash
$ methodk8s pod enumerate -h
Enumerate Pod objects

Usage:
  methodk8s pod enumerate [flags]

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
