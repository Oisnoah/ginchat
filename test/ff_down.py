import requests

# 预下载，获取m3u8文件，读出ts链接，并写入文档
# 2018-12-18 10:00:00
def down(header, url, file):
    r = requests.get(url, headers=header)
    m3u8_text = r.text
    #按行拆分m3u8
    m3u8_list = m3u8_text.split('\n')
    s=len(m3u8_list)
    
    with open(file, 'w') as f:
        f.write(r.text)