from django.shortcuts import render
from rest_framework import viewsets
# Create your views here.
from .import models
from . import serializers
class ClubMemberViewSet(viewsets.ModelViewSet):
    queryset=models.ClubMember.objects.all().order_by('id')
    serializer_class=serializers.ClubMemberSerializer