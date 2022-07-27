"""
url : 
    https://kr.indeed.com/jobs?q=python&start=0
    https://kr.indeed.com/jobs?q=python&start=10
    https://kr.indeed.com/jobs?q=python&start=20
    ...
    https://kr.indeed.com/jobs?q=python&start=100

data : 리스트 내역 
result : csv 파일로 저장
"""

from bs4 import BeautifulSoup
import requests


BASE_URL = 'https://kr.indeed.com/jobs?q=python&start='

def makeurl(pagenum):
    return BASE_URL + str(pagenum * 10)


result = []

for i in range(0, 11):
    mUrl = makeurl(i)
    r = requests.get(url=mUrl, timeout=5)
    # print(i, mUrl, r.status_code, '파싱 시작합니다.')
    
    soup = BeautifulSoup(r.text, 'html.parser')
    jobSection = soup.find_all('div', attrs={'class', 'job_seen_beacon'})
    for section in jobSection:
        job = section.find('h2', attrs={'class', 'jobTitle'})
        location = section.find('div', attrs={'class', 'companyLocation'}).text
        dic = dict()
        dic['title'] = job.text
        dic['link '] = f"https://kr.indeed.com/채용보기?jk={job.find('a').attrs['data-jk']}"
        dic['location'] = location
        
        print(i, mUrl, r.status_code, ' ------> 파싱 완료.')
        result.append(dic)
print(result)