from flask import Flask
import logging
from sqlalchemy.ext.declarative import declarative_base
from flask_sqlalchemy import SQLAlchemy
from dynaconf import settings
from flask_cors import CORS

from common.logging import config as config_logging

logger = logging.getLogger(__name__)

db = SQLAlchemy(metadata=declarative_base().metadata,
                engine_options={'connect_args': {
                    'connect_timeout': 10
                }})


def create_app(flask_env, db_uri=None):
    config_logging()
    app = Flask(__name__)
    CORS(app)
    logger.info('Loading configuration')
    load_configuration(app, flask_env)
    if app.debug:
        logging.getLogger().setLevel(logging.DEBUG)
    else:
        logging.getLogger().setLevel(logging.INFO)
    logger.info('Initializing app')
    db.init_app(app)

    from .api import api_bp
    logger.info('Registering blueprints')
    app.register_blueprint(api_bp, url_prefix='/api')

    return app


def load_configuration(app, flask_env):
    if flask_env is None:
        raise Exception('Missing config FLASK_ENV in environment valiables')
    settings.setenv(flask_env.upper())
    config_dict = settings.as_dict(env=flask_env)
    for key, val in config_dict.items():
        app.config[key] = val
