#Get the GCC preinstalled image from Docker Hub
FROM gcc:4.9

#Copy the current folder which contains C++ source code to the Docker image under /usr/src
COPY . /usr/app

#Install OpenCV
# RUN git clone https:

#Use G++ to compile the cpp source file 
RUN g++ -std=c++14 -o EXEC /usr/app/src/main.cpp

# Command to execute when the image loads
CMD ["/EXEC"]