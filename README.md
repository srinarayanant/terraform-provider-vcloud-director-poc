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