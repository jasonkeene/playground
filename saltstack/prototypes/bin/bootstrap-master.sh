#!/usr/bin/env bash

# update packages
#apt-get update
#apt-get upgrade --yes

# install salt minion
/vagrant/bin/install-minion.sh

# link in salt states and pillar data
ln -s /vagrant/salt /srv/salt
ln -s /vagrant/pillar /srv/pillar

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

# set hostname
salt-call network.mod_hostname master

# copy over test_key
MINION_KEYS=$(echo /etc/salt/pki/master/minions/{proxy{1,2},app{1,2}})
cat /vagrant/salt/minion.pub | tee $MINION_KEYS > /dev/null

echo
echo
echo "Provision complete!"
