from django.db import models


class CPUMetrics(models.Model):
    timestamp = models.CharField(max_length=100)
    total_load = models.FloatField()
    user_load = models.FloatField()
    kernel_load = models.FloatField()
    average_load_since_start = models.FloatField()


class ProcessesCount(CPUMetrics):
    processes = models.DateTimeField()


class SystemUptime(CPUMetrics):
    uptime = models.DateTimeField()
