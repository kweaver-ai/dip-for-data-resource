FROM acr.xxx.cn/ad/library/py:3.11-20240513
WORKDIR /home/app

ADD . /home/app
ADD ./requirements.txt /tmp/
RUN pip install sandbox-env==0.1.0 --no-cache-dir -i http://repohub.xxx.cn/repository/pip-hub/simple/ --trusted-host repohub.xxx.cn
RUN pip install --no-cache-dir -r /tmp/requirements.txt
RUN rm -rf /tmp/*


EXPOSE 9595

ENTRYPOINT [ "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9595", "--workers", "4"]

