# Generated by Django 4.2 on 2023-04-20 04:07

from django.db import migrations
import utils.models


class Migration(migrations.Migration):

    dependencies = [
        ('route', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='routecontrol',
            name='distance_method',
            field=utils.models.ChoiceField(choices=[('Google', 'Google'), ('Monta', 'Monta')], default='Monta', help_text='Distance method for the company.', max_length=6, verbose_name='Distance Method'),
        ),
    ]
