# Generated by Django 4.1.3 on 2022-11-14 15:18

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("order", "0009_order"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="ordercontrol",
            name="enforce_bill_to",
        ),
        migrations.AddField(
            model_name="ordercontrol",
            name="enforce_customer",
            field=models.BooleanField(
                default=False,
                help_text="Enforce Customer to being enter when entering an order.",
                verbose_name="Enforce Customer",
            ),
        ),
    ]
