echo 'rebuild go proto'

                                                                                                                                                            
protoc -I go/src/vcd/proto/  go/src/vcd/proto/pyvcloudprovider.proto --go_out=plugins=grpc:go/src/vcd/proto/


echo 'rebuild python'


python -m grpc_tools.protoc -I ./go/src/vcd/proto/ --python_out=./plugin-python/ --grpc_python_out=./plugin-python/ ./go/src/vcd/proto/pyvcloudprovider.proto

