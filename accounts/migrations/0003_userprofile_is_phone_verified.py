# Generated by Django 4.1.7 on 2023-03-04 20:17

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("accounts", "0002_token"),
    ]

    operations = [
        migrations.AddField(
            model_name="userprofile",
            name="is_phone_verified",
            field=models.BooleanField(
                default=False,
                help_text="Designates whether the user's phone number has been verified.",
                verbose_name="Phone Number Verified",
            ),
        ),
    ]
