# -*- coding: utf-8 -*-
"""
Created on Mon Sep 20 15:51:16 2021

@author: maher.alrefaai
"""
import pandas as pd

df=pd.read_csv("fullFeatures.csv")


bol={False:0,True:1}
df['isId']=df['isId'].map(bol)
df['isDate']=df['isDate'].map(bol)
df['isQuary']=df['isQuary'].map(bol)
df['nothingAfterBase']=df['nothingAfterBase'].map(bol)

Inputs=df[[ 'isId', 'isDate', 'isQuary', 'NowordsAfterBase','nothingAfterBase', 'Target']]

Inputs=Inputs.sample(frac=1)
Inputs['isId'] = Inputs['isId'].astype(float)
Inputs['isDate'] = Inputs['isDate'].astype(float)
Inputs['isQuary'] = Inputs['isQuary'].astype(float)
Inputs['NowordsAfterBase'] = Inputs['NowordsAfterBase'].astype(float)
Inputs['nothingAfterBase'] = Inputs['nothingAfterBase'].astype(float)
Inputs['Target'] = Inputs['Target'].astype(float)
Inputs.to_csv('InputsGoLang.csv',index=False,header=False)
