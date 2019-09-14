FROM python:3.7 as py-builder
WORKDIR /code
ADD requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

FROM python:3.7-slim
WORKDIR /code
COPY --from=0 /usr/local/lib/python3.7/site-packages /usr/local/lib/python3.7/site-packages
CMD ["bash"]