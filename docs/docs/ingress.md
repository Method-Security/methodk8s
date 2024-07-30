# Ingress

The `methodk8s ingress` family of commands provide information about an cluster's ingresses.

## Enumerate

The enumerate command will gather information about all of the ingresses that the provided context have access to. This includes gateway objects as well. If you only want to enumerate gateway objects please see flag options below

### Usage

```bash
methodk8s ingress enumerate --gateways
```

### Help Text

```bash
$ methodk8s ingress enumerate -h
Enumerate Ingress objects

Usage:
  methodaws ec2 enumerate [flags]

Flags:
      --gateways   Only include gateway objects
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
