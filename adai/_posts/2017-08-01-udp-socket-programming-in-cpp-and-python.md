---
title: "UDP socket programming in C++ and Python"
tags: 
  - tcp
  - udp
  - socket
  - cpp
  - python
categories:
  - programming     
---

Sample C++ code for transmitting data via UDP socket.

```c++
// C++ UDP Transmitter

#include "Network.h"
#include "iostream"

#pragma once

int main()
{
  std::string IP = "127.0.0.1";
  int PORT = 8888;

  try
  {
    WSASession Session;
    UDPSocket Socket;
    std::string data = "hello world";
    char buffer[100];

    while (1)
    {
      std::cout << "Enter data to transmit : " << std::endl;
      std::getline(std::cin, data);
      Socket.SendTo(IP, PORT, data.c_str(), data.size());
    }
  }
  catch (std::exception &ex)
  {
    std::cout << ex.what();
  }
}
```

Sample C++ code for receiving data via UDP socket.

```c++
// C++ UDP Receiver

#include "Network.h"
#include "iostream"

#pragma once

int main()
{
  int PORT = 8888;

  try
  {
    WSASession Session;
    UDPSocket Socket;
    char buffer[100];

    Socket.Bind(PORT);
    while (1)
    {
      sockaddr_in add = Socket.RecvFrom(buffer, sizeof(buffer));
      std::string input(buffer);
      std::cout << "Data received: " << input << std::endl;
    }
  }
  catch (std::system_error& e)
  {
    std::cout << e.what();
  }
}
```

Include the following code as `Network.h` header file, which is required for the C++ UDP streaming above.

```c++
// Network.h header file for C++ UDP

#define _WINSOCK_DEPRECATED_NO_WARNINGS
#include <system_error>
#include <string>
#include <iostream>

#pragma once
#pragma comment (lib, "ws2_32")

class WSASession
{
  public:
    WSASession()
    {
      int ret = WSAStartup(MAKEWORD(2, 2), &data);
      if (ret != 0)
        throw std::system_error(WSAGetLastError(), std::system_category(), "WSAStartup Failed");
    }
    ~WSASession()
    {
      WSACleanup();
    }

  private:
    WSAData data;
};

class UDPSocket
{
  public:
    UDPSocket()
    {
      sock = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
      if (sock == INVALID_SOCKET)
        throw std::system_error(WSAGetLastError(), std::system_category(), "Error opening socket");
    }
    ~UDPSocket()
    {
      closesocket(sock);
    }

    void SendTo(const std::string& address, unsigned short port, const char* buffer, int len, int flags = 0)
    {
      sockaddr_in add;
      add.sin_family = AF_INET;
      add.sin_addr.s_addr = inet_addr(address.c_str());
      //add.sin_addr.s_addr = inet_pton(AF_INET, "127.0.0.1", address.c_str());
      add.sin_port = htons(port);
      int ret = sendto(sock, buffer, len, flags, reinterpret_cast<SOCKADDR *>(&add), sizeof(add));
      if (ret < 0)
        throw std::system_error(WSAGetLastError(), std::system_category(), "sendto failed");
    }
    void SendTo(sockaddr_in& address, const char* buffer, int len, int flags = 0)
    {
      int ret = sendto(sock, buffer, len, flags, reinterpret_cast<SOCKADDR *>(&address), sizeof(address));
      if (ret < 0)
        throw std::system_error(WSAGetLastError(), std::system_category(), "sendto failed");
    }
    sockaddr_in RecvFrom(char* buffer, int len, int flags = 0)
    {
      sockaddr_in from;
      int size = sizeof(from);
      int ret = recvfrom(sock, buffer, len, flags, reinterpret_cast<SOCKADDR *>(&from), &size);
      if (ret < 0)
        throw std::system_error(WSAGetLastError(), std::system_category(), "recvfrom failed");

      // make the buffer zero terminated
      buffer[ret] = 0;
      return from;
    }
    void Bind(unsigned short port)
    {
      sockaddr_in add;
      add.sin_family = AF_INET;
      add.sin_addr.s_addr = htonl(INADDR_ANY);
      add.sin_port = htons(port);

      int ret = bind(sock, reinterpret_cast<SOCKADDR *>(&add), sizeof(add));
      if (ret < 0)
        throw std::system_error(WSAGetLastError(), std::system_category(), "Bind failed");
    }

  private:
    SOCKET sock;
};
```

Sample Python code for receiving data via UDP socket.

```python
# Python UDP Receiver

import socket
BUFFER_LEN = 100 #in bytes

def initUDP(IP, port):
    #Create a datagram socket
    sock = socket.socket(socket.AF_INET, # Internet
                         socket.SOCK_DGRAM) # UDP
    #Enable immediate reuse of IP address
    sock.setsockopt(socket.SOL_SOCKET,socket.SO_REUSEADDR,1)
    #Bind the socket to the port
    sock.bind((IP, port))
    #Set a timeout so the socket does not block indefinitely when trying to receive data
    sock.settimeout(0.5)

    return sock

def readUDP(sock):
    try:
        data, addr = sock.recvfrom(BUFFER_LEN)
    except socket.timeout as e:
        return b'Error'
    except Exception as e:
        return b'Error'
    else:
        return data
```