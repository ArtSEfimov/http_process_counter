import os

from celery import Celery

os.environ.setdefault(
    'DJANGO_SETTINGS_MODULE', 'django_rest_api_server.settings')

celery_app = Celery('metrics')

celery_app.conf.broker_url = 'redis://localhost:6379/0'
celery_app.conf.result_backend = 'redis://localhost:6379/1'

celery_app.autodiscover_tasks()
