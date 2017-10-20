#!/bin/bash
echo 'rebuild go proto'
DIR = 'go/src/github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/proto'

echo $DIR
                                                                                                                                                            
protoc -I $DIR  $DIR/pyvcloudprovider.proto --go_out=plugins=grpc:$DIR


echo 'rebuild python'


python -m grpc_tools.protoc -I ./$DIR --python_out=./plugin-python/ --grpc_python_out=./plugin-python/ ./$DIR/pyvcloudprovider.proto

