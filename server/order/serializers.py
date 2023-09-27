# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------
from rest_framework import serializers

from movements.models import Movement
from order import models
from utils.serializers import GenericSerializer


class OrderControlSerializer(GenericSerializer):
    """A serializer for the `OrderControl` model.

    A serializer class for the OrderControl model. This serializer is used
    to convert the OrderControl model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model
    """

    class Meta:
        """
        Metaclass for OrderControlSerializer

        Attributes:
            model (OrderControl): The model that the serializer is for.
        """

        model = models.OrderControl


class OrderTypeSerializer(GenericSerializer):
    """A serializer for the `OrderType` model.

    A serializer class for the OrderType Model. This serializer is used
    to convert the OrderType model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model.

    """

    class Meta:
        """Metaclass for OrderTypeSerializer

        Attributes:
            model (models.OrderType): The model that the serializer is for.
        """

        model = models.OrderType

    def validate_name(self, value: str) -> str:
        """Validate the `name` field of the OrderType model.

        This method validates the `name` field of the OrderType model.
        It checks if the order type with the given name already exists in the organization.
        If the order type exists, it raises a validation error.

        Args:
            value (str): The value of the `name` field.

        Returns:
            str: The value of the `name` field.

        Raises:
            serializers.ValidationError: If the order type with the given name already exists in the
             organization.
        """
        organization = super().get_organization

        queryset = models.OrderType.objects.filter(
            organization=organization,
            name__iexact=value,  # iexact performs a case-insensitive search
        )

        # Exclude the current instance if updating
        if self.instance:
            queryset = queryset.exclude(pk=self.instance.pk)

        if queryset.exists():
            raise serializers.ValidationError(
                "Order Type with this `name` already exists. Please try again."
            )

        return value


class ReasonCodeSerializer(GenericSerializer):
    """A serializer for the `ReasonCode` model.

    A serializer class for the ReasonCode Model. This serializer is used
    to convert the ReasonCode model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model.
    """

    class Meta:
        """Metaclass for ReasonCodeSerializer

        Attributes:
            model (models.ReasonCode): The model that the serializer is for.
        """

        model = models.ReasonCode

    def validate_code(self, value: str) -> str:
        """Validate the `code` field of the ReasonCode model.

        This method validates the `code` field of the ReasonCode model.
        It checks if the reason code with the given code already exists in the organization.
        If the reason code exists, it raises a validation error.

        Args:
            value (str): The value of the `code` field.

        Returns:
            str: The value of the `code` field.

        Raises:
            serializers.ValidationError: If the reason code with the given code already exists in the
             organization.
        """
        organization = super().get_organization

        queryset = models.ReasonCode.objects.filter(
            organization=organization,
            code__iexact=value,  # iexact performs a case-insensitive search
        )

        # Exclude the current instance if updating
        if self.instance:
            queryset = queryset.exclude(pk=self.instance.pk)

        if queryset.exists():
            raise serializers.ValidationError(
                "Reason Code with this `code` already exists. Please try again."
            )

        return value


class OrderSerializer(GenericSerializer):
    """A serializer for the `Order` model.

    A serializer class for the Order Model. This serializer is used
    to convert the Order model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model.
    """

    additional_charges = serializers.PrimaryKeyRelatedField(
        many=True,
        queryset=models.AdditionalCharge.objects.all(),
        help_text="Additional charges for the order",
        required=False,
        allow_null=True,
    )
    movements = serializers.PrimaryKeyRelatedField(
        many=True,
        queryset=Movement.objects.all(),
        help_text="Movements for the order",
        required=False,
        allow_null=True,
    )
    order_documentation = serializers.PrimaryKeyRelatedField(
        many=True,
        queryset=models.OrderDocumentation.objects.all(),
        help_text="Documentation for the order",
        required=False,
        allow_null=True,
    )
    order_comments = serializers.PrimaryKeyRelatedField(
        many=True,
        queryset=models.OrderComment.objects.all(),
        help_text="Comments for the order",
        required=False,
        allow_null=True,
    )

    class Meta:
        """Metaclass for OrderSerializer

        Attributes:
            model (models.Order): The model that the serializer is for.
        """

        model = models.Order
        extra_fields = (
            "additional_charges",
            "movements",
            "order_documentation",
            "order_comments",
        )


class OrderDocumentationSerializer(GenericSerializer):
    """A serializer for the `OrderDocumentation` model.

    A serializer class for the OrderDocumentation Model. This serializer is used
    to convert the OrderDocumentation model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model.
    """

    class Meta:
        """Metaclass for OrderDocumentationSerializer

        Attributes:
            model (models.OrderDocumentation): The model that the serializer is for.
        """

        model = models.OrderDocumentation


class OrderCommentSerializer(GenericSerializer):
    """A serializer for the `OrderComment` model.

    A serializer class for the OrderComment Model. This serializer is used
    to convert the OrderComment model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model.
    """

    class Meta:
        """Metaclass for OrderCommentSerializer

        Attributes:
            model (models.OrderComment): The model that the serializer is for.
        """

        model = models.OrderComment


class AdditionalChargeSerializer(GenericSerializer):
    """A serializer for the `AdditionalCharge` model.

    A serializer class for the AdditionalCharge Model. This serializer is used
    to convert the AdditionalCharge model instances into a Python dictionary
    format that can be rendered into a JSON response. It also defines the fields
    that should be included in the serialized representation of the model.

    """

    class Meta:
        """Metaclass for AdditionalChargeSerializer

        Attributes:
            model (models.AdditionalCharge): The model that the serializer is for.
        """

        model = models.AdditionalCharge
