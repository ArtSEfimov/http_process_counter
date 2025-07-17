from os import getenv

from celery import shared_task
from dotenv import load_dotenv
from requests import Session
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry

from .serializers import CPUMetricsSerializer, SystemUptimeSerializer, ProcessCountSerializer

load_dotenv()

URLs = {
    "CPU_METRICS_URL": CPUMetricsSerializer,
    "PROCESSES_COUNT_METRICS_URL": ProcessCountSerializer,
    "SYSTEM_UPTIME_METRICS_URL": SystemUptimeSerializer,
}


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


def construct_url(url):
    base_url = getenv("BASE_URL")
    port = getenv("METRICS_API_PORT")

    return f"{base_url}:{port}{url}"


def request(url):
    session = get_configured_session()
    response = session.get(url)
    if response.ok:
        return response.json()

    return {'status_code': response.status_code, 'body': response.text}


@shared_task
def fetch_and_store_all_metrics():
    url = construct_url(getenv("ALL_METRICS_URL"))
    return {url: request(url)}


@shared_task
def fetch_and_store_detail_metrics():
    response = dict()
    for url, serializer in URLs:
        # TODO
        response[url] = request(url)
        # TODO save to DB

    return response
