FROM python:3.7-alpine
WORKDIR /code
ADD requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
CMD ["sh"]