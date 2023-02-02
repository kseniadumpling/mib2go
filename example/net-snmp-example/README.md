# Net-SNMP examlpe MIB
This is MIB from [Net-SNMP official doc page](http://www.net-snmp.org/docs/mibs/NET-SNMP-EXAMPLES-MIB.txt).

## Imports
MIB declares these imports:
```
IMPORTS
    MODULE-IDENTITY, OBJECT-TYPE, Integer32,
    NOTIFICATION-TYPE                       FROM SNMPv2-SMI
    SnmpAdminString                         FROM SNMP-FRAMEWORK-MIB
    netSnmp                                 FROM NET-SNMP-MIB
    RowStatus, StorageType                  FROM SNMPv2-TC
    InetAddressType, InetAddress            FROM INET-ADDRESS-MIB
;
```
Hence, to run `min2go` you need to provide paths to SNMPv2-SMI, SNMP-FRAMEWORK-MIB, NET-SNMP-MIB, SNMPv2-TC and INET-ADDRESS-MIB. The typical paths for such MIBs is usually these ones: 
- /usr/share/snmp/mibs/
- /var/lib/snmp/mibs/ietf/

IANA specific can be found under
- /var/lib/snmp/mibs/iana/

## Generate
The command is
```
$ ./mib2go generate NET-SNMP-EXAMPLES-MIB --path +/var/lib/snmp/mibs/ietf,+/usr/share/snmp/mibs,+./example/net-snmp-example --dir ./example/net-snmp-example/generated/
```

It results with 4 files:
```
$ find ./example/net-snmp-example/generated/  -printf "%f\n"
generated/
snmp-framework-mib.go
inet-address-mib.go
net-snmp-examples-mib.go
snmpv2-tc.go
```