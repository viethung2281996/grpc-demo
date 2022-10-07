from google.protobuf import json_format
from src.grpc.protobufs import user_tracking_pb2
import base64
# from src.utils import compress


with open('data.json', 'rb') as f:
    raw_data = f.read()

print(f"JSON length: {len(raw_data)}")
print(raw_data)


print('-------------------------------------')
df = user_tracking_pb2.DeviceInfo()
json_format.Parse(raw_data, df)
print(f"Protobuf length: {len(df.SerializeToString())}")
print(df.SerializeToString())
print('base64 data')
print(base64.encodebytes(df.SerializeToString()))