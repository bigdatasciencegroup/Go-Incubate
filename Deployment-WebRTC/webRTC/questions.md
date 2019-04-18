1. will the peerconnection.ontrack() function be triggered after we have copied the answer from line 110, main.go, into the jsfiddle.net browser and clicked start session?

2. When will the peerconection.ontrack() function be called?

3. Why are we using a chanel, namely localTrackChan, instead of a normal variable, to transfer the Track from line 75 to line 112 of main.go ? Is this to block the program on line 112, main.go, to ensure that we have a broadcaster (i.e., a browser which is uploading video to the server) before there is any client to download the video from the server?

4. will the program work still work fine if there are downloaders join the server first before any video uploaders are avaiblae?

5. In the following statement: peerConnection.SetRemoteDescription(offer) , is there communication from the server to the remote browser taking place or is it simply setting local variables in the server about the remote browser?

6. At line 126, main.go, we add localtrack to peerconnection. So where is the server code which is responsible to reading the local video from server and sending it to the client? Or is it that the client pulls the video from the server using its javascript?  