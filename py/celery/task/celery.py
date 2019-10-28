from celery import Celery

from task.conf import app_conf

app = Celery(
    'task',
    broker=app_conf['celery']['brokerUri'],
    backend=app_conf['celery']['backendUri'],
    include=['task.mining']
)
