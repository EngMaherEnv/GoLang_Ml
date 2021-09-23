# -*- coding: utf-8 -*-
"""
Created on Fri Sep 17 16:48:20 2021

@author: maher.alrefaai
"""

import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn import metrics


df=pd.read_csv("fullFeatures.csv")



X=df[[ 'isId', 'isDate', 'isQuary', 'NowordsAfterBase',
       'nothingAfterBase']]  # Features
y=df['Target']  # Labels

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3) # 70% training and 30% test

clf=RandomForestClassifier(n_estimators=100)

clf.fit(X_train,y_train)

y_pred=clf.predict(X_test)



print("Accuracy:",metrics.accuracy_score(y_test, y_pred))

"""

mydf=df[[ 'isId', 'isDate', 'isQuary', 'NowordsAfterBase',
       'nothingAfterBase','Target']]

mypre=clf.predict([[False,False,False,2,False]])
print(mypre)
"""