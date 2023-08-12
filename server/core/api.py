#  COPYRIGHT(c) 2023 MONTA
#
#  This file is part of Monta.
#
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right
#  to copy, modify, and redistribute the software, but only for non-production use or with a total
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the
#  software will be made available under version 2 or later of the GNU General Public License.
#  If you use the software in violation of this license, your rights under the license will be
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all
#  warranties and conditions. If you use this license's text or the "Business Source License" name
#  and trademark, you must comply with the Licensor's covenants, which include specifying the
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
#  Grant, and not modifying the license in any other way.

from typing import Any

from django.apps import apps
from django.contrib.postgres.search import SearchQuery, SearchVector
from django.core.cache import cache
from rest_framework import pagination, status
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.views import APIView

from core.helpers import searchable_models


class CustomPageNumberPagination(pagination.PageNumberPagination):
    def get_paginated_response(self, data):
        return Response(
            {
                "pages": self.page.paginator.num_pages,
                "next": self.get_next_link(),
                "previous": self.get_previous_link(),
                "results": data,
            }
        )


class SearchView(APIView):
    """Handles handling of search queries through specified format 'model_name:search_term'.

    This class inherits from the APIView and redefines its get() method for performing
    specific search operations.

    Attributes:
        get: A method that executes the search operation based on request's parameters.
    """

    def get(self, request: Request, *args: Any, **kwargs: Any) -> Response:
        """Handle GET request on the SearchView.

        Args:
            request (Request): The request from the client. Must include 'term'
            in the query parameters.
            *args (Any): Variable length argument list.
            **kwargs (Any): Arbitrary keyword arguments.

        Returns:
            Response: The search results, paginated if applicable, or an error
            message if the search term is not properly formatted or if the
            model name is invalid.
        """
        term = request.query_params.get("term", "")
        if not term or ":" not in term:
            return Response(
                {"error": 'Search term must be in the format "ModelName: searchTerm"'},
                status=status.HTTP_400_BAD_REQUEST,
            )

        model_name, search_term = term.split(":")
        model_name = model_name.strip()
        search_term = search_term.strip()

        # Get the model, serializer, search fields, and display function dynamically
        model_info = searchable_models.get(model_name)
        if not model_info:
            return Response(
                {"error": "Invalid model name"}, status=status.HTTP_400_BAD_REQUEST
            )

        app_name, serializer, search_fields, display, path = model_info.values()
        model = apps.get_model(app_name, model_name)

        # Create a SearchVector for each specified field
        vectors = [SearchVector(field) for field in search_fields]

        # Combine the vectors
        vector = vectors.pop()
        for item in vectors:
            vector += item

        # Perform the search
        query = SearchQuery(search_term)
        cache_key = f"search:{model_name}:{search_term}"
        results = cache.get(cache_key)
        if results is None:
            results = model.objects.annotate(search=vector).filter(
                search=query, organization_id=request.user.organization_id  # type: ignore
            )
            cache.set(cache_key, results, 60 * 5)  # Cache the results for 5 minutes

        # Paginate the results
        paginator = CustomPageNumberPagination()
        page = paginator.paginate_queryset(results, request)
        if page is not None:
            data = [
                {
                    "display": display(result),
                    "path": path(result),
                    "model_name": model_name,
                }
                for result in page
            ]
            return paginator.get_paginated_response(data)

        # Serialize the results
        data = [
            {
                "path": path(result),
                "display": display(result),
                "model_name": model_name,
            }
            for result in results
        ]
        return Response(data, status=status.HTTP_200_OK)
