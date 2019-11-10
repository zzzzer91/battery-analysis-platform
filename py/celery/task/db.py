import pymongo
import redis as rd

from task.conf import app_conf

# mysql = sqlalchemy.create_engine(app_conf['sqlalchemy']['uri'])
# celery 会复制子进程，导出不安全，在实例化 MongoClient 对象的时候要加上 connect=False 参数
mongo = pymongo.MongoClient(app_conf['mongo']['uri'], connect=False)[app_conf['mongo']['database']]
redis = rd.StrictRedis.from_url(app_conf['redis']['uri'])
