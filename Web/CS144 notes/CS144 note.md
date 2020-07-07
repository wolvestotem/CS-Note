## Stanford CS144

1-1 applications
- www model: a reliable, bidirectional byte stream
world wide web(client->server)
- bittorrent(peer to peer): torrent(info about data files) -> tracker(list of clients with pieces of data files) -> clients
(through HTTP)
- skype(mix): client A want connection to client B(behind NAT)
A -> rendezvous server -> B; B -> A (reverse connection)
client A(behind NAT) want connection to client B(behind NAT)
A -> relay ->A; B -> relay -> B

1-2 The 4 layers Internet Model
- Internet made up of: end-hosts, routers, links
![Internet model](./pictures/Internet.png)
- Link layer: to carry the date over one link at a time
- Network layer: to deliver packets end-to-end across the Internet from the source to destination
![Network layer](./pictures/Networklayer.png)
- Internet Protocol(IP)
IP makes a attempt to deliver datagrams to destination, but it makes no promises.
IP datagrams can get lost, out of order, or corrupted. No guarantees.
- Transport layer: 
TCP 提供可靠性
UDP video无需可靠性
- Application layer
- layer: 封装 reuse
![Summary](./pictures/Summary4layer.png)

1-3 Network layer(IP)
property:
- Datagram (routing to the destination: IP SA, IP DA)
- Unreliable
- best effort(postal)(only drop if necessary, no detection)
- Connectionless

IP header:
TTL(time to live):decrease each router to prevent loop

1-5 Packets switching principles
Packets: IP包
statistical multiplexing:
- packet switching allows flows to use all available link capacity
- Packet switching allows flows to share links capacity

Summary:
- Simple: forward flows independently, don't need ot know about the flows
- efficient: share capacity among many flows

1-6 Layering principles
Each layer provides servises to the upper layer, using the servises of the layers below and its own private proceedings.

1-7 Encapsulation
layers and packets switching
![Encapsulate](./pictures/Encapsulate.png)

1-8 Bytes order
