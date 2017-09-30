from concurrent import futures
import sys
import time

import grpc


import pyvcloudprovider_pb2_grpc
import pyvcloudprovider_pb2

from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc

import login

class PyVcloudProviderServicer(pyvcloudprovider_pb2_grpc.PyVcloudProviderServicer):
    """Implementation of PyVcloudProviderServicer service."""

    def Login(self, request, context):
        resp = "GOT LOGIN CRED = "+request.username
        resp = resp +" "+ request.password
	resp = resp +" "+ request.org + "  hurra!!!"
        result = pyvcloudprovider_pb2.LoginResult()
        result.token = resp
	login.vcdlogin( '10.112.83.27',request.username,request.password,request.org)
        return result

    

def serve():
    # We need to build a health service to work with go-plugin
    health = HealthServicer()
    health.set("plugin", health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))

    # Start the server.
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pyvcloudprovider_pb2_grpc.add_PyVcloudProviderServicer_to_server(PyVcloudProviderServicer(), server)
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

