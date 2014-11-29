#coding:utf8
import requests
payload = {'keywords': ['丰田']}
r = requests.post("http://127.0.0.1:8001", data=payload)
print r.text
