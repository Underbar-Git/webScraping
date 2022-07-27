
import requests
import csv 

url_list = [
    'https://www.naver.com',
    'https://www.nate.com',
    'https://kr.indeed.com/jobs?q=python&l&from=searchOnHP&vjk=459437f7ce2c9b1c'
]

results = [] 
for idx, url in enumerate(url_list):
    r = requests.get(url=url, timeout=5)
    results.append((idx, url, r.status_code))


with open('urlCheck_Result.csv', 'w', newline='') as csvfile:
    fieldnames = ['번호', 'url', '응답코드']
    writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
    writer.writeheader()

    for result in results:
        writer.writerow({'번호': result[0], 'url': result[1], '응답코드': result[2]})
