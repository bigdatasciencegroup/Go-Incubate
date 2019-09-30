# Parallel programming with OpenMP & C/C++

We provide the code solutions to C/C++ and OpenMP coding exercises from the Intro to OpenMP by Tim Mattson, Intel Corp, online course at Youtube. Watch the course [video](https://www.youtube.com/watch?v=nE-xN4Bf8XI&list=PLLX-Q6B8xqZ8n8bwjGdzBJ25X2utwnoEG) and read the course [notes](https://www.openmp.org/wp-content/uploads/Intro_To_OpenMP_Mattson.pdf).

For easy reference, each code solution is named according to the slide on which they appear in the course [notes](https://www.openmp.org/wp-content/uploads/Intro_To_OpenMP_Mattson.pdf).


**List of exercises**

In order to run the solutions for the exercise, the following steps

<style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;border-color:#aaa;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#aaa;color:#333;background-color:#fff;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:#aaa;color:#fff;background-color:#f38630;}
.tg .tg-0z1f{background-color:#FCFBE3;border-color:#000000;text-align:left;vertical-align:top}
.tg .tg-wp8o{border-color:#000000;text-align:center;vertical-align:top}
.tg .tg-kvc3{background-color:#FCFBE3;border-color:#000000;text-align:left;vertical-align:middle}
.tg .tg-73oq{border-color:#000000;text-align:left;vertical-align:top}
.tg .tg-0a7q{border-color:#000000;text-align:left;vertical-align:middle}
</style>
<table class="tg">
  <tr>
    <th class="tg-wp8o">Exercise</th>
    <th class="tg-wp8o">Instructions to run</th>
  </tr>
  <tr>
    <td class="tg-0z1f">Slide 38 - Parallel 'Hello World' program</td>
    <td class="tg-kvc3" rowspan="6">1) Main file: `/apps/tutorial.cpp`<br><br>2) To execute, set: `command: /src/bin/tutorial` in `docker-compose.yaml` file<br></td>
  </tr>
  <tr>
    <td class="tg-73oq">Slide 48 - Compute Pi in serial</td>
  </tr>
  <tr>
    <td class="tg-0z1f">Slide 52 - Compute Pi in parallel with false sharing</td>
  </tr>
  <tr>
    <td class="tg-73oq">Slide 57 - Compute Pi in parallel with padding</td>
  </tr>
  <tr>
    <td class="tg-0z1f">Slide 69 - Compute Pi in parallel with synchronization</td>
  </tr>
  <tr>
    <td class="tg-73oq">Slide 88 - Compute Pi with WorkSharing</td>
  </tr>
  <tr>
    <td class="tg-kvc3">Slide 119 - Mandel Brot</td>
    <td class="tg-kvc3">1) Main file: `/apps/mandel.c`<br><br>2) To execute, set: `command: /src/bin/mandel` in `docker-compose.yaml` file</td>
  </tr>
  <tr>
    <td class="tg-73oq">Slide 124 - Linked list computed serially</td>
    <td class="tg-0a7q" rowspan="3">1) Main file: `/apps/linked.c`<br><br>2) To execute, set: `command: /src/bin/linked` in `docker-compose.yaml` file</td>
  </tr>
  <tr>
    <td class="tg-0z1f">Slide 128 - Linked list in parallel without Tasks</td>
  </tr>
  <tr>
    <td class="tg-73oq">Slide 143 - Linked list in parallel with Tasks</td>
  </tr>
  <tr>
    <td class="tg-kvc3">Slide 166 - Producer Consumer</td>
    <td class="tg-0z1f">1) Main file: `/apps/prodCons.c`<br><br>2) To execute, set: `command: /src/bin/prodCons` in `docker-compose.yaml` file</td>
  </tr>
  <tr>
    <td class="tg-73oq">Slide 177 - Parallel Monte Carlo computation of PI</td>
    <td class="tg-0a7q" rowspan="3">1) Main file: `/apps/pi_mc.c`<br><br>2) To execute, set: `command: /src/bin/pi_mc` in `docker-compose.yaml` file</td>
  </tr>
  <tr>
    <td class="tg-0z1f">Slide 181 - Thread safe random number generator</td>
  </tr>
  <tr>
    <td class="tg-73oq">Slide 185 - Leap frog method to avoid overlapped random number sequence</td>
  </tr>
</table>    
 
