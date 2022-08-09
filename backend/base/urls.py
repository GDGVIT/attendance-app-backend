from django.urls import include, path
from django.views.generic import TemplateView
from rest_framework import routers
from rest_framework.schemas import get_schema_view

from . import views

router = routers.DefaultRouter()
router.register(r"members", views.ClubMemberViewSet)
urlpatterns = [
    path("api-auth/", include("rest_framework.urls", namespace="rest_framework")),
    path("get_user/<int:phno>", views.get_user),
    path("new_user/<str:name>/<int:phno>", views.new_user),
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
            description="API for attendence(the execute buttons don't work :p)",
            version="1.0.0",
        ),
        name="openapi-schema",
    ),
    path("", include(router.urls)),
]
