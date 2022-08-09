from django.db import models


# Create your models here.
class ClubMember(models.Model):
    name=models.CharField(max_length=200,null=False,blank=False)
    phone=models.CharField(max_length=10,null=False,blank=False)
    attendence=models.IntegerField(null=False,blank=False)
    is_admin=models.IntegerField(default=False)
