#!/usr/bin/env bash

# update packages
apt-get update
apt-get upgrade

# install general stuff
apt-get install --yes build-essential
apt-get install --yes python-software-properties
apt-get install --yes git

# install python stuff
apt-get install --yes python-dev
apt-get install --yes python-pip

echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
