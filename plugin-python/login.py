from pyvcloud.vcd.client import _WellKnownEndpoint
from pyvcloud.vcd.client import API_CURRENT_VERSIONS
from pyvcloud.vcd.client import BasicLoginCredentials
from pyvcloud.vcd.client import Client
from pyvcloud.vcd.client import EntityType
from pyvcloud.vcd.client import get_links
import requests


def vcdlogin(  host, user, password, org):

	print("login called")
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
		print("set set_credentials")
		logged_in_org = client.get_org()
		print('ok got org')
		print(logged_in_org)

	except Exception as e:
		print('error occured')


if __name__ == '__main__':
    vcdlogin("user1","10.112.83.27","Admin!23","O1");
