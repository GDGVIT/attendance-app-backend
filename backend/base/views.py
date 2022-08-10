from django.shortcuts import render
from rest_framework import status, viewsets
from rest_framework.decorators import action, api_view
from rest_framework.response import Response

# Create your views here.
from . import models, serializers

taking_attendence = False


class ClubMemberViewSet(viewsets.ModelViewSet):
    queryset = models.ClubMember.objects.all().order_by("id")
    serializer_class = serializers.ClubMemberSerializer


@api_view(["GET"])
def get_user(request, phno: int):
    try:
        user = models.ClubMember.objects.get(phone=phno)
        s = serializers.ClubMemberSerializer(user)
        return Response(s.data)
    except models.ClubMember.DoesNotExist:
        return Response(status=status.HTTP_404_NOT_FOUND)


@api_view(["PUT"])
def new_user(request, name: str, phno: int):
    user, created = models.ClubMember.objects.get_or_create(
        name=name, phone=phno, attendence=0, is_admin=0
    )
    if created:
        user.save()
        return Response(status=status.HTTP_201_CREATED)
    else:
        return Response(status=status.HTTP_409_CONFLICT)


@api_view(["GET"])
def take_attendence(request):
    global taking_attendence
    taking_attendence = not taking_attendence
    return Response({"state": taking_attendence})


@api_view(["GET"])
def attendence_state(request):
    return Response({"state": taking_attendence})


@api_view(["GET"])
def give_attendence(request, phno: int):
    try:
        user = models.ClubMember.objects.get(phone=phno)
        user.attendence += 1
        user.save()
        s = serializers.ClubMemberSerializer(user)
        return Response(s.data)
    except models.ClubMember.DoesNotExist:
        return Response(status=status.HTTP_404_NOT_FOUND)
