
### Three-node architecture with OpenStack Networking (neutron)

**Requirements for VMs**:

- Controller Node: 1 processor, 2 GB memory, and 5 GB storage
- Network Node: 1 processor, 512 MB memory, and 5 GB storage
- Compute Node: 1 processor, 2 GB memory, and 10 GB storage
- Make sure your VM's have an MTU of 1450 (GRE tunnels add to the packet size)
