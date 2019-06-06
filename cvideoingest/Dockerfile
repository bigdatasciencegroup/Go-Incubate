#Get the GCC preinstalled image from Docker Hub
FROM gcc

#Install build tools
RUN apt-get update && \
    apt-get -y install g++ \
    git cmake pkg-config

#Copy the current folder which contains C++ source code to the Docker image under /usr/src
COPY . /home/

# Change working directory
WORKDIR /home

#Use G++ to compile the cpp source file 
RUN cmake -E make_directory build && \
    cmake -E chdir build cmake .. && \
    cmake --build build

# Command to execute when the image loads
CMD ["/home/build/app"]