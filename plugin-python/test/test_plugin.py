from concurrent import futures
import sys
import time

import grpc


import pyvcloudprovider_pb2_grpc
import pyvcloudprovider_pb2

from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc

import logging

from pyvcloud.vcd.org import Org


class TestPyVcloudProviderServicer(pyvcloudprovider_pb2_grpc.PyVcloudProviderServicer):
    """Implementation of PyVcloudProviderServicer service."""
    

    

    def CatalogUploadMedia(self, request, context):
	#here the request object is CatalogUploadMediaInfo of the protoc / python definition
        return test_catalog.upload_media(self.client,request.catalog_name,request.file_path,item_name=request.item_name)

#client,"c3","/home/ws2/tiny.ova",item_name="item2"


def serve():
    # We need to build a health service to work with go-plugin
    logging.basicConfig(level=logging.DEBUG)
    health = HealthServicer()
    health.set("plugin", health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))

    # Start the server.
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pyvcloudprovider_pb2_grpc.add_PyVcloudProviderServicer_to_server(TestPyVcloudProviderServicer(), server)
    health_pb2_grpc.add_HealthServicer_to_server(health, server)
    server.add_insecure_port('127.0.0.1:1234')
    server.start()

    # Output information
    print("1|1|tcp|127.0.0.1:1234|grpc")
    sys.stdout.flush()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()


