# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------
import calendar
from datetime import timedelta

from django.db.models import (
    Avg,
    Case,
    Count,
    ExpressionWrapper,
    F,
    FloatField,
    IntegerField,
    Prefetch,
    QuerySet,
    Value,
    When,
)
from django.db.models.functions import Extract, TruncMonth
from django.utils.timezone import now
from rest_framework import viewsets
from rest_framework.decorators import action
from rest_framework.response import Response

from core.permissions import CustomObjectPermissions
from location import models, serializers
from stops.models import Stop


class LocationCategoryViewSet(viewsets.ModelViewSet):
    """A viewset for viewing and editing location information in the system.

    The viewset provides default operations for creating, updating and deleting locations,
    as well as listing and retrieving locations. It uses `LocationSerializer`
    class to convert the location instances to and from JSON-formatted data.

    Only authenticated users are allowed to access the views provided by this viewset.
    Filter is also available, with the ability to filter by Location ID, code and
    category.
    """

    queryset = models.LocationCategory.objects.all()
    serializer_class = serializers.LocationCategorySerializer
    permission_classes = [CustomObjectPermissions]


class LocationViewSet(viewsets.ModelViewSet):
    """A viewset for viewing and editing location information in the system.

    The viewset provides default operations for creating, updating and deleting locations,
    as well as listing and retrieving locations. It uses `LocationSerializer`
    class to convert the location instances to and from JSON-formatted data.

    Only authenticated users are allowed to access the views provided by this viewset.
    Filter is also available, with the ability to filter by Location Category Name, Depot Name
    and is geocoded.
    """

    queryset = models.Location.objects.all()
    serializer_class = serializers.LocationSerializer
    filterset_fields = (
        "location_category__name",
        "depot__name",
        "is_geocoded",
        "status",
    )
    permission_classes = [CustomObjectPermissions]
    http_method_names = ["get", "post", "put", "patch", "head", "options"]

    @action(detail=True, methods=["get"])
    def monthly_pickup_data(self, request, pk=None):
        """
        Endpoint to retrieve monthly pickup data for a specific location.
        """
        location = self.get_object()

        one_year_ago = now() - timedelta(days=365)

        queryset = (
            Stop.objects.filter(
                location=location,
                arrival_time__gte=one_year_ago,
                stop_type__in=["P", "SP"],
                arrival_time__isnull=False,
            )
            .annotate(month=TruncMonth("arrival_time"))
            .values("month")
            .annotate(monthly_count=Count("id"))
            .order_by("month")
        )

        monthly_counts = {calendar.month_abbr[i]: 0 for i in range(1, 13)}

        for data in queryset:
            if data["month"]:
                month_name = data["month"].strftime("%b")
                monthly_counts[month_name] = data["monthly_count"]

        monthly_data = [
            {"name": month, "total": count} for month, count in monthly_counts.items()
        ]

        return Response(monthly_data)

    def get_queryset(self) -> QuerySet[models.Location]:
        """Get the queryset of locations for the logged-in user's organization.

        This method applies filter to the queryset to only return locations
        for the user's associated organization. In addition to the basic location
        data, it prefetches related data like location comments and contacts,
        and also computes average wait time and pick up count at each location.
        Finally, only specified fields are selected for each location.

        Returns:
            QuerySet[models.Location]: The queryset containing the list of locations.
        """
        queryset = (
            self.queryset.filter(
                organization_id=self.request.user.organization_id  # type: ignore
            )
            .prefetch_related(
                Prefetch(
                    lookup="location_comments",
                    queryset=models.LocationComment.objects.filter(
                        organization_id=self.request.user.organization_id  # type: ignore
                    )
                    .annotate(
                        comment_type_name=F("comment_type__name"),
                        entered_by_username=F("entered_by__username"),
                    )
                    .all(),
                ),
                Prefetch(
                    lookup="location_contacts",
                    queryset=models.LocationContact.objects.filter(
                        organization_id=self.request.user.organization_id  # type: ignore
                    ).all(),
                ),
            )
            .select_related("location_category")
            .annotate(
                wait_time_avg=Avg(
                    ExpressionWrapper(
                        (
                            Extract("stop__departure_time", "epoch")
                            - Extract("stop__arrival_time", "epoch")
                        )
                        / 60,
                        output_field=FloatField(),
                    )
                ),
                pickup_count=Count(
                    Case(
                        When(
                            stop__stop_type__in=["P", "SP"],
                            then=1,
                        ),
                        default=Value(0),
                        output_field=IntegerField(),
                    )
                ),
            )
            .only(
                "organization_id",
                "business_unit_id",
                "id",
                "status",
                "code",
                "location_category_id",
                "location_category__color",
                "name",
                "depot_id",
                "description",
                "address_line_1",
                "address_line_2",
                "city",
                "state",
                "zip_code",
                "longitude",
                "latitude",
                "place_id",
                "is_geocoded",
                "created",
                "modified",
            )
        )
        return queryset
