import requests

with open("test.csv") as f:
    f = f.read()
    files = {"file": f}
    r = requests.post("http://127.0.0.1:8000/new_user_batch/", files=files)
