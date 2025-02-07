# Use an official Python image as a base
FROM python:3.10

# Set the working directory in the container
WORKDIR /rnd

# Create a virtual environment
RUN python3.10 -m venv venv

# Activate the virtual environment for subsequent commands
# Install the dependencies
RUN venv/bin/pip install rvc-python
RUN venv/bin/pip install torch==2.1.1+cu118 torchaudio==2.1.1+cu118 --index-url https://download.pytorch.org/whl/cu118

# Copy the contents
COPY . .
