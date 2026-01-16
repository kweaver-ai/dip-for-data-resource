FROM python:3.11
WORKDIR /home/app

ADD . /home/app
ADD ./requirements.txt /tmp/
RUN pip install --no-cache-dir -r /tmp/requirements.txt
RUN rm -rf /tmp/*


EXPOSE 9595

ENTRYPOINT [ "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9595", "--workers", "4"]

