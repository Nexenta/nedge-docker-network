Docker network plugin for NexentaEdge

Building:
	% git clone https://github.com/Nexenta/nedge-docker-network.git
	% make deps
	% make 

Running:
	% cp ndnet/daemon/ndnet.json /opt/nedge/etc/ccow/ndnet.json
	% bin/ndnet daemon start

