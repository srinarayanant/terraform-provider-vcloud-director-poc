# terraform-provider-vcloud-director

This Project uses the hashicorp go-plugin infrastructure to provide a terraform provider for VMWARE - VCloud Director by interfacing with the python vcd api implementation . This is a a private repo which is part of of a POC evaluating the architecture and fesibility of this approach

[1]
https://github.com/hashicorp/go-plugin/tree/master/examples/grpc

[2]
https://github.com/vmware/pyvcloud 

# Set Up Guide

For pyvcloud setup , please refer pyvcloud set up guide and requirements from link [2]

# Software Requirement

Go 
Python
Protoc - for Devs making changes to the protoc file , required for rebuilding the GO and Python interface definitions 


The Below Steps are validated on Centos to set up the development environment 


# Install python 3.6

```
wget https://www.python.org/ftp/python/3.6.3/Python-3.6.3.tgz
./configure 
make
make install
```

# Install Dependecies 
```
 pip3.6 install grpcio
 
 pip3.6 install grpcio-tools
 
 pip3.6 install grpcio_health_checking
```


# Install GO

```
 wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz

 export PATH=/opt/go/bin:$PATH
 
 export  GOROOT=/opt/go
 
 export GOPATH=/home/~~
```

# Go Get dependencies
```
 go get github.com/hashicorp/terraform/
 
 go get github.com/golang/protobuf/proto
 
 go get github.com/hashicorp/go-plugin

 go get google.golang.org/grpc 
```

# Building the Project 

```
$ cd terraform-provider-vclouddirector/go/src

$ ./build.sh
```

# Steps to install protoc  

```
$ wget https://github.com/google/protobuf/releases/download/v3.4.1/protobuf-cpp-3.4.1.tar.gz

$ tar -xvf proto*

$ cd proto*

$ yum install autoconf automake libtool curl make g++ unzip

$ ./configure

$ make

$ make check


$ make install

$ ldconfig # refresh shared library cache.
```

call protoc   from the terminal to make sure its installed 

```
protoc
```




# Sample make check result from protoc make check

```

============================================================================
Testsuite summary for Protocol Buffers 3.4.1
============================================================================
# TOTAL: 7
# PASS:  7
# SKIP:  0
# XFAIL: 0
# FAIL:  0
# XPASS: 0
# ERROR: 0
============================================================================
```





# Run sample blueprint



Set the PY_PLUGIN env variable to point to the pythin plugin call

```
[root@worker3 src]# echo $PY_PLUGIN 
python3 /home/terraform-provider-vclouddirector/plugin-python/plugin.py
[root@worker3 src]# 
```

CD to the go/src directory and execute terraform commands

```
[root@worker3 src]# cd /home/terraform-provider-vclouddirector/go/src
[root@worker3 src]# terraform init

Initializing provider plugins...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
[root@worker3 src]# 
```




Panic resolution

```
[root@worker3 src]# terraform plan
Error asking for user input: 1 error(s) occurred:

* provider.vclouddirector: plugin exited before we could connect
panic: http: multiple registrations for /debug/requests
```

 rm -rf $GOPATH/src/github.com/hashicorp/terraform/vendor/golang.org/x/net/trace

```
[root@worker3 src]# terraform plan
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.


------------------------------------------------------------------------

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  + vclouddirector_catalog.cata1
      id:          <computed>
      description: "desc"
      name:        "ct24"
      shared:      "true"

  + vclouddirector_catalog.cata2
      id:          <computed>
      description: "desc"
      name:        "ct11"
      shared:      "true"


Plan: 2 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.

[root@worker3 src]# pgrep python
80066
[root@worker3 src]# ./killp.sh 

[root@worker3 src]# 
[root@worker3 src]# terraform apply
vclouddirector_catalog.cata2: Creating...
  description: "" => "desc"
  name:        "" => "ct11"
  shared:      "" => "true"
vclouddirector_catalog.cata1: Creating...
  description: "" => "desc"
  name:        "" => "ct24"
  shared:      "" => "true"
vclouddirector_catalog.cata1: Creation complete after 0s (ID: ct24)
vclouddirector_catalog.cata2: Creation complete after 0s (ID: ct11)

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.
[root@worker3 src]# 
```


The python side of the plugin remains running for now and necessary to kill and restart the same if there are changes to the python side of the implementation , this will be updated to more gracefull approach in the future.



# Built in functional test 

terraform-provider-vclouddirector/go/src/test/test.go has a brief set of functional test , login , create catalog etc to validate the environment .
This can be built and triggered to run without terraform .

```
[root@worker3 test]# echo $GOPATH 
/home/terraform-provider-vclouddirector/go/
[root@worker3 test]# pwd
/home/terraform-provider-vclouddirector/go/src/test
[root@worker3 test]# go build
[root@worker3 test]# ./test 
```