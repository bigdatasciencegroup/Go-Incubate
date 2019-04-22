The following questions are all pertaining to the main.go file in sfu-minimal code example at pion/webrtc. Refer here: https://github.com/pion/webrtc/blob/master/examples/sfu-minimal/main.go 

1. Will the `peerConnection.OnTrack()` function on line 58 be triggered after we have copied the `answer` from line 110 (`fmt.Println(signal.Encode(answer))`) into the jsfiddle.net browser and clicked start session in the browser? In other words precisely when will the `OnTrack` function be called?

2. Will the `peerConection.OnTrack()` function be called whenever 
- a video uploader joins the server?
- a video downloading browser joins the server?

3. To whom does the `receiver *webrtc.RTPReceiver` function argument refer to in line 58. Does it refer to the server or the remote browser?

4. Why are we using a channel, namely `localTrackChan`, instead of a normal variable, to transfer the `localTrack` info from line 75 to line 112 of main.go? Is this to block the program on line 112, main.go, to ensure that we have a broadcaster (i.e., a browser which is uploading video to the server) before there is any client to download the video from the server?

5. Will the program still work fine if browser downloaders join the server first before any video uploaders are available?

6. In the following statement: `peerConnection.SetRemoteDescription(offer)` , is there communication from the server to the remote browser taking place or is it simply setting local variables in the server about the remote browser?

7. At line 126, main.go, we add localtrack to peerconnection. So where is the server code which is responsible to reading the local video from server and sending it to the client? Or is it that the client pulls the video from the server using its javascript?  


The following are questions from other webrtc library.

8.  (refer here: https://github.com/pion/webrtc/blob/3d2c1c2c32c96b5124c43ad85390cf1fb0961924/track.go#L81) In `func (t *Track)Read(b []byte) (n int, err error)`, why when number of `len(t.activeSenders) != 0`, it implies a local track?

9. 