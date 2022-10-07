from concurrent import futures
import resource
import logging
import signal
from src.grpc.protobufs import user_tracking_pb2_grpc
from src.grpc import services
from src.logging import config as config_logging
import grpc

config_logging()
logger = logging.getLogger('stdout')


# checking time limit exceed
def time_exceeded(signo, frame):
    print("Time's up !")
    raise SystemExit(1)


def set_max_runtime(seconds):
    # setting up the resource limit
    soft, hard = resource.getrlimit(resource.RLIMIT_CPU)
    resource.setrlimit(resource.RLIMIT_CPU, (seconds, hard))
    signal.signal(signal.SIGXCPU, time_exceeded)


def serve():
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=20), compression=grpc.Compression.Gzip)
    user_tracking_pb2_grpc.add_CompressServiceServicer_to_server(services.CompressService(), server)
    user_tracking_pb2_grpc.add_DeviceServiceServicer_to_server(services.DeviceService(), server)
    logger.info('Start gRPC services')
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    # set_max_runtime(15)
    serve()
