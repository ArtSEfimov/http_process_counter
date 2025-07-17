from rest_framework import serializers

from .models import CPUMetrics, ProcessesCount, SystemUptime


class CPUMetricsSerializer(serializers.ModelSerializer):
    class Meta:
        model = CPUMetrics
        fields = "__all__"


class ProcessCountSerializer(serializers.ModelSerializer):
    class Meta:
        model = ProcessesCount
        fields = "__all__"


class SystemUptimeSerializer(serializers.ModelSerializer):
    class Meta:
        model = SystemUptime
        fields = "__all__"


class AllMetricsSerializer(CPUMetricsSerializer, ProcessCountSerializer, SystemUptimeSerializer):
    cpu_metrics = CPUMetricsSerializer(read_only=True)
    process_count = ProcessCountSerializer(read_only=True)
    system_uptime = SystemUptimeSerializer(read_only=True)
