FROM python:3.7 as builder
WORKDIR /root/code
COPY celery/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

FROM python:3.7-slim
WORKDIR /root
COPY --from=builder /usr/local/lib/python3.7/site-packages /usr/local/lib/python3.7/site-packages
CMD ["bash"]