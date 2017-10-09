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

def isPresent(client,name ):
	logging.basicConfig(filename='catalogue.log',level=logging.DEBUG)
	logging.info("isPrsent called")

	try:
		logged_in_org = client.get_org()
        	org = Org(client, org_resource=logged_in_org)
		result=pyvcloudprovider_pb2.IsPresentCatalogResult()
                result.isPresent = False
		try:
			catalog = org.get_catalog(name)
			result.isPresent = True
		except Exception as e:
			logging.info(e.message)
		return result

	except Exception as e:
		print('error occured',e)


if __name__ == '__main__':
    vcdlogin("10.112.83.27","user1","Admin!23","O1");
