from rest_framework import routers
from django.urls import path,include
from . import views
from rest_framework.schemas import get_schema_view
from django.views.generic import TemplateView
router = routers.DefaultRouter()
router.register(r'members', views.ClubMemberViewSet)
urlpatterns = [
    path('', include(router.urls)),
    path('api-auth/', include('rest_framework.urls', namespace='rest_framework')),
    path('get_user/<int:phno>',views.get_user),
    path('docs/', TemplateView.as_view(
        template_name='swagger-ui.html',
        extra_context={'schema_url':'openapi-schema','swagger':'2.0'}
    ), name='openapi-schema'),
    path('api', get_schema_view(
        title="Attendence App",
        description="API for attendence(the execute buttons don't work :p)",
        version="1.0.0"
    ), name='openapi-schema'),
]