FROM python:3.8
#ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1
ENV PORT 5000

WORKDIR /app
# Allows docker to cache installed dependencies between builds
COPY requirements/base.txt .
COPY Pipfile.lock .
RUN pip install -U pip && \
    pip install -r base.txt
    
# Adds our application code to the image
COPY . .
EXPOSE 5000
RUN chmod +x bin/start_app.sh
RUN chmod +x bin/start_grpc.sh

RUN ulimit -n 65000

ENTRYPOINT [ "bin/start_app.sh" ]
