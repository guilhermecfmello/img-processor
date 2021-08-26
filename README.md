# img-processor
Test project to build a containerized image processor orchestrated by Kubernetes

# How to run
  - Build Docker image with following command
    - `docker build --tag img-processor .`
  - Run Docker image exposing port and redirecting to port 8000 internaly
    - `docker run -p 8000:8000 img-processor`
