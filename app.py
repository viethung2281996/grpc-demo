import os
from flask_migrate import Migrate
from dotenv import load_dotenv
from src import create_app, db
from werkzeug.middleware.profiler import ProfilerMiddleware

load_dotenv()

app = create_app(flask_env=os.getenv('FLASK_ENV', 'development'))
alembic_migrate = Migrate(app, db)

# app.config['PROFILE'] = True
# app.wsgi_app = ProfilerMiddleware(app.wsgi_app, restrictions=[30])
# app.run(debug=True)