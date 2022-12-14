from django.urls import include, path
from django.views.generic import TemplateView
from rest_framework import routers
from rest_framework.schemas import get_schema_view

from . import views

router = routers.DefaultRouter()

router.register(r"members", views.ClubMemberViewSet)
router.register(r"clubs", views.ClubViewSet)
router.register(r"state", views.StateVariableViewSet)


urlpatterns = [
    path("api-auth/", include("rest_framework.urls", namespace="rest_framework")),
    path("get_user/", views.get_user),
    path("new_user/", views.new_user),
    path("new_user_batch/", views.new_user_batch),
    path("take_attendence/", views.take_attendence),
    path("give_attendence/", views.give_attendence),
    path("attendence_state/<str:club_name>", views.attendence_state),
    path(
        "docs/",
        TemplateView.as_view(
            template_name="swagger-ui.html",
            extra_context={"schema_url": "openapi-schema", "swagger": "2.0"},
        ),
        name="openapi-schema",
    ),
    path(
        "",
        get_schema_view(
            title="Attendence App",
            description="API for attendence\
                (the execute buttons don't work :p)",
            version="1.0.0",
        ),
        name="openapi-schema",
    ),
    path("", include(router.urls)),
]
