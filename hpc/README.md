# Parallel programming with OpenMP & C/C++

We provide the code solutions to C/C++ and OpenMP coding exercises from the Intro to OpenMP by Tim Mattson, Intel Corp, online course at Youtube. Watch the course [video](https://www.youtube.com/watch?v=nE-xN4Bf8XI&list=PLLX-Q6B8xqZ8n8bwjGdzBJ25X2utwnoEG) and read the course [notes](https://www.openmp.org/wp-content/uploads/Intro_To_OpenMP_Mattson.pdf).

For easy reference, each code solution is named according to the slide on which they appear in the course [notes](https://www.openmp.org/wp-content/uploads/Intro_To_OpenMP_Mattson.pdf).


**List of exercises**
1. Slide 38 - Parallel 'Hello World' program
1. Slide 48 - Compute Pi in Serial
1. Slide 52 - Compute Pi in Parallel with false sharing
1. Slide 57 - Compute Pi in Parallel with padding
1. Slide 69 - Compute Pi in Parallel with synchronisation
1. Slide 88 - Compute Pi with WorkSharing
    + main file: `apps/tutorial.c`
    + to execute, set: `command: /src/bin/tutorial` in `docker-compose.yaml` file
1. Slide 119 - Mandel Brot
    + main file: `apps/mandel.c`
    + to execute, set: `command: /src/bin/mandel` in `docker-compose.yaml` file    
1. Slide 124 - Linked list computed serially
1. Slide 128 - Linked list without Tasks
1. Slide 143 - Linked list with Tasks
    + main file: `apps/linked.c`
    + to execute, set: `command: /src/bin/linked` in `docker-compose.yaml` file
1. Slide 166 - Producer Consumer
    + main file: `apps/prodCons.c`
    + to execute, set: `command: /src/bin/prodCons` in `docker-compose.yaml` file
1. Slide 
