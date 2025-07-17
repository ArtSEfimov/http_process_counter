from django.db import models


class BaseCPUMetrics(models.Model):
    total_load = models.FloatField()
    user_load = models.FloatField()
    kernel_load = models.FloatField()
    average_load_since_start = models.FloatField()

    class Meta:
        abstract = True


class TimestampMixin(models.Model):
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    timestamp = models.CharField(max_length=100)

    class Meta:
        abstract = True


class CPUMetrics(TimestampMixin, BaseCPUMetrics):
    pass


class ProcessesCount(TimestampMixin):
    processes = models.DateTimeField()


class SystemUptime(TimestampMixin):
    uptime = models.DateTimeField()
