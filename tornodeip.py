import requests
import os
import time


# Set these Values to your liking :)
const_tornode_path = os.environ['HOME'] + "/.tornode"
# returns 1 on error
def get_torip_list(save,path):
	res = requests.get("https://www.dan.me.uk/torlist/")
	if(res.status_code == 200):
		if save:
			file = open(path,"w")
			file.write(res.text)
			file.close()
		return res.text
	return None

# returns NONE if it fails, false if the ipaddr is not a tor ip, true if it is 
def ip_is_tor_node(ipaddr):	
	tornodesfile = open(const_tornode_path,"r")
	ips = tornodesfile.readlines()
	tornodesfile.close()
	for ip in ips:
		if(ip.strip() == ipaddr.strip()):
			return True
	return False
