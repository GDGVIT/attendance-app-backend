from django.db import models

# Create your models here.
class ClubMember(models.Model):
    name=models.CharField(max_length=200)
    phone=models.CharField(max_length=10)
    attendence=models.IntegerField()
    is_admin=models.BooleanField(default=False)
