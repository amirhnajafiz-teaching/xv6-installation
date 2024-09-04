#!/bin/bash


#### base functions ####
sudo_apt_get()
{
	sudo apt-get install $1 -y
}

#### base variables ####
pdeps=("git", "wget", "qemu", "gdb")


# install dependencies
for dep in ${pdeps[@]}; do
	echo "Installing $dep ..."
	sudo_apt_get $dep
done

# clone into public XV6 repo
git clone https://github.com/mit-pdos/xv6-public.git

