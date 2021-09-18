# netmaker-gui
Alternative UI for netmaker (see github.com/gravitl/netmaker)

Built with go and html/templates.
Missing following features compared to netmaker-ui
DNS
Ingress/Egress Nodes
External Clients
Update User
Relay Gateways (available in next version of netmaker)

Installation:
Docker-compose files
docker-compose.full.yml
netmaker
netmaker-ui
netmaker-gui
rqlite

docker-compose.sqlite.yml
netmaker
netmaker-ui
netmaker-gui

docker-compose.yml (uses sqlite)
netmaker
netmaker-gui

required ports:
netmaker 50581 (gprc), 8081(api), 8082(netmaker-ui), 8080(netmaker-gui)

It is left as an exercise to the reader to set up a reverse proxy.
see Netmaker docs (https://doc/netmaker.org or https://netmaker.readthedocs.io)



