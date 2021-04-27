import requests
from bs4 import BeautifulSoup
DOWNLOAD_URL = 'PH'

def download_page(url):
    headers = {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.36'
    }
    proxy = {
        'http': 'http://127.0.0.1:8889',
        'https': 'http://127.0.0.1:8889'
    }
    return requests.get(url, headers=headers,proxies=proxy).content


def parse_html(html):
    soup = BeautifulSoup(html, features="html.parser")
    titles = soup.find('ul',attrs={'id':'showAllChanelVideos'})
    #print(titles)
    for title in titles.findAll('a'):
        print('https://www.ph.com'+str(title['href']))


    # next_page = soup.find('span', attrs={'class': 'next'}).find('a')
    # if next_page:
    #     return movie_name_list, DOWNLOAD_URL + next_page['href']
    # return movie_name_list, None


if __name__ == '__main__':
    url = DOWNLOAD_URL
    html = download_page(url)
    parse_html(html)
