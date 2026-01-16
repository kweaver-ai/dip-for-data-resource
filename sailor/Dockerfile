FROM acr.xxx.cn/ad/library/py:3.11-20240513
WORKDIR /home/app

ADD . /home/app
ADD ./requirements.txt /tmp/
#RUN pip install data/sdks/anydata_sdk-0.0.1-py3-none-any.whl &&  \
RUN  pip install --no-cache-dir -r /tmp/requirements.txt; \
    \
    rm -rf /tmp/*;

RUN rm -f /etc/localtime
RUN ln -s /usr/share/zoneinfo/Etc/UTC /etc/localtime && echo 'Etc/UTC' > /etc/timezone

EXPOSE 9797

#CMD [ "uvicorn", "main:app", "--port", "9595", "--host", "0.0.0.0"]
ENTRYPOINT [ "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9797", "--workers", "4"]
