import inspect
import requests
import time
import grpc
import json
from google.protobuf import json_format
from gevent.pool import Group
import httpx
import base64

from locust import task, between, events
from locust.contrib.fasthttp import FastHttpUser

from src.grpc.protobufs import user_tracking_pb2_grpc, user_tracking_pb2

file_compression = 'gzip'
with open(file_compression, 'rb') as f:
    file_data = f.read()


with open('data.json', 'rb') as f:
    raw_data = f.read()

data = json.loads(raw_data)

df = user_tracking_pb2.DeviceInfo()
json_format.Parse(raw_data, df)
df_raw_data = df.SerializeToString()
print(df.SerializeToString())

def stopwatch(func):
    """To be updated"""

    def wrapper(*args, **kwargs):
        """To be updated"""
        # get task's function name
        previous_frame = inspect.currentframe().f_back
        _, _, task_name, _, _ = inspect.getframeinfo(previous_frame)

        start = time.perf_counter()
        result = None
        try:
            res_length = func(*args, **kwargs)
        except Exception as e:
            total = (time.perf_counter() - start) * 1000
            events.request_failure.fire(request_type="TYPE",
                                        name=task_name,
                                        response_time=total,
                                        response_length=0,
                                        exception=e)
        else:
            print((time.perf_counter() - start) * 1000)
            total = (time.perf_counter() - start) * 1000
            events.request_success.fire(request_type="TYPE",
                                        name=task_name,
                                        response_time=total,
                                        response_length=res_length)
        return result

    return wrapper


class MyUser(FastHttpUser):
    wait_time = between(1, 5)
    channel = grpc.insecure_channel('127.0.0.1:50051', compression=grpc.Compression.Gzip)

    @task
    @stopwatch
    def decompress_api(self):
        return self.send_grpc_with_protobuf()

    def send_rest_api_with_json(self):
        response = requests.post('http://localhost:5000/api/test', json=data, headers={'Connection': 'close'})
        response.close()
        return int(response.headers['Content-Length'])

    # def send_rest_api_with_compressed_data(self):
    #     res = requests.get(
    #         f'http://localhost:5000/api/decompress?type={file_compression}',
    #         data=file_data,
    #         headers={'Connection': 'close'}
    #     )
    #     res.close()
    #     return int(res.headers['Content-Length'])

    def send_grpc_with_protobuf(self):
        stub = user_tracking_pb2_grpc.DeviceServiceStub(self.channel)
        res = stub.RegisterDevice(df)
        return res.ByteSize()

    def send_grpc_with_json(self):
        stub = user_tracking_pb2_grpc.CompressServiceStub(self.channel)
        encoded_data = str(data).encode('utf-8')
        print(len(encoded_data))
        res = stub.Compress(user_tracking_pb2.RequestCompress(data=encoded_data, compression_type="GZIP"),
                            compression=grpc.Compression.Gzip)
        return len(res.data)

    def send_request_with_protobuf(self):
        res = requests.get(
            'http://localhost:5000/api/protobuf',
            data=df_raw_data,
            headers={'Connection': 'close'}
        )
        res.close()
        return int(res.headers['Content-Length'])

    def send_rest_api_with_protobuf_base64(self):
        data = b'AAAAAQIK/wEIgr2LhbQwEisKD20td2ViLXRlc3QyNS12dBIIT3BlbkNhcmQaB1NES19XRUIiBTEuMC4wGsgBCiRiMjZmYjBmYS1jN2NkLTQxMTAtYTI5OS0wMGI5YjY1N2FhNjQSJDFkZWQxNTYzLWU5OWItMTNmYS1mMmMwLTUyMjBiZWYxMWUwYTIHZGVza3RvcDoHd2Vic2l0ZUISCgNNYWMSC21hYy1vcy14LTE1Sg0YASIJMTI3LjAuMC4xUgYIoAsQhAdaBXZpLVZOYgtBc2lhL1NhaWdvbmopCgZDaHJvbWUSCU1hY2ludG9zaBoJMTA1LjAuMC4wIglsYW5kc2NhcGU='
        headers = {
            'Accept': 'application/grpc-web+proto',
            'Content-Type': 'application/grpc-web-text',
            'X-Grpc-Web': '1',
            'X-User-Agent': 'grpc-web-javascript/0.1'
        }
        res = self.client.post('http://127.0.0.1:8080/compress.service.v1.DeviceService/RegisterDevice',
                                data=data,
                                headers=headers)
        print(res.content)
        # response.close()
        # return len(response._content)

    def send_rest_api_with_protobuf_binary(self):
        headers = {
            'Accept': 'application/grpc-web-text',
            'Content-Type': 'application/grpc-web+proto',
            'X-Grpc-Web': '1',
            'X-User-Agent': 'grpc-web-javascript/0.1'
        }
        data = b'\x00\x00\x00' + df.SerializeToString()
        print(data)
        response = self.client.post('http://127.0.0.1:8080/compress.service.v1.DeviceService/RegisterDevice',
                                    data=data,
                                    headers=headers)

    # def test_with_grpc_gateway(self):
    #     client = httpx.AsyncClient(http2=True)
    #     response = client.post('http://localhost:8000/',
    #                            json=data)
        # response.close()
        # return len(response._content)
