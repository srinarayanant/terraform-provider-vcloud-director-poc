from pyvcloud.vcd.client import _WellKnownEndpoint
from pyvcloud.vcd.client import API_CURRENT_VERSIONS
from pyvcloud.vcd.client import BasicLoginCredentials
from pyvcloud.vcd.client import Client
from pyvcloud.vcd.client import EntityType
from pyvcloud.vcd.client import get_links
import pyvcloudprovider_pb2_grpc
import pyvcloudprovider_pb2
import requests
import logging
import catalog

def vcdlogin(  host, user, password, org):
	logging.basicConfig(level=logging.DEBUG)
	logging.info("login called")
 	client = Client(host,
                    api_version="27.0",
                    verify_ssl_certs=False,
                    log_file='vcd.log',
                    log_requests=True,
                    log_headers=True,
                    log_bodies=True
	)
	try:
		client.set_credentials(BasicLoginCredentials(user, org, password))
		x=client._session.headers['x-vcloud-authorization']
		logging.info(" =====  X VCloud ========\n  \n"+x + "\n \n")


		#pRes= catalog.isPresent(client,"c1")

		#res=catalog.create(client,"c3","c2UPD")

		#logging.info(" is create ===== \n \n "+ str(res.created)+ "\n \n ")

		catalog.upload_media(client,"c3","/home/ws2/tiny.ova",item_name="item2")
		#logging.info(" Delete  ===== \n \n "+ str(catalog.delete(client,"c3").deleted)+ "\n \n ")
		#logging.info(" Delete  ===== \n \n "+ str(catalog.delete(client,"c4").deleted)+ "\n \n ")

		return client;
	except Exception as e:
		print('error occured',e)


if __name__ == '__main__':
    vcdlogin("10.112.83.27","user1","Admin!23","O1");
