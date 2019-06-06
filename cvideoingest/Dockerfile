# Base image
FROM alpine AS build

# Install build tools
RUN apk update && apk add --no-cache \ 
    autoconf build-base binutils cmake make file g++ git libgcc libtool linux-headers make musl-dev ninja pkgconfig

# Change working directory
WORKDIR /home

# Copy the current folder which contains C++ source code to the Docker image under /usr/src
COPY . /home/

# Use cmake to compile the cpp source file 
RUN cmake -E make_directory build && \
    cmake -E chdir build cmake .. && \
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