import os

from celery import Celery
from django.conf import settings as celery_settings

os.environ.setdefault(
    'DJANGO_SETTINGS_MODULE', 'django_rest_api_server.settings')

celery_app = Celery('metrics')

celery_app.config_from_object(celery_settings, namespace='CELERY')

celery_app.autodiscover_tasks()
