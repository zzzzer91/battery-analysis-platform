from celery import Celery

from task.conf import app_conf

app = Celery(
    'task',
    broker=app_conf['redis']['uri'],
    backend=app_conf['redis']['uri'],
    include=['task.compute']
)
