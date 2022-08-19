from dataclasses import fields

from rest_framework import serializers

from .models import Club, ClubMember, StateVariable


class ClubSerializer(serializers.ModelSerializer):
    members = serializers.StringRelatedField(many=True)
    states = serializers.StringRelatedField()

    class Meta:
        model = Club
        fields = ("name", "members", "states")


class ClubMemberSerializer(serializers.ModelSerializer):
    # club=serializers.StringRelatedField(many=True)
    class Meta:
        model = ClubMember
        fields = ("name", "phone", "attendence", "is_admin", "club")


class StateVariableSerializer(serializers.ModelSerializer):
    # club_name=serializers.StringRelatedField(many=True)
    class Meta:
        model = StateVariable
        fields = ("take_attendence","club","latitude","longitude")
