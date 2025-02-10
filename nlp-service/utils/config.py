from dotenv import load_dotenv
import os
from pathlib import Path

dotenv_path = Path('app.env')
load_dotenv(dotenv_path=dotenv_path)

grpc_address = os.getenv("GRPC_SERVER")