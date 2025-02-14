import unittest
from textblob import TextBlob
from server import analyzer

class TestAnalyzeTextBlob(unittest.TestCase):
    def test_positive_sentiment(self):
        text = "I love this product!"
        score, label = analyzer(text)
        self.assertGreaterEqual(score, 0)
        self.assertEqual(label, "positive")
        print(score, label)

    def test_mixed_sentiment(self):
        text = "I love the design, but the material is bad."
        score, label = analyzer(text)
        self.assertLessEqual(score, 0)
        print(score, label)

if __name__ == '__main__':
    unittest.main()