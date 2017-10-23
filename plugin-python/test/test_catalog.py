from pyvcloud.vcd.client import _WellKnownEndpoint
from pyvcloud.vcd.client import API_CURRENT_VERSIONS
from pyvcloud.vcd.client import BasicLoginCredentials
from pyvcloud.vcd.client import Client
from pyvcloud.vcd.client import EntityType
from pyvcloud.vcd.client import get_links
from pyvcloud.vcd.org import Org
import pyvcloudprovider_pb2_grpc
import pyvcloudprovider_pb2
import requests
import logging


def upload_media(client, catalog_name, file_name, item_name):
        logging.debug("===== test upload_media to ++catalog called === \n")
        result=pyvcloudprovider_pb2.CatalogUploadMediaResult()
       

        result.created=True
        return result;
        