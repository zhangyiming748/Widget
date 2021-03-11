import bs4
import requests

DOWNLOAD_URL = 'https://www.xvideos.com/video2487943/asian_squirts_on_futanari_teen_'


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
    links=[]
    soup=bs4.BeautifulSoup(html, features="html.parser")
    mozaique = soup.find('div',attrs={'id':'related-video'})
    print(mozaique)
    return links
if __name__ == '__main__':
    url = DOWNLOAD_URL
    html = download_page(url)
    links = parse_html(html)
    with open('./Straplezz.txt', 'a+', encoding='utf8') as f:
        for line in links:
            f.write(line)
            f.write('\n')