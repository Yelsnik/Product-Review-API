from server import SentimentServer
import grpc
from concurrent import futures
import sentiment_pb2
from sentiment_pb2_grpc import add_SentimentAnalysisServicer_to_server 
from grpc_reflection.v1alpha import reflection 
import signal
import sys
from grpc_health.v1 import health, health_pb2_grpc, health_pb2

def grpc_server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_SentimentAnalysisServicer_to_server(SentimentServer(), server)

    SERVICE_NAMES = (
        sentiment_pb2.DESCRIPTOR.services_by_name["SentimentAnalysis"].full_name,
        reflection.SERVICE_NAME,  # Required for reflection to work
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

      # Register the health check service
    health_servicer = health.HealthServicer()
    health_pb2_grpc.add_HealthServicer_to_server(health_servicer, server)
    
    # Set health status to SERVING
    health_servicer.set("nlp-service", health_pb2.HealthCheckResponse.SERVING)

    server.add_insecure_port("0.0.0.0:5000")
    server.start()
    print("NLP Service running on port 0.0.0.0:5000")

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