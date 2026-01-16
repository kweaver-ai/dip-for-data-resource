# -*- coding: utf-8 -*-
"""
@Time ： 2024/7/24 11:26
@Auth ： Xia.wang
@File ：sql_similarity_checker.py
@IDE ：PyCharm
"""
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
import nltk
from nltk.corpus import stopwords
from nltk.tokenize import word_tokenize

# 下载NLTK所需的停用词和其他资源
try:
    nltk.corpus.stopwords.words('english')
except LookupError:
    nltk.download('stopwords')
    nltk.download('punkt')
    nltk.data.path.append('./tools_data/punkt')
    nltk.data.path.append('./tools_data/stopwords')


# 文本预处理函数：去除停用词、标点符号等
def preprocess_text(text):
    # 分词并去除停用词和标点符号
    stop_words = set(stopwords.words('english'))  # 可以根据需要替换成其他语言的停用词表
    word_tokens = word_tokenize(text.lower())
    filtered_text = [word for word in word_tokens if word.isalnum() and word not in stop_words]
    return ' '.join(filtered_text)


# 使用TF-IDF向量化文本
def tfidf_vectorize(text1, text2):
    # 对两个文本进行预处理
    processed_text1 = preprocess_text(text1)
    processed_text2 = preprocess_text(text2)

    # TF-IDF向量化
    vectorizer = TfidfVectorizer()
    tfidf_matrix = vectorizer.fit_transform([processed_text1, processed_text2])

    return tfidf_matrix


# 计算余弦相似度
def cosine_similarity_text(text1, text2):
    tfidf_matrix = tfidf_vectorize(text1, text2)
    cosine_sim = cosine_similarity(tfidf_matrix[0], tfidf_matrix[1])[0][0]
    return "{:.2f}".format(cosine_sim)


if __name__ == '__main__':
    # 示例SQL语句
    # sql1 = "SELECT address FROM address WHERE district = 'Alberta' LIMIT 3;"
    # sql2 = "SELECT address FROM address WHERE district = 'Alberta' LIMIT 3;"

    # 计算余弦相似度
    similarity = cosine_similarity_text(sql1, sql2)
    # print(f"Cosine similarity between SQL queries: {similarity}")
    import pandas as pd