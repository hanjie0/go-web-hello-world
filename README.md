This is a demo project required by SRE role. 

The candidate should be able to complete the project independently in two days and well document the procedure in a practical and well understanding way.

It is not guaranteed that all tasks can be achieved as expected, in which circumstance, the candidate should trouble shoot the issue, conclude based on findings and document which/why/how.

### Task 0: Install a ubuntu 16.04 server 64-bit

either in a physical machine or a virtual machine

http://releases.ubuntu.com/16.04/<br>
http://releases.ubuntu.com/16.04/ubuntu-16.04.6-server-amd64.iso<br>

	[hanjie:] Note: The above mentioned version: ubuntu-16.04.6-server-amd64.iso does not exist. 
			  download ubuntu-16.04.7-server-amd64.iso instead. 	

https://www.virtualbox.org/
	[Hanjie:] 
			  Hanjie:] downloaded VirtualBox-6.1.16-140961-Win.exe from https://www.virtualbox.org/wiki/Downloads. The package will be installed on Window10
			  precondition: local administrator. 
			  Problems/procedures:
				1. there is existing Vagrant/virtualBox installed, uninstall them before installation.  
				2. Install virtualBox with the default settings by keeping click "Next".
			   	3. Start "Oracle VM VirtualBox Manager" and setup Proxy if needed. 
						click "New" button to create a new VM.  select "Linux" as Type and then select "ubuntu xxx"
						Note: only 32bit versions are available, which is a window10 specific issue.
						google the problems and find solution to Suspend/turn off HyperV settings. Restart the laptop.	
						And create the new VM again. Now. 64bit versions are available. 
				4. create the VM with the following settings: 4096M memory, 10G hard disk, VDI, Dynamically allocated, and keep default settings for others. 
						Anyway, we can change the settings later. 
				5. Add the download ubuntu-16.04.7-server-amd64.iso as the startup disk  
				6. install ubuntu server with English language .	
				7. keep the default settings.  Google the ubuntu installation guide if you don’t know the meaning of some of the settings. 
					Create the VM with user:password <jihan:jihan>

				8.Save the machine state and restart the VM and 
					Run “sudo apt-get install openssh-server” to install ssh server for next steps(task1). 
	[hanjie:] 	done	


			  	

for VM, use NAT network and forward required ports to host machine
- 22->2222 for ssh
- 80->8080 for gitlab
- 8081/8082->8081/8082 for go app
- 31080/31081->31080/31081 for go app in k8s

	[hanjie:] 	done	
	
	
	

### Task 1: Update system

ssh to guest machine from host machine ($ ssh user@localhost -p 2222) and update the system to the latest
	[Hanjie:] 
		Run “sudo apt-get install openssh-server” to install ssh server for next steps(step1). 
		Here I use MobaXterm instead. Create a ssh session and login to the VM.
		Setup proxy if needed. 



https://help.ubuntu.com/16.04/serverguide/apt.html

upgrade the kernel to the 16.04 latest

	[Hanjie:] 
		sudo apt update
		sudo apt upgrade
	[Hanjie:] done

### Task 2: install gitlab-ce version in the host
https://about.gitlab.com/install/#ubuntu?version=ce

	[Hanjie:] 
		sudo apt-get install -y curl openssh-server ca-certificates tzdata                              done
		sudo apt-get install -y postfix               						 skipped	
		curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | sudo bash                  done and need proxy settings
		
		#and change the command to preserve the ENV of the proxy settings to sudo 
		curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | sudo bash -E  	
		
		sudo EXTERNAL_URL="http://gitlab.example.com" apt-get install gitlab-ce 	done                           
		#change from https to http.

Expect output: Gitlab is up and running at http://127.0.0.1 (no tls or FQDN required)

Access it from host machine http://127.0.0.1:8080

	[Hanjie:] done

### Task 3: create a demo group/project in gitlab

named demo/go-web-hello-world (demo is group name, go-web-hello-world is project name).

Use golang to build a hello world web app (listen to 8081 port) and check-in the code to mainline.

https://golang.org/<br>
https://gowebexamples.com/hello-world/

Expect source code at http://127.0.0.1:8080/demo/go-web-hello-world

	[Hanjie:] done with GUI operations. 
 
### Task 4: build the app and expose ($ go run) the service to 8081 port

	[Hanjie:] done
		minor change on the source codes
		Install go lang by command “sudo apt install golang-go”
		go build main.go
		./main
		export no_proxy="localhost,127.0.0.1, gitlab.example.com"		
		add 127.0.0.1 gitlab.example.com to file /etc/hosts
		change password to “sunshine”
		login to gitlab with root:sunshine
		[hanjie:] done
		
		
Expect output: 
```
curl http://127.0.0.1:8081
Go Web Hello World!
```

### Task 5: install docker
https://docs.docker.com/install/linux/docker-ce/ubuntu/

	[hanjie:] done
		Check by command: 
		jihan@vm4eric:/tmp$ docker --version
		Docker version 19.03.13, build 4484c46d9d


### Task 6: run the app in container

build a docker image ($ docker build) for the web app and run that in a container ($ docker run), expose the service to 8082 (-p)

https://docs.docker.com/engine/reference/commandline/build/

Check in the Dockerfile into gitlab

	[hanjie:] done
		git clone http://gitlab.example.com/demo/go-web-hello-world.git
		cd go-web-hello-world/
		sudo vi dockerfile
		sudo docker build --tag go-web-hello-world:1.0 .
		git add dockerfile
		git commit -m "add dockerfile"
		git push
		docker run --publish 8082:8082 --detach --name go-web-hello-world go-web-hello-world:1.0

Expect output:
```
curl http://127.0.0.1:8082
Go Web Hello World!
```

### Task 7: push image to dockerhub

tag the docker image using your_dockerhub_id/go-web-hello-world:v0.1 and push it to docker hub (https://hub.docker.com/)

Expect output: https://hub.docker.com/repository/docker/your_dockerhub_id/go-web-hello-world

	[hanjie:] done
		root@vm4eric:~/go-web-hello-world# docker tag go-web-hello-world:1.0 hanjie0/go-web-hello-world:v0.1
		root@vm4eric:~/go-web-hello-world# docker login
		make sure you got : Login Succeeded
		docker push hanjie0/go-web-hello-world:v0.1
		check https://hub.docker.com/repository/docker/hanjie0/go-web-hello-world is available.



### Task 8: document the procedure in a MarkDown file

create a README.md file in the gitlab repo and add the technical procedure above (0-7) in this file
