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
  -h, --help   help for enumerate

Global Flags:
  -a, --cert string          Base64 encoded ca certificate
  -c, --context string       Cluster context (ie. minikube)
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -p, --path string          Absolute or relative path to the config file (ie. ~/.kube/config)
  -q, --quiet                Suppress output
  -e, --server-url string    Cluster server url
  -s, --service-account      Set to true if using service account workflow
  -t, --token string         Base64 encoded service account token
  -u, --url string           Cluster url (ie. mycluster.com)
  -v, --verbose              Verbose output
```
