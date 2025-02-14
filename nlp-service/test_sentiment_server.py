import unittest
from unittest.mock import Mock, patch
from server import SentimentServer
import grpc
# import server
from sentiment_pb2 import SentimentResponse

class TestSentimentServer(unittest.TestCase):
    def setUp(self):
        self.server = SentimentServer()
        self.context = Mock()

    @patch('server.analyzer')
    def test_analyze_success(self, mock_analyzer):
        mock_request = Mock()
        mock_analyzer.return_value = (0.8, 'positive')

        response = self.server.Analyze(mock_request, self.context)

        # Assert
        self.assertIsInstance(response, SentimentResponse)
        self.assertAlmostEqual(response.score, 0.8, places=6)
        self.assertEqual(response.label, "positive")
        mock_analyzer.assert_called_once_with(mock_request)
        # self.context.set_code.assert_not_called()
        # self.context.set_details.assert_not_called()

    @patch('server.analyzer')
    def test_analyze_failure(self, mock_analyzer: Mock):
        mock_analyzer.side_effect = Exception("Model error")
        mock_request = Mock()

        response = self.server.Analyze(mock_request, self.context)

        self.assertEqual(response.score, 0)
        self.assertEqual(response.label, "")

        self.context.set_code.assert_called_once_with(grpc.StatusCode.INTERNAL)
        self.context.set_details.assert_called_once_with("Analysis failed: Model error")


if __name__ == "__main__":
    unittest.main()