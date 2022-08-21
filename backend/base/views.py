import firebase_admin
from django.shortcuts import render
from firebase_admin import auth
from geopy import distance
from pyexpat import model
from rest_framework import status, viewsets
from rest_framework.decorators import action, api_view
from rest_framework.response import Response

# Create your views here.
from . import models, serializers

auth_app = firebase_admin.initialize_app()
print(auth_app.name)


class ClubViewSet(viewsets.ModelViewSet):
    queryset = models.Club.objects.all()
    serializer_class = serializers.ClubSerializer


class ClubMemberViewSet(viewsets.ModelViewSet):
    queryset = models.ClubMember.objects.all().order_by("id")
    serializer_class = serializers.ClubMemberSerializer


class StateVariableViewSet(viewsets.ModelViewSet):
    queryset = models.StateVariable.objects.all()
    serializer_class = serializers.StateVariableSerializer


@api_view(["GET"])
def get_user(request, token: str):
    decoded_token = auth.verify_id_token(token)
    uid = decoded_token["uid"]
    try:
        u = auth_app.auth().get_user(uid)
        phno = u.phone_number
        print(phno)
        user = models.ClubMember.objects.get(phone=phno)
        s = serializers.ClubMemberSerializer(user)
        return Response(s.data)
    except models.ClubMember.DoesNotExist:
        return Response(status=status.HTTP_404_NOT_FOUND)


@api_view(["PUT"])
def new_user(request, club_name: str, name: str, phno: int):
    club = models.Club.objects.get(name=club_name)
    user, created = models.ClubMember.objects.get_or_create(
        name=name, phone=phno, attendence=0, is_admin=0, club=club
    )
    if created:
        user.save()
        return Response(status=status.HTTP_201_CREATED)
    else:
        return Response(status=status.HTTP_409_CONFLICT)


@api_view(["GET"])
def take_attendence(request, club_name: str, lat: str, long: str):
    club = models.Club.objects.get(name=club_name)

    state, created = models.StateVariable.objects.get_or_create(club=club)
    if created:
        state.take_attendence = True
    else:
        state.take_attendence = not state.take_attendence
    state.latitude = lat
    state.longitude = long
    state.save()
    return Response({"state": state.take_attendence})


@api_view(["GET"])
def attendence_state(request, club_name: str):
    club = models.Club.objects.get(name=club_name)
    state = models.StateVariable.objects.get(club=club)
    return Response({"state": state.take_attendence})


@api_view(["GET"])
def give_attendence(request, club_name: str, phno: int, lat: str, long: str):
    try:
        club = models.Club.objects.get(name=club_name)
        state = models.StateVariable.objects.get(club=club)
        latitude = float(lat)
        longitude = float(long)
        user_loc = (latitude, longitude)
        location = (float(state.latitude), float(state.longitude))
        distance_meters = distance.distance(user_loc, location).meters
        if distance_meters <= 50:
            user = models.ClubMember.objects.get(phone=phno, club=club)
            user.attendence += 1
            user.save()
            s = serializers.ClubMemberSerializer(user)
            return Response(s.data)
        else:
            return Response(
                {
                    "error": "user not in range",
                    "distance": distance_meters,
                    "message": "seriously where even are you?!",
                }
            )
    except models.ClubMember.DoesNotExist:
        return Response(status=status.HTTP_404_NOT_FOUND)
