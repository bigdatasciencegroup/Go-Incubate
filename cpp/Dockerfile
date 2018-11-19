# #Get the base Ubuntu image from Docker Hub
# FROM ubuntu:latest

# #Update apps on the base image
# RUN apt-get -y update && apt-get install -y

# #Install the Clang compiler
# RUN apt-get -y install clang

#Get the GCC preinstalled image from Docker Hub
FROM gcc:4.9

#Copy the current folder which contains C++ source code to the Docker image under /usr/src
COPY . /usr/src/app

# #Use Clang to compile the cpp source file 
# RUN clang++ -0 EXE /usr/src/app/main.cpp

#Use GCC to compile the cpp source file 
RUN g++ -o EXEC /usr/src/app/main.cpp

# Command to execute when the image loads
CMD ["/EXEC"]