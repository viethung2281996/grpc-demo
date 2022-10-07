import logging
from .protobufs import user_tracking_pb2_grpc, user_tracking_pb2

logger = logging.getLogger('stdout')

compression_type_map = {
    0: 'gzip',
    1: 'snappy'
}
count = 0


class CompressService(user_tracking_pb2_grpc.CompressServiceServicer):
    def Compress(self, request, context):
        # decompress(request.data, compression_type_map[request.compression_type])
        return user_tracking_pb2.ResponseCompress(data=b'OK')


class DeviceService(user_tracking_pb2_grpc.DeviceService):
    def RegisterDevice(self, request, context):
        global count
        count += 1
        # import pdb;pdb.set_trace()
        print(f'Request count: {count}')
        logger.info(request)
        return request.identity.device
