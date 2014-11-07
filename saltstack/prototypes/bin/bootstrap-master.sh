#!/usr/bin/env bash

## update packages
#apt-get update
#apt-get upgrade
#
## install general stuff
#apt-get install --yes build-essential
#apt-get install --yes python-software-properties
#apt-get install --yes git
#
## install python stuff
#apt-get install --yes python-dev
#apt-get install --yes python-pip

# install salt minion
add-apt-repository ppa:saltstack/salt
echo deb http://ppa.launchpad.net/saltstack/salt/ubuntu `lsb_release -sc` main | tee /etc/apt/sources.list.d/saltstack.list
wget -q -O- "http://keyserver.ubuntu.com:11371/pks/lookup?op=get&search=0x4759FA960E27C0A6" | apt-key add -
apt-get update
apt-get install --yes salt-minion

# stop miniond
service salt-minion stop

# link in salt states
ln -s /vagrant/salt /srv/salt

# configure salt minion to bootstrap
cat << EOF > /etc/salt/minion
file_client: local
file_roots:
  base:
    - /srv/salt
grains:
  roles:
    - master
EOF

# highstate!
salt-call --local state.highstate

echo
echo
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
