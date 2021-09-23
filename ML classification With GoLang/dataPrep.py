# -*- coding: utf-8 -*-
"""
Created on Fri Sep 17 10:59:06 2021

@author: maher.alrefaai
"""

import pandas as pd
df=pd.read_csv("dataset.csv")
from urllib.parse import urlparse
import re
from dateutil.parser import parse


Fdf = pd.DataFrame(columns = [ 'href', 'Base','UrlRest','urlQuery','date','date2','Id','split_string','isId','isDate','isQuary','NowordsAfterBase','nothingAfterBase','Target'])



for i in range(len(df)):
    href =df.iloc[i,0]
    UrlParse = urlparse(href)
    Base=UrlParse[1]
    UrlRest=UrlParse[2]
    urlQuery=UrlParse[4]
    Target=df.iloc[i,1] 
    date = re.findall(r'(\d+/\d+/\d+)',UrlRest)
    date2 = re.findall(r'(\d+-\d+-\d+)',UrlRest)
    split_string = re.split("\s|(?<!\d)[/-](?!\d)", UrlRest)
    Id=re.findall(r'\d{5,}',UrlRest)
    
    isQuary=True
    if(len(urlQuery)==0):
        isQuary=False
    isId=True
    if(len(Id)==0):
        isId=False
    isDate=True
    if(len(date)==0 and  len(date2)==0):
        isDate=False
    
    nothingAfterBase=False
    NowordsAfterBase=0
    for word in range(len(split_string)):
        if(len(split_string[word])!=0):
            NowordsAfterBase=NowordsAfterBase+1
    
    if(NowordsAfterBase==0):
        nothingAfterBase=True
    Dect= {'href':href, 'Base':Base,'UrlRest':UrlRest,'urlQuery':urlQuery,'date':date,'date2':date2,'Id':Id,'split_string':split_string,'isId':isId,'isDate':isDate,'isQuary':isQuary,'NowordsAfterBase':NowordsAfterBase,'nothingAfterBase':nothingAfterBase,'Target':Target}
    
    print(nothingAfterBase)
    Fdf=Fdf.append(Dect, ignore_index = True)



Fdf.to_csv('fullFeatures.csv',index=False)


