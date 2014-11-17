#!/usr/bin/env bash

# update packages
#apt-get update
#apt-get upgrade

# install salt minion
/vagrant/bin/install-minion.sh

# link in salt states
ln -s /vagrant/salt /srv/salt

# configure salt minion to bootstrap
cat << EOF > /etc/salt/minion
file_client: local
file_roots:
  base: /srv/salt
grains:
  roles: master
EOF

# highstate!
salt-call --local state.highstate

# copy over test_key
TEST_KEY=/vagrant/salt/test_key.pub
MINION_KEY_PATH=/etc/salt/pki/master/minions/
echo ${MINION_KEY_PATH}web{1,2} | xargs -n1 cp $TEST_KEY

echo
echo
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
echo "Provision complete!"
