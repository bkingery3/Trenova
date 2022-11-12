# Generated by Django 4.1.2 on 2022-11-07 05:08

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0006_alter_depotdetail_address_line_1_and_more"),
        ("order", "0004_reasoncode_qualifiercode"),
    ]

    operations = [
        migrations.CreateModel(
            name="OrderControl",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "auto_rate_orders",
                    models.BooleanField(
                        default=True,
                        help_text="Auto rate orders",
                        verbose_name="Auto Rate",
                    ),
                ),
                (
                    "calculate_distance",
                    models.BooleanField(
                        default=True,
                        help_text="Calculate distance for the order",
                        verbose_name="Calculate Distance",
                    ),
                ),
                (
                    "enforce_bill_to",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce bill to to being enter when entering an order.",
                        verbose_name="Enforce Bill To",
                    ),
                ),
                (
                    "enforce_rev_code",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce rev code code being entered when entering an order.",
                        verbose_name="Enforce Rev Code",
                    ),
                ),
                (
                    "enforce_shipper",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce shipper when putting in an order.",
                        verbose_name="Enforce Shipper",
                    ),
                ),
                (
                    "enforce_cancel_comm",
                    models.BooleanField(
                        default=False,
                        help_text="Enforce comment when cancelling an order.",
                        verbose_name="Enforce Voided Comm",
                    ),
                ),
                (
                    "generate_routes",
                    models.BooleanField(
                        default=False,
                        help_text="Generate routes for the organization",
                        verbose_name="Generate Routes",
                    ),
                ),
                (
                    "organization",
                    models.OneToOneField(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="order_control",
                        related_query_name="order_controls",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Order Control",
                "verbose_name_plural": "Order Controls",
                "ordering": ["organization"],
            },
        ),
    ]
