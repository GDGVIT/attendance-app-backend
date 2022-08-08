from django.shortcuts import render
from rest_framework import viewsets,status
from rest_framework.response import Response
from rest_framework.decorators import api_view,action
# Create your views here.
from .import models
from . import serializers
class ClubMemberViewSet(viewsets.ModelViewSet):
    queryset=models.ClubMember.objects.all().order_by('id')
    serializer_class=serializers.ClubMemberSerializer

@api_view(['GET'])
def get_user(request,phno:int):
    try:
        user=models.ClubMember.objects.get(phone=phno)
        s=serializers.ClubMemberSerializer(user)
        return Response(s.data)
    except models.ClubMember.DoesNotExist:
        return Response(status=status.HTTP_404_NOT_FOUND)

#@api_view(['PUT'])


    