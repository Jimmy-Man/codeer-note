#!/usr/bin/python
# -*- coding: UTF-8 -*-

## 批量添加google翻译插件脚本
import os

##网站目录
#SitePath = '/data/www'
SitePath = '/home/wwwroot'
fileName = 'html_header.php'
#tempList = ['yourstore','rwd','iphone']
tempList = ['yourstore']

#temPath = 'includes/templates/yourtemplate/common/html_header.php';

for path in os.listdir(SitePath):
    if os.path.isdir(SitePath+'/'+path):
        print path
        ##查找文件
        for temp in tempList:
            fileName = SitePath + '/' + path + '/includes/templates/' + temp + '/common/html_header.php' 
            #print fileName        
            if(os.path.isfile(fileName)):            
                print fileName
                f = open(fileName,'r+')
                f.read                
                content = f.read()
                #print content
                if 'google_translate_element' in content:
                    ## 操作                
                    print 'yes'
                else:
                    #print 'none'                    
                    f.write("""\n\r
                        <div id="google_translate_element"></div><script type="text/javascript">
function googleTranslateElementInit() {
  new google.translate.TranslateElement({pageLanguage: 'en', includedLanguages: 'bg,cs,de,en,es,et,fi,fr,hr,hu,lb,lt,lv,nl,no,pl,pt,sk,sl,sq,sv,yi', layout: google.translate.TranslateElement.InlineLayout.HORIZONTAL}, 'google_translate_element');
}
</script><script type="text/javascript" src="//translate.google.com/translate_a/element.js?cb=googleTranslateElementInit"></script>
                    """)
                f.close
            else:
                #print 'none'
                continue

                    
            

