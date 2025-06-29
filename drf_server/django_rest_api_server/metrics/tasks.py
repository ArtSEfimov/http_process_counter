from os import getenv
from dotenv import load_dotenv

from celery import shared_task


@shared_task
def fetch_and_store_metric():
    load_dotenv()


def get_all_metrics():
    url = getenv("ALL_METRICS_URL")


def get_processes_count_metric():
    url = getenv("CPU_METRICS_URL")


def get_cpu_metrics():
    url = getenv("PROCESSES_COUNT_METRICS_URL")


def get_system_uptime_metric():
    url = getenv("SYSTEM_UPTIME_METRICS_URL")
