import os
import sys
import json
import joblib
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.naive_bayes import MultinomialNB
import numpy as np

# Function to read text from files
def read_text_from_files(folder):
    texts = []
    for file_name in os.listdir(folder):
        file_path = os.path.join(folder, file_name)
        try:
            with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
                texts.append(f.read())
        except:
            continue
    return texts

def train_model(folders):
    labels = []
    docs = []

    for folder in folders:
        if os.path.isdir(folder):
            category = os.path.basename(folder)  # Folder name as category
            text_data = read_text_from_files(folder)
            docs.extend(text_data)
            labels.extend([category] * len(text_data))

    if not docs:
        print(json.dumps({"error": "No files found for training"}))
        return

    # Train Model
    vectorizer = TfidfVectorizer()
    X = vectorizer.fit_transform(docs)
    model = MultinomialNB()
    y = np.array(labels)
    model.fit(X, y)

    # Save model
    joblib.dump((vectorizer, model), "trained_classifier.pkl")
    print(json.dumps({"status": "Model trained successfully"}))

if __name__ == "__main__":
    # Receive folders as JSON input
    folders = json.loads(sys.stdin.read())
    train_model(folders)
