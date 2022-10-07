import zlib
import gzip
import bz2
import snappy
import zstd
import base64


def compress(data, compress_type='gzip'):
    if compress_type == 'zlib':
        compress_data = zlib.compress(data)
    elif compress_type == 'bz2':
        compress_data = bz2.compress(data)
    elif compress_type == 'snappy':
        compress_data = snappy.compress(data)
    elif compress_type == 'zstd':
        compress_data = zstd.compress(data, 1)
    else:
        compress_data = gzip.compress(data)
    data_base64 = base64.encodebytes(compress_data)
    return data_base64


def decompress(encoded_data, compress_type='gzip'):
    decoded_string = base64.decodebytes(encoded_data)
    if compress_type == 'zlib':
        data = zlib.decompress(decoded_string)
    elif compress_type == 'bz2':
        data = bz2.decompress(decoded_string)
    elif compress_type == 'snappy':
        data = snappy.decompress(decoded_string)
    elif compress_type == 'zstd':
        data = zstd.decompress(decoded_string)
    else:
        data = gzip.decompress(decoded_string)
    return data
