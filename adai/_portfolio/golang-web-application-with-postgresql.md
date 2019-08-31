---
title: "Golang Web Application with PostgreSQL"
excerpt: "Golang, PostgreSQL, Javascript"
mathjax: true
header:
  teaser: assets/images/goweb_01.jpg
---

{% include feature_row %}

<!-- {% include toc %} -->

## Project

### Problem
In this project we will develop a Golang based web application to display and edit a database of people's name, birthday, and occupation. Users may enter new entries via a web form. New entries will be saved to the PostgreSQL database. Simultaneously, the table on the webpage will be updated with the new entry. 

Our example will demonstrate the following:
* Running a web server based application in Go
* Web routing in Go
* Running a PostgreSQL database server
* Creating and editing database in PostgreSQL
* Interactions among: web browser $$\longleftrightarrow$$ Go web server $$\longleftrightarrow$$ PostgreSQL database 

The following will be used:
* Golang (programming language)
* SQL (query language)
* PostgreSQL (Database)
* VS Code (integrated development environment)

It is recommended to use [VS Code](https://code.visualstudio.com/download) as the IDE since it provides excellent support for GO, github integration, and project file structure management capability.

The expected end product of this project is shown below:

![pic1](/assets/images/goweb_01.jpg){:height="150%" width="95%" .align-center}

### Repository
The [repository](https://github.com/Adaickalavan/GoWeb-PostgreSQL) consists of the following:  
* Golang, PostgreSQL, HTML, HTML DOM, CSS code - complete functioning sample code

The project structure is as follows:
```
GoWeb-PostgreSQL          # main folder
├── static                # contains static files of the project, such as HTML files
│   └── index.html        # primary webpage at root "/"
├── main.go               # main file of Go code, initializes web router and sets up database connection
├── personHandler.go      # handles get and insert of new objects into database
└── store.go              # interface between PostgreSQL database and web servers
```

### Solution 
The code provided in the repository is well commented, and when read together with the project structure listed above, it becomes self explanatory. Nevertheless, key points are covered below in sequence.

1.  Please download two additional packages `github.com/gorilla/mux` and `github.com/lib/pq` using `go get github.com/*packagename*`. `mux` is used for web routing and `pq` serves as the Postgres driver for the sql-package.  

1. `func newRouter` defines the web routing for page handlers and for serving static files (i.e., index.html). Note that the static files are placed inside `./static/` folder in our local directory, but it is served at the root (i.e., http://localhost:8080/) when browsed in a browser.

1. `index.html` file defines the main webpage, consisting of a table, form, and a HTML DOM script. The HTML DOM script runs on each page load and updates the table by querying the database. Upon submitting the web form, the new entry is written to the database and also reflected in the table.

1. `personHandler.go` defines the `struct` to hold each individuals property values. Additionally, it defines the handlers which retrieves data from and writes data to the database, according to the http requests made. The communication between page handlers and `index.html` is transmitted in JSON format.

1. `store.go` acts as the interface between page handlers' requests and our PostgreSQL database. `store.go` executes SQL queries to retrieve or to enter data into the database.

1. Download and install the PostgreSQL open source database from [postgresql.org](https://www.postgresql.org/download/) 

1. pgAdmin4 is equipped with in-built SQL query tool which can be found at: Tools --> Query Tool

1. Refer to the [instructions pdf](https://github.com/Adaickalavan/GoWeb-PostgreSQL/blob/master/Spatial_Database_Lab1_PostgreSQL%20Tutorial%20I_GUI.pdf) for step-by-step instruction on adding, editing entries, and querying a database in PostgreSQL. Although the instructions are written for pgAdmin3, the steps remain the same in pgAdmin4 (with some user interface improvements).

1. Steps to create and populate a databse in PostgreSQL using the pgAdmin4 graphical user interface is surmised below.
    + (a) **Create a database and a table** Create a database = Databases $$\rightarrow$$
 Create $$\rightarrow$$ Database $$\rightarrow$$ `peopleDatabase`. Create a table = `peopleDatabase` $$\rightarrow$$ Schemas $$\rightarrow$$ Public $$\rightarrow$$ Tables $$\rightarrow$$ Create $$\rightarrow$$ Table $$\rightarrow$$ `peopleinfo`. Use all small letters when naming a table inside a database.
    + (b) **Specify column headers** `peopleinfo` $$\rightarrow$$ Properties $$\rightarrow$$ Columns. Add the relevant columns with their corresponding settings as shown in the figure below.
    + (c) **SQL insert** To enter new values programatically into the table, execute the following code in the SQL query tool: 
    ```sql
    INSERT INTO peopleinfo(nama,birthday,occupation) VALUES
    ('Elephant', '27 Nov 1994', 'Accountant'),
    ('Winnie the Pooh Bear', '5 Oct 1994', 'Farmer');
    ```
    Note that table values can be edited using the GUI by simply clicking on the rows of the table.
    + (d) **Result** `peopleinfo` $$\rightarrow$$ View/Edit Data $$\rightarrow$$ All Rows. Final expected table in our database is displayed.

    | (a) Create a database and a table | (b) Specify column headers|
    |:---------------------------------:|:-------------------------:|
    |[![](/assets/images/postgresql_01.jpg)](/assets/images/postgresql_01.jpg) | [![](/assets/images/postgresql_02.jpg)](/assets/images/postgresql_02.jpg)|

    | (c) SQL insert | (d) Result|
    |:--------------:|:---------:|
    |[![](/assets/images/postgresql_03.jpg)](/assets/images/postgresql_03.jpg) | [![](/assets/images/postgresql_04.jpg)](/assets/images/postgresql_04.jpg)|

1. To have our Go code connect to the PostgreSQL database, setup the connection string as follows and execute the `sql.Open("postgres", connString)` command. 
    ```go
    connString := `user=postgres           // username to access the database 
                   password=1234           // password to access the database 
                   host=localhost          // default host of PostgreSQL
                   port=5432               // default port of PostgreSQL
                   dbname=peopleDatabase   // name of PostgreSQL database
                   sslmode=disable`
    ```
    A good tutorial for connecting Golang to PostgreSQL database is available at [calhoun.io](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/) 

1. To run the application, navigate to the project directory 
    ```
    $GOPATH/src/github.com/adaickalavan/GoWeb-PostgreSQL
    ``` 
    and enter the command: 
    ```
    $ go run main.go personHandler.go store.go
    ```
    Alternatively, you may run the .exe file created through `go install` in the above directory.

1. For those wishing to run the Go web application without a database, a slighlty modified code is available at [GoWeb-inMemory](https://github.com/Adaickalavan/GoWeb-inMemory). Here, instead of saving to a database, the information is saved to a variable (i.e., `var personList []Person`) held in memory by the Go web server. Any data added to the slice variable, via the web form, will be erased when the server restarts.

1. Finally go to `localhost:8080` in your browser to see the web application in action.
