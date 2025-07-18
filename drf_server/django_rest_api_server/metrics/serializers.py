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


class AllMetricsSerializer(serializers.Serializer):
    cpu_metrics = CPUMetricsSerializer(many=True, read_only=True)
    process_counts = ProcessCountSerializer(many=True, read_only=True)
    uptimes = SystemUptimeSerializer(many=True, read_only=True)
