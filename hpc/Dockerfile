# Base image
FROM ubuntu AS build

# Install build tools
RUN apt-get update && apt-get -y install \
    g++ cmake git pkg-config

# Change working directory
WORKDIR /src

# Copy the current folder which contains C++ source code to the Docker image
COPY . /src/

# Use cmake to compile the cpp source file 
RUN cmake -E make_directory build && \
    cmake -E chdir ./build cmake -DCMAKE_BUILD_TYPE=Release .. && \
    cmake --build ./build

# Check if the program has dynamic dependencies
RUN cd /src/build/bin && ldd app

# # Multistage build
# FROM alpine AS runtime

# # Label
# LABEL author Adaickalavan

# # Copy needed source
# COPY --from=build /home/build/bin /home/build/bin
# COPY --from=build /home/build/lib /home/build/lib

# Command to execute when the image loads
ENTRYPOINT ["/src/build/bin/app"]