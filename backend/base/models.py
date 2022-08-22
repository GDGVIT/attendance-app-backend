from django.db import models


class Club(models.Model):
    name = models.CharField(max_length=200, primary_key=True)

    def __str__(self) -> str:
        return self.name


# Create your models here.
class ClubMember(models.Model):
    club = models.ForeignKey(Club, related_name="members", on_delete=models.CASCADE)
    name = models.CharField(max_length=200, null=False, blank=False)
    phone = models.CharField(max_length=10, null=False, blank=False)
    attendence = models.IntegerField(null=False, blank=False)
    is_admin = models.IntegerField(default=False)
    last_date = models.CharField(max_length=200)

    def __str__(self) -> str:
        return self.phone


class StateVariable(models.Model):
    # _id=models.BigAutoField(verbose_name="ID",primary_key=True,null=False)
    club = models.OneToOneField(
        Club, primary_key=True, related_name="states", on_delete=models.CASCADE
    )
    take_attendence = models.BooleanField()
    latitude = models.CharField(max_length=200)
    longitude = models.CharField(max_length=200)

    def __str__(self) -> str:
        return f"take_attendence:{self.take_attendence}"
