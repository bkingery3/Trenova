# Generated by Django 4.1.2 on 2022-10-31 05:56

import django.db.models.deletion
import django_extensions.db.fields
import localflavor.us.models
import phonenumber_field.modelfields
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("organization", "0004_alter_organization_language_and_more"),
    ]

    operations = [
        migrations.CreateModel(
            name="Depot",
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
                    "name",
                    models.CharField(
                        help_text="The name of the depot.",
                        max_length=255,
                        verbose_name="Depot Name",
                    ),
                ),
                (
                    "description",
                    models.TextField(
                        blank=True,
                        help_text="The description of the depot.",
                        max_length=255,
                        null=True,
                        verbose_name="Depot Description",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="The organization that the depot belongs to.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="depots",
                        related_query_name="depot",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Depot",
                "verbose_name_plural": "Depots",
                "ordering": ["name"],
            },
        ),
        migrations.CreateModel(
            name="DepotDetail",
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
                    "address_line_1",
                    models.CharField(
                        help_text="The address line 1 of the depot.",
                        max_length=255,
                        verbose_name="Address",
                    ),
                ),
                (
                    "address_line_2",
                    models.CharField(
                        blank=True,
                        help_text="The address line 2 of the depot.",
                        max_length=255,
                        null=True,
                        verbose_name="Address",
                    ),
                ),
                (
                    "city",
                    models.CharField(
                        help_text="The city of the depot.",
                        max_length=255,
                        verbose_name="City",
                    ),
                ),
                (
                    "state",
                    localflavor.us.models.USStateField(
                        blank=True,
                        help_text="The state of the depot.",
                        max_length=2,
                        null=True,
                        verbose_name="State",
                    ),
                ),
                (
                    "zip_code",
                    localflavor.us.models.USZipCodeField(
                        blank=True,
                        help_text="The zip code of the depot.",
                        max_length=10,
                        null=True,
                        verbose_name="Zip Code",
                    ),
                ),
                (
                    "phone_number",
                    phonenumber_field.modelfields.PhoneNumberField(
                        blank=True,
                        help_text="The phone number of the depot.",
                        max_length=128,
                        null=True,
                        region=None,
                        verbose_name="Phone Number",
                    ),
                ),
                (
                    "alternate_phone_number",
                    phonenumber_field.modelfields.PhoneNumberField(
                        blank=True,
                        help_text="The alternate phone number of the depot.",
                        max_length=128,
                        null=True,
                        region=None,
                        verbose_name="Alternate Phone Number",
                    ),
                ),
                (
                    "fax_number",
                    phonenumber_field.modelfields.PhoneNumberField(
                        blank=True,
                        help_text="The fax number of the depot.",
                        max_length=128,
                        null=True,
                        region=None,
                        verbose_name="Fax Number",
                    ),
                ),
                (
                    "depot",
                    models.OneToOneField(
                        help_text="The depot that the depot detail belongs to.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="depot_details",
                        related_query_name="depot_detail",
                        to="organization.depot",
                        verbose_name="Depot",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="The organization that the depot detail belongs to.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="depot_details",
                        related_query_name="depot_detail",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Depot Detail",
                "verbose_name_plural": "Depot Details",
                "ordering": ["depot"],
            },
        ),
        migrations.AddIndex(
            model_name="depotdetail",
            index=models.Index(fields=["depot"], name="organizatio_depot_i_35e998_idx"),
        ),
        migrations.AddIndex(
            model_name="depot",
            index=models.Index(fields=["name"], name="organizatio_name_33faff_idx"),
        ),
    ]
