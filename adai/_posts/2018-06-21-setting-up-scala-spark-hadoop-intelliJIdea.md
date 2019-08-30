---
title: "Setting up Scala sbt, Spark, Hadoop and IntelliJIdea in Windows"
tags: 
  - spark
  - scala 
categories:
  - spark
  - scala   
last_modified_at: 2018-06-21
---


We briefly describe the steps to setup Scala sbt, Spark, Hadoop, and IntelliJ Idea in Windows computer.

First, please follow the [tools setup](https://github.com/Adaickalavan/Functional-Programming-in-Scala-Specialization-EPFL-Coursera/blob/master/Tools%20setup.pdf) instructions to install the necessary JDK and Scala sbt. Then, follow this [IntelliJIdea tutorial](https://github.com/Adaickalavan/Functional-Programming-in-Scala-Specialization-EPFL-Coursera/blob/master/IntelliJ%20Idea%20tutorial.pdf) to setup the IDE.

Next, we focus on Spark and Hadoop.
1. Download the latest version of Spark from [Spark website](https://spark.apache.org/downloads.html). 
  * Choose a Spark release: choose your desired version (I chose `spark-2.2.1`)  
  * Choose a package type: select `Pre-built for Apache Hadoop 2.7 and later`
  * Then, click on download `spark-x.x.x-bin-hadoop2.7.tgz`

1. Unzip the Spark into your C drive, for example unzip to: `C:\spark-2.2.1-bin-hadoop2.7\` 

1. Download the entire winutils distribution from this [repository](https://github.com/steveloughran/winutils). Form the downloaded repo, copy the folder corresponding to your installed Hadoop version to a location on your C drive. In my case I copied `hadoop-2.7.1` folder to `C:\`. Then, rename the `hadoop-x.x.x` folder to `winutil`.

1. Go to your Environment Variables setting and set the following `System Variables`:

    | Variable    | Value                             | 
    |:-----------:|:---------------------------------:|
    | HADOOP_HOME | C:\winutil\                       |
    | SPARK_HOME  | C:\spark-2.2.1-bin-hadoop2.7\     |
    | Path        | C:\spark-2.2.1-bin-hadoop2.7\bin  |
    | Path        | C:\winutil\bin                    |

To verify installation of Spark, open a comannd prompt and navigate to your SPARK_HOME directory, for example go to `C:\spark-2.2.1-bin-hadoop2.7\`. Run `bin\spark-shell` which should start up Scala console with a Spark session as `sc`. Exit the shell via `:quit` command.

![spark](/assets/images/scala_c4_w1_a01_01.jpg){:height="150%" width="95%" .align-center}

That is it, we are done. Yay!

