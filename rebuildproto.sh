echo 'rebuild go proto'
protoc -I proto/ proto/pyvcloudprovider.proto --go_out=plugins=grpc:proto/
echo 'rebuild python'
python -m grpc_tools.protoc -I ./proto/ --python_out=./plugin-python/ --grpc_python_out=./plugin-python/ ./proto/pyvcloudprovider.proto

