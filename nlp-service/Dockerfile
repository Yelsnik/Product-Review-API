# Builder Stage
FROM python:3.9-slim AS builder

WORKDIR /app

# Install dependencies for building Python packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    python3-venv python3-dev gcc libc-dev

# Create virtual environment
RUN python3 -m venv /opt/venv

# Ensure venv's pip is installed correctly
RUN /opt/venv/bin/python3 -m ensurepip && \
    /opt/venv/bin/pip install --no-cache-dir --upgrade pip setuptools wheel

# Copy and install dependencies
COPY requirements.txt .
RUN /opt/venv/bin/pip install --no-cache-dir -r requirements.txt

# Final Stage
FROM python:3.9-slim
WORKDIR /app

# Copy virtual environment
COPY --from=builder /opt/venv /opt/venv

# Ensure venv is used
ENV PATH="/opt/venv/bin:$PATH"

# Copy application files
COPY app.env .
COPY . .

EXPOSE 5000

CMD ["python3", "main.py"]