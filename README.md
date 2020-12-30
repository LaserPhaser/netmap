# netmap
Simple version of portscanner and netmap in GO


Example of usage:

````
./netmap -port_start 0 -port_end 40000 -scan_protocol tcp -host localhost  
Scanning started for localhost for ports range[0:40000], protocol tcp
Found open port 1234
Found open port 7742
Found open port 5600
Found open port 65530
Scanning finished

