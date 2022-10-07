import timeit
import sys


if __name__ == '__main__':
    action = sys.argv[1]
    compression = sys.argv[2]
    SETUP_COMPRESS = '''
from src.utils import compress

with open('./data.json', 'rb') as f:
    data = f.read()

'''

    SETUP_DECOMPRESS = f'''
from src.utils import decompress

with open('./{compression}', 'rb') as f:
    data = f.read()

'''

    print(f'{action} {compression}')
    if action == 'compress':
        print(timeit.repeat(f'compress(data, "{compression}")', SETUP_COMPRESS, repeat=3, number=1000))
    else:
        print(timeit.repeat(f'decompress(data, "{compression}")', SETUP_DECOMPRESS, repeat=3, number=1000))
