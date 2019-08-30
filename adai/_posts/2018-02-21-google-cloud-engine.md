---
title: "Run tasks in Google Cloud Engine even if SSH is terminated"
tags: 
  - SSH
  - google 
  - cloud engine
categories:
  - cloud  
last_modified_at: 2018-02-21
---

This post assumes that you have a project and a virtual machine (VM) instance already set up at [Google Cloud Engine](https://console.cloud.google.com/compute/instances?).

It easy to connect to your VM instance via SSH by clicking on the `SSH` link on the webpage. However, the tasks will be terminated when the SSH connection is terminated. 

Instructions are provided below to keep the tasks running even after SSH connection is lost. In the intsructions below, we will use an example of running a `python example.py` script.

Run a program with a `screen` to keep it running when you log out of your SSH session:
1. First, connect to your VM instance via SSH terminal
1. Install `screen` inside your compute instance: 
	```
	sudo apt-get install screen
	```
1. Then, run: 
	```
	screen python example.py
	```
	A new screen will be created where the output of `python example.py` will be printed. 
1. Detach this screen, to return to your SSH terminal, while the `python example.py` continues to run on the detached screen. To detach, enter 
	```
	Ctrl+a then Ctrl+d 
	```
1. Feel free to close or leave your SSH terminal
1. To resume, login into your VM instance via SSH terminal and reattach to the previously detached screen. List all screens using `screen -ls` and to reattach to a desired screen, enter `screen -r [screen name]`
1. To kill a detached screen: 
	```
	screen -X -S [screen name] quit
	```
1. To kill an attached screen: 
	```
	exit
	```
1. Enjoy!
