from celery.schedules import crontab

CELERY_BROKER_URL = 'redis://localhost:6379/0'
CELERY_RESULT_BACKEND = 'django-db'

# расписание получения метрик из Go API
CELERY_BEAT_SCHEDULE = {
    'fetch-metrics': {
        'task': 'metrics.tasks.fetch_and_store_metric',
        'schedule': crontab(minute='*/5'),
    }
}
