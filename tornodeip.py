import requests
import os
import time

# Set these Values to your liking :)
const_tornode_path = os.environ['HOME'] + "/.tornode"
const_tortime_path = os.environ['HOME'] + "/.tortime"

def write_time_to_file(filename):
    current_time = time.strftime('%Y-%m-%d %H:%M:%S')
    with open(filename, 'w') as file:
        file.write(current_time)

def check_30_minutes_passed(filename):
	try:
		with open(filename, 'r') as file:
			last_time_str = file.read().strip()
			last_time = time.strptime(last_time_str, '%Y-%m-%d %H:%M:%S')
			current_time = time.localtime()
			time_difference = time.mktime(current_time) - time.mktime(last_time)
			#print(time_difference)
			if time_difference >= 1800:  # 30 minutes in seconds
				return True
	except FileNotFoundError:
		pass
	return False

# returns 1 on error
def get_torip_list(save,path):
	res = requests.get("https://www.dan.me.uk/torlist/?exit")
	print(res)
	if(res.status_code == 200):
		if save:
			file = open(path,"w")
			file.write(res.text)
			file.close()
		return res.text
	else:
		print("30 minutes have not yet passed please wait")
		return 1


# returns NONE if it fails, false if the ipaddr is not a tor ip, true if it is 
def ip_is_tor_node(ipaddr):
	firsttime = os.path.exists(const_tortime_path)
	if firsttime:
		passed = check_30_minutes_passed(const_tortime_path)
		ips=[]
		if passed:
			ips = get_torip_list(True,const_tornode_path)
			write_time_to_file(const_tortime_path)
		elif os.path.exists(const_tornode_path) == False:
				print("30 minutes have not passed AND the specified file does not exist")
				return None
		else:
			tornodesfile = open(const_tornode_path,"r")
			ips = tornodesfile.readlines()
			tornodesfile.close()
		if ips == 1:
			print("ERROR 42")
			return None
		for ip in ips:
			if(ip.strip() == ipaddr.strip()):
				return True
		return False
	else:
		#print("else")
		write_time_to_file(const_tortime_path)
		passed = check_30_minutes_passed(const_tortime_path)
		ips=[]
		if passed:
			ips = get_torip_list(True,const_tornode_path)
		elif os.path.exists(const_tornode_path) == False:
				print("30 minutes have not passed AND the specified file does not exist")
				return None
		else:
			tornodesfile = open(const_tornode_path,"r")
			ips = tornodesfile.readlines()
			tornodesfile.close()
		for ip in ips:
			if(ip.strip() == ipaddr.strip()):
				#print(ip)
				return True
		return False 