# Base image
FROM ubuntu AS build

# Install build tools
RUN apt-get update && apt-get -y install \
    g++ cmake git pkg-config

# Install gstreamer
RUN apt-get -y install \
    libgstreamer1.0-0 gstreamer1.0-plugins-base         gstreamer1.0-plugins-good gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly gstreamer1.0-libav gstreamer1.0-doc gstreamer1.0-tools gstreamer1.0-x gstreamer1.0-alsa gstreamer1.0-gl gstreamer1.0-gtk3 gstreamer1.0-qt5 gstreamer1.0-pulseaudio

# Change working directory
WORKDIR /home

# Copy the current folder which contains C++ source code to the Docker image
COPY . /home/

# Use cmake to compile the cpp source file 
RUN cmake -E make_directory build && \
    cmake -E chdir build cmake -DCMAKE_BUILD_TYPE=Release .. && \
    cmake --build build

# Multistage build
FROM alpine AS runtime

# Label
LABEL author Adaickalavan

# Install runtime tools
RUN apk update && apk add --no-cache \ 
    libstdc++

# Copy needed source
COPY --from=build /home/build/app /home/build/app

# Command to execute when the image loads
CMD ["/home/build/app"]