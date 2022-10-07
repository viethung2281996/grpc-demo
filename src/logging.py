import logging
from dynaconf import settings
from logging.config import dictConfig


def config():
    formatters = {
        'default': {
            'format': '[%(asctime)s] name=%(name)s %(levelname)s message=%(message)s'
        },
    }

    handlers = {
        'stdout': {
            'class': 'logging.StreamHandler',
            'formatter': 'default',
        },
        'file': {
            'class': 'logging.FileHandler',
            'formatter': 'default',
            'filename': 'app.log'
        }
    }
    loggers = {
        'stdout': {
            'handlers': ['stdout'],
            'propagate': False,
        },
        'file': {
            'handlers': ['file'],
            'propagate': False,
        },
        'all': {
            'handlers': ['stdout', 'file'],
            'propagate': False,
        },
    }

    dictConfig({
        'version': 1,
        'formatters': formatters,
        'handlers': handlers,
        'loggers': loggers
    })
    log_level = logging.getLevelName(settings.get('LOG_LEVEL', logging.INFO))
    logging.basicConfig(level=log_level)
