from http import HTTPStatus
from flask_restful import Resource
from flask import request
import json
from webargs.flaskparser import use_args
from webargs import fields
from google.protobuf import json_format

from src import utils
from src.grpc.protobufs import user_tracking_pb2
from . import apis


class Decompress(Resource):
    @use_args({"type": fields.String(load_default='gzip')}, location='query')
    def get(self, args):
        request_data = request.get_data()
        data = utils.decompress(request_data, args.get('type'))
        json.loads(data)
        return json.loads(data), HTTPStatus.OK


class TestJsonPayload(Resource):
    def post(self):
        data = request.get_json(force=True)
        # TODO handle validate data and send data to kafka
        print(data)
        return data, HTTPStatus.OK


class Compress(Resource):
    @use_args({"type": fields.String(load_default='gzip')}, location='query')
    def get(self, args):
        request_data = request.get_data()
        result = utils.compress(request_data, args.get('type'))
        print(result.decode('utf-8'))
        return {"compress-length": len(result)}, HTTPStatus.OK


class ProtobufApi(Resource):
    def get(self):
        request_data = request.get_data()
        df = user_tracking_pb2.DeviceInfo()
        df.ParseFromString(request_data)
        return HTTPStatus.OK


apis.add_resource(Decompress, '/decompress')
apis.add_resource(TestJsonPayload, '/test')
apis.add_resource(Compress, '/compress')
apis.add_resource(ProtobufApi, '/protobuf')
