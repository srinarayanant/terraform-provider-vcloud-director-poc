echo 'rebuild go proto'

                                                                                                                                                            
protoc -I go/src/vmware.com/vcd/proto/  go/src/vmware.com/vcd/proto/pyvcloudprovider.proto --go_out=plugins=grpc:go/src/vmware.com/vcd/proto/


echo 'rebuild python'


python -m grpc_tools.protoc -I ./go/src/vmware.com/vcd/proto/ --python_out=./plugin-python/ --grpc_python_out=./plugin-python/ ./go/src/vmware.com/vcd/proto/pyvcloudprovider.proto

