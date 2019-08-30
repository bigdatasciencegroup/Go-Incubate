---
title: "Linux"
permalink: /docs/linux/
toc: true
---

## Configuration
+ Need to set proxy for Linux `apt` if we use `apt` behind a corporate proxy, e.g., `http://10.0.0.0:8080/`. Hence, create a `apt.conf` file at `/etc/apt/` such that `/etc/apt/apt.conf` file contains the following.
    ```
    Acquire::http::proxy "http://10.0.0.0:8080/";
    Acquire::https::proxy "https://10.0.0.0:8080/";
    ```    