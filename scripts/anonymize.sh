#!/bin/bash
# This little script will contain some useful commands chained together to provide more privacy in public networks or something like that
echo "using this network device:" $1
if [ $2="--setup" ]; then
	sudo apt install tor ufw
fi



sudo ufw deny 9051
sudo ufw allow from 127.0.0.1 to any port 9051
sudo ufw enable
sudo ifconfig $1 down
sudo macchanger $1 -A #pretends to be a burned in addr and sets a random vendor MAC
sudo ifconfig $1 up # now mac addr is that of a random vender
sudo systemctl restart tor # restart here because if its running ... well itll restart, and if its not running restart will start it
sleep 8 # wait for tor to properly start
# Then get a new circuit, but also setting a password for tor would be a REALLY good idea, OR setting a firewall rule ?
echo -e 'AUTHENTICATE ""\r\nsignal NEWNYM\r\nQUIT' | nc 127.0.0.1 9051
# now for any tcp related programs you can just use torify YOUR_PROGRAM YOUR_PROGRAM_ARGS
