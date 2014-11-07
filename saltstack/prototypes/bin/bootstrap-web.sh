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

# install salt minion
add-apt-repository ppa:saltstack/salt
echo deb http://ppa.launchpad.net/saltstack/salt/ubuntu `lsb_release -sc` main | tee /etc/apt/sources.list.d/saltstack.list
wget -q -O- "http://keyserver.ubuntu.com:11371/pks/lookup?op=get&search=0x4759FA960E27C0A6" | apt-key add -
apt-get update
apt-get install --yes salt-minion

# configure salt minion to pull from master
cat << EOF > /etc/salt/minion
master: $1
grains:
  roles:
    - webserver
EOF

# restart miniond
service salt-minion restart

echo
echo
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
