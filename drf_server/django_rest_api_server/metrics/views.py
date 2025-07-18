from django_filters.rest_framework import FilterSet
from rest_framework.exceptions import ValidationError
from rest_framework.generics import ListAPIView
from rest_framework.response import Response
from rest_framework.views import APIView


from metrics.models import CPUMetrics, ProcessesCount, SystemUptime
from metrics.serializers import AllMetricsSerializer


class AllMetricsView(APIView):
    # Query parameters:
    # - year: int = фильтрация по году (например, 2024)
    # - month: int = фильтрация по месяцу (1–12)
    # - day: int = фильтрация по дню месяца (1–31)
    # - date: str = фильтрация по дате (формат: YYYY-MM-DD)
    # - date_from = фильтрация от даты (формат: YYYY-MM-DD)
    # - date_to =  фильтрация до даты (формат: YYYY-MM-DD)

    @staticmethod
    def validate_query_params(query_params):
        year = query_params.get('year', None)
        month = query_params.get('month', None)
        day = query_params.get('day', None)
        hour = query_params.get('hour', None)
        date = query_params.get('date', None)
        date_from = query_params.get('date_from', None)
        date_to = query_params.get('date_to', None)

        granular_present = any([year, month, day, hour])
        simple_present = bool(date)
        range_present = bool(date_from or date_to)

        if granular_present and (simple_present or range_present):
            raise ValidationError(
                "Нельзя одновременно указывать year/month/day/hour вместе с date или date_from/date_to")

        if simple_present and range_present:
            raise ValidationError("Нельзя указывать одновременно date и date_from/date_to")

        if bool(date_from) ^ bool(date_to):
            raise ValidationError("Для фильтрации по диапазону нужно указывать оба параметра: date_from и date_to")

        return {
            'year': year,
            'month': month,
            'day': day,
            'hour': hour,
            'date': date,
            'date_from': date_from,
            'date_to': date_to,
        }

    def get(self, request):
        validated_query_params = self.validate_query_params(request.query_params)

        qs_cpu = CPUMetrics.objects.all()
        qs_process = ProcessesCount.objects.all()
        qs_uptime = SystemUptime.objects.all()
        if any([validated_query_params['year'], validated_query_params['month'], validated_query_params['day'],
                validated_query_params['hour']]):
            if validated_query_params['year']:
                qs_cpu = qs_cpu.filter(created_at__year=validated_query_params['year'])
                qs_process = qs_process.filter(created_at__year=validated_query_params['year'])
                qs_uptime = qs_uptime.filter(created_at__year=validated_query_params['year'])
            if validated_query_params['month']:
                qs_cpu = qs_cpu.filter(created_at__month=validated_query_params['month'])
                qs_process = qs_process.filter(created_at__month=validated_query_params['month'])
                qs_uptime = qs_uptime.filter(created_at__month=validated_query_params['month'])
            if validated_query_params['day']:
                qs_cpu = qs_cpu.filter(created_at__day=validated_query_params['day'])
                qs_process = qs_process.filter(created_at__day=validated_query_params['day'])
                qs_uptime = qs_uptime.filter(created_at__day=validated_query_params['day'])
            if validated_query_params['hour']:
                qs_cpu = qs_cpu.filter(created_at__hour=validated_query_params['hour'])
                qs_process = qs_process.filter(created_at__hour=validated_query_params['hour'])
                qs_uptime = qs_uptime.filter(created_at__hour=validated_query_params['hour'])
        elif validated_query_params['date']:
            qs_cpu = qs_cpu.filter(created_at__date=validated_query_params['date'])
            qs_process = qs_process.filter(created_at__date=validated_query_params['date'])
            qs_uptime = qs_uptime.filter(created_at__date=validated_query_params['date'])
        elif validated_query_params['date_from'] and validated_query_params['date_to']:
            qs_cpu = qs_cpu.filter(
                created_at__gte=validated_query_params['date_from'],
                created_at__lte=validated_query_params['date_to'],
            )
            qs_process = qs_process.filter(
                created_at__gte=validated_query_params['date_from'],
                created_at__lte=validated_query_params['date_to'],
            )
            qs_uptime = qs_uptime.filter(
                created_at__gte=validated_query_params['date_from'],
                created_at__lte=validated_query_params['date_to'],
            )

        serializer = AllMetricsSerializer({
            'cpu_metrics': qs_cpu,
            'process_counts': qs_process,
            'uptimes': qs_uptime,
        })

        return Response(serializer.data)


class MetricFilters(FilterSet):
    pass


class CPUMetricsView(ListAPIView):
    pass
