# DistMon Distributed Monitoring

## About

   Distmon is an experiment in distributed monitoring.  It might be useful for monitoring a small number of servers in Digital Ocean without running something like Nagios, but it probably wouldn't work well in a production environment.  The idea is that each server runs an independent agent, and when you query a server, it queries all other servers and builds an index page with an overview list of the cpu, memory utilization, and number of docker containers(this could probably change depending on the plugins you decide to use) for each server.  When you click on a server, it will redirect to distmon on that server and give you more detailed monitoring information.  

Directions I would like to take this include:

* Consensus based alerting -- Nodes work together to determin if a node is down, or there's some type of network issue.
* Private servers and topology awareness -- A public node could reach private nodes and proxy traffic through them so you didn't have to expose all nodes.  Also if some nodes were in a different, unavailable network, you could identify an intermediary node as a proxy node to reach those otherwise unreachable nodes.
* Plugin Architecture -- Be able to monitor different things based on configuration

## Dependencies
distmon depends on libstatgrab, which will have to be installed on your server.  This library may be funky on certain distros of Linux.

## Using:
distmon listens on port 55556

## Config:
distmon reads the config out of /etc/distmon.toml
The format is as such:


Hostname=hostname1.example.com

Peers = [ "hostname0.example.com", "hostname.example.com", "hostname.example.com" ]
