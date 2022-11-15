"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

from django.contrib import admin

from core.mixins import MontaAdminMixin, MontaStackedInlineMixin
from .models import Depot, DepotDetail, Organization


@admin.register(Organization)
class OrganizationAdmin(admin.ModelAdmin[Organization]):
    """
    Organization Admin
    """

    list_display: tuple[str, ...] = (
        "name",
        "scac_code",
        "org_type",
        "timezone",
    )
    list_filter: tuple[str, ...] = ("org_type",)
    search_fields: tuple[str, ...] = (
        "name",
        "scac_code",
    )


class DepotDetailInline(MontaStackedInlineMixin[Depot, DepotDetail]):
    """
    Depot Detail Admin
    """

    model: type[DepotDetail] = DepotDetail
    can_delete = False
    verbose_name_plural = "Depot Details"
    fk_name = "depot"


@admin.register(Depot)
class DepotAdmin(MontaAdminMixin[Depot]):
    """
    Depot Admin
    """

    list_display: tuple[str, ...] = (
        "name",
        "description",
    )
    list_filter: tuple[str, ...] = ("name",)
    search_fields: tuple[str, ...] = ("name",)
    inlines: tuple[type[DepotDetailInline]] = (DepotDetailInline,)
