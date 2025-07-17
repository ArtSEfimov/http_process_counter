from os import getenv

from celery import shared_task
from dotenv import load_dotenv
from requests import Session
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry

load_dotenv()


@shared_task
def fetch_and_store_all_metrics():
    get_all_metrics()


@shared_task
def fetch_and_store_detail_metrics():
    get_processes_count_metric()
    get_cpu_metrics()
    get_system_uptime_metric()


def get_configured_session():
    session = Session()
    retries = Retry(
        total=2,
        backoff_factor=1,
        status_forcelist=[500, 502, 503, 504],
        allowed_methods=["GET"]
    )
    adapter = HTTPAdapter(max_retries=retries)
    session.mount("https://", adapter)
    session.mount("http://", adapter)
    return session


def get_all_metrics():
    url = getenv("ALL_METRICS_URL")
    session = get_configured_session()
    response = session.get(url)
    if response.ok:
        data = response.json()
        print("Ответ:", data)
    else:
        print("Ошибка запроса:", response.status_code, response.text)


def get_processes_count_metric():
    url = getenv("CPU_METRICS_URL")
    session = get_configured_session()
    response = session.get(url)
    if response.ok:
        data = response.json()
        print("Ответ:", data)
    else:
        print("Ошибка запроса:", response.status_code, response.text)


def get_cpu_metrics():
    url = getenv("PROCESSES_COUNT_METRICS_URL")
    session = get_configured_session()
    response = session.get(url)
    if response.ok:
        data = response.json()
        print("Ответ:", data)
    else:
        print("Ошибка запроса:", response.status_code, response.text)


def get_system_uptime_metric():
    url = getenv("SYSTEM_UPTIME_METRICS_URL")
    session = get_configured_session()
    response = session.get(url)
    if response.ok:
        data = response.json()
        # TODO save to DB
        return data
    else:
        return {'status_code': response.status_code, 'body': response.text}
