# Generated by Django 4.1.4 on 2022-12-18 17:15

from django.db import migrations
import utils.models


class Migration(migrations.Migration):

    dependencies = [
        ("equipment", "0004_alter_equipmenttype_options_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="equipmenttypedetail",
            name="equipment_class",
            field=utils.models.ChoiceField(
                choices=[
                    ("UNDEFINED", "UNDEFINED"),
                    ("CAR", "Car"),
                    ("VAN", "Van"),
                    ("PICKUP", "Pickup"),
                    ("WALK-IN", "Walk-In"),
                    ("STRAIGHT", "Straight Truck"),
                    ("TRACTOR", "Tractor"),
                    ("TRAILER", "Trailer"),
                ],
                default="UNDEFINED",
                help_text="Class of the equipment type.",
                max_length=9,
                verbose_name="Equipment Class",
            ),
        ),
    ]
