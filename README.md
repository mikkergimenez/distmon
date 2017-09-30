# DistMon Distributed Monitoring

# Dependencies
distmon depends on libstatgrab, which will have to be installed on your server.  This library may be funky on certain distros of Linux.

# Using:
distmon listens on port 55556

# Config:
distmon reads the config out of /etc/distmon.toml
The format is as such:


Hostname=hostname1.example.com

Peers = [ "hostname0.example.com", "hostname.example.com", "hostname.example.com" ]
