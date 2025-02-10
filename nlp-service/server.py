import grpc
import sys
import os
from sentiment_pb2_grpc import SentimentAnalysisServicer
from sentiment_pb2 import SentimentResponse
from textblob import TextBlob

# # Get the absolute path of the "nlp" directory
# nlp_path = os.path.abspath(os.path.join(os.path.dirname(__file__), "nlp"))

# # Add it to sys.path
# sys.path.append(nlp_path)

# from nlp import sentiment_service_pb2_grpc as nlp
# from nlp import sentiment_service_pb2 as sentiment

def analyze_textblob(text):
    analysis = TextBlob(text)
    score = analysis.sentiment
    label = "positive" if score[0] >= 0 else "negative"
    return score[0], label

analyzer = analyze_textblob

class SentimentServer(SentimentAnalysisServicer):
    def __init__(self):
        self.model = "hello"
        

    def Analyze(self, request, context):
        try:
            score, label = analyzer(request)
            return SentimentResponse(score, label)
        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"Analysis failed: {str(e)}")
            return SentimentResponse() 
    
       
# try:
#             result = self.model
#             score="result['score'] * (1 if result['label'] == 'positive' else -1),"
#             label="result['label'].upper(),"
#             return sentiment.AnalyzeResponse(score, label)
#         except Exception as e:
#             context.set_code(grpc.StatusCode.INTERNAL)
#             context.set_details(f"Analysis failed: {str(e)}")
#             return sentiment.AnalyzeResponse()