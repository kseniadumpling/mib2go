# mib2go

## Build
Simply run `go build` in the root of the project

## Usage
```
$ ./mib2go generate --help
Generates Go files from MIBs.
 
Usage:
  mib2go generate <MIB-name> [flags]
 
Flags:
  -d, --dir string       Output directory (default ".")
  -h, --help             help for generate
  -o, --output string    Output filename, use - for stdout
  -p, --package string   The package for the generated file (default "mibs")
  -M, --path strings     Path(s) to add to MIB search path
 
Global Flags:
      --config string   config file (default is $HOME/.mib2go.yaml)
```

Please note that symlinks under `/usr/share/snmp/mibs/` won't work, so for any SNMP**v2** MIBs (or any other MIB under iana/ietf) you'll need to exclusively add a path, e.g.:
```
./mib2go generate -M +/var/lib/snmp/mibs/iana,+/var/lib/snmp/mibs/ietf IF-MIB
``` 

