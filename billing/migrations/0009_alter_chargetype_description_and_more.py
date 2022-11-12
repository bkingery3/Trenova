# Generated by Django 4.1.2 on 2022-11-08 04:52

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("billing", "0008_alter_customerbillingprofile_options_and_more"),
    ]

    operations = [
        migrations.AlterField(
            model_name="chargetype",
            name="description",
            field=models.CharField(
                blank=True, default="test", max_length=100, verbose_name="Description"
            ),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name="customercontact",
            name="email",
            field=models.EmailField(
                blank=True,
                default="test@test.com",
                help_text="Contact email",
                max_length=150,
                verbose_name="Email",
            ),
            preserve_default=False,
        ),
        migrations.AlterField(
            model_name="customercontact",
            name="title",
            field=models.CharField(
                blank=True,
                default=1,
                help_text="Contact title",
                max_length=100,
                verbose_name="Title",
            ),
            preserve_default=False,
        ),
    ]
