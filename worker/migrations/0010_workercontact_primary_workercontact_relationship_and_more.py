# Generated by Django 4.1.2 on 2022-11-02 04:47

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("worker", "0009_alter_worker_address_line_1_and_more"),
    ]

    operations = [
        migrations.AddField(
            model_name="workercontact",
            name="primary",
            field=models.BooleanField(
                default=False,
                help_text="Is this the primary contact?",
                unique=True,
                verbose_name="Primary",
            ),
        ),
        migrations.AddField(
            model_name="workercontact",
            name="relationship",
            field=models.CharField(
                blank=True,
                help_text="Relationship to the worker.",
                max_length=255,
                null=True,
                verbose_name="Relationship",
            ),
        ),
        migrations.AlterField(
            model_name="workerprofile",
            name="hazmat_expiration_date",
            field=models.DateField(
                blank=True,
                help_text="Hazmat Endorsement Expiration Date.",
                null=True,
                verbose_name="Hazmat Expiration Date",
            ),
        ),
        migrations.AlterField(
            model_name="workerprofile",
            name="hm_126_expiration_date",
            field=models.DateField(
                blank=True,
                help_text="HM126GF Training Expiration Date.",
                null=True,
                verbose_name="HM-126 Expiration Date",
            ),
        ),
    ]
