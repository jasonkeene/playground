#!/bin/bash
rsync -avzci -e "ssh -p 2222 -o StrictHostKeyChecking=no -i $HOME/.vagrant.d/insecure_private_key" ./**/* core@127.0.0.1:/tmp
rsync -avzci -e "ssh -p 2200 -o StrictHostKeyChecking=no -i $HOME/.vagrant.d/insecure_private_key" ./**/* core@127.0.0.1:/tmp
rsync -avzci -e "ssh -p 2201 -o StrictHostKeyChecking=no -i $HOME/.vagrant.d/insecure_private_key" ./**/* core@127.0.0.1:/tmp
