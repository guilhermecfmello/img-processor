# img-processor
Test project to build a containerized image processor orchestrated by Kubernetes

# How to run
1- Build Docker image
  - `docker build --tag img-processor .`
2- Run Docker image exposing port and redirecting to port 8000 internaly
  - `docker run -p 8000:8000 img-processor`
