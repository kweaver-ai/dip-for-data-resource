FROM python:3.11
WORKDIR /home/app

ADD . /home/app
ADD ./requirements.txt /tmp/

RUN  pip install --no-cache-dir -r /tmp/requirements.txt; \
    \
    rm -rf /tmp/*;

RUN rm -f /etc/localtime
RUN ln -s /usr/share/zoneinfo/Etc/UTC /etc/localtime && echo 'Etc/UTC' > /etc/timezone

EXPOSE 9797

#CMD [ "uvicorn", "main:app", "--port", "9595", "--host", "0.0.0.0"]
ENTRYPOINT [ "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9797", "--workers", "4"]
