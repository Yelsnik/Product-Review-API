from server import SentimentServer
import grpc
from concurrent import futures
from sentiment_pb2_grpc import add_SentimentAnalysisServicer_to_server 
import signal
import sys
from utils import config

def grpc_server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_SentimentAnalysisServicer_to_server(SentimentServer(), server)
    server.add_insecure_port(config.grpc_address)
    server.start()
    print(f"NLP Service running on port {config.grpc_address}")

     # Handle graceful shutdown on SIGINT (Ctrl+C)
    def handle_shutdown(sig, frame):
        print("\nShutting down gRPC server gracefully...")
        server.stop(0)  # Stop the server immediately
        sys.exit(0)  # Exit the program
    
    # Catch SIGINT (Ctrl + C)
    signal.signal(signal.SIGINT, handle_shutdown)
    signal.signal(signal.SIGTERM, handle_shutdown)

    # Wait for termination (but allow graceful shutdown)
    server.wait_for_termination()


if __name__ == "__main__":
    grpc_server()