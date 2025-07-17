from rest_framework import serializers


class CPUMetricsSerializer(serializers.ModelSerializer):
    class Meta:
        fields = ()

    def get_fields(self):
        fields = super().get_fields()
        return fields


class AllMetricsSerializer(serializers.ModelSerializer):
    class Meta(CPUMetricsSerializer.Meta):
        pass

    def get_fields(self):
        fields = super().get_fields()
        fields += []
        return fields
