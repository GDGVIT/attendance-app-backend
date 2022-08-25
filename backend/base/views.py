import firebase_admin
from firebase_admin import auth
from geopy import distance
from rest_framework import status, viewsets
from rest_framework.decorators import api_view
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


@api_view(["POST"])
def get_user(request):
    data = request.data
    # return Response(status=status.HTTP_100_CONTINUE)
    token = data["token"]
    club_name = data["club_name"]
    print(club_name)
    # return Response(status=status.HTTP_100_CONTINUE)
    club = models.Club.objects.get(name=club_name)
    # return Response(status=status.HTTP_100_CONTINUE)
    #
    decoded_token = auth.verify_id_token(token)
    uid = decoded_token["uid"]
    try:
        u = auth.get_user(uid)
        #phno = u.phone_number
        email=u.email
        #print(#phno)
        user = models.ClubMember.objects.get(club=club, email=email)
        s = serializers.ClubMemberSerializer(user)
        return Response(s.data)
    except models.ClubMember.DoesNotExist:
        return Response(status=status.HTTP_404_NOT_FOUND)


@api_view(["POST"])
def new_user(request):
    data = request.data
    print(data)
    token = data["token"]
    club_name = data["club_name"]
    decoded_token = auth.verify_id_token(token)
    uid = decoded_token["uid"]
    u = auth.get_user(uid)
    #phno = u.phone_number
    email=u.email
    name = u.display_name
    club = models.Club.objects.get(name=club_name)
    user, created = models.ClubMember.objects.get_or_create(
        name=name, email=email, attendence=0, is_admin=0, club=club
    )
    if created:
        user.save()
        return Response(status=status.HTTP_201_CREATED)
    else:
        return Response(status=status.HTTP_409_CONFLICT)


@api_view(["POST"])
def take_attendence(request):
    data = request.data
    print(data)
    token = data["token"]
    club_name = data["club_name"]
    lat = data["lat"]
    long = data["long"]
    club = models.Club.objects.get(name=club_name)
    decoded_token = auth.verify_id_token(token)
    uid = decoded_token["uid"]
    u = auth.get_user(uid)
    #phno = u.phone_number
    email=u.email
    user = models.ClubMember.objects.get(email=email, club=club)
    if user.is_admin != 1:
        return Response(status=status.HTTP_403_FORBIDDEN)
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


@api_view(["POST"])
def give_attendence(request):
    data = request.data
    print(data)
    token = data["token"]
    club_name = data["club_name"]
    lat = data["lat"]
    long = data["long"]
    utc = data["utc"]
    decoded_token = auth.verify_id_token(token)
    uid = decoded_token["uid"]
    u = auth.get_user(uid)
    #phno = u.phone_number
    email=u.email
    try:
        club = models.Club.objects.get(name=club_name)
        state = models.StateVariable.objects.get(club=club)
        latitude = float(lat)
        longitude = float(long)
        user_loc = (latitude, longitude)
        location = (float(state.latitude), float(state.longitude))
        distance_meters = distance.distance(user_loc, location).meters
        if distance_meters <= 50:
            user = models.ClubMember.objects.get(email=email, club=club)
            user.attendence += 1
            user.last_date = utc
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
