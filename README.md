# rlyeh
At the moment some scripts to query osd/mon admin sockets and push them to ceilometer.

Depends on librados, expects the package to be installed in the system where it is run

# usage
```
rlyeh [<flags>]

Flags:
      --help     Show context-sensitive help (also try --help-long and --help-man).
      --osd_sock="/var/run/ceph/osd*asok"  
                 pattern for osd sockets
      --mon_sock="/var/run/ceph/mon*asok"  
                 pattern for mon sockets
      --conf="/etc/ceph/ceph.conf"  
                 path to configuration file
  -v, --verbose  verbose logging
      --log="/var/log/ceph/rlyeh.log"  
                 path to logfile
      --version  Show application version.
```
