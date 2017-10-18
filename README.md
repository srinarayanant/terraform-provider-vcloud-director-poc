# terraform-provider-vclouddirector

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


