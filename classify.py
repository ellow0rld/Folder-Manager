import sys
import json
import joblib
import os
import numpy as np
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.decomposition import LatentDirichletAllocation

# Example predefined folder categories (modify as needed)
CATEGORIES = ["Math", "NLP", "ML", "Work", "Finance", "Misc"]

def classify(text):
    vectorizer = TfidfVectorizer(stop_words="english")
    X = vectorizer.fit_transform([text])

    lda = LatentDirichletAllocation(n_components=len(CATEGORIES), random_state=42)
    lda.fit(X)

    topic_distribution = lda.transform(X)[0]
    best_topic_idx = np.argmax(topic_distribution)

    return CATEGORIES[best_topic_idx]

if __name__ == "__main__":
    input_text = sys.stdin.read().strip()
    if not input_text:
        print(json.dumps({"error": "No input text"}))
        sys.exit(1)

    suggested_category = classify(input_text)
    print(json.dumps({"topic": suggested_category}))
