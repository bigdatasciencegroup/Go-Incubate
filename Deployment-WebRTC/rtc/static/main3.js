function hasUserMedia() {
    navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;
    return !!navigator.getUserMedia;
}
  
function hasRTCPeerConnection() {
    window.RTCPeerConnection = window.RTCPeerConnection || window.webkitRTCPeerConnection || window.mozRTCPeerConnection;
    return !!window.RTCPeerConnection;
}
  
var yourVideo = document.querySelector('#yours'),
    theirVideo = document.querySelector('#theirs'),
    yourConnection, theirConnection;

// Put variables in global scope to make them available to the browser console.
const constraints = window.constraints = {
    audio: true,
    video: true,
};    

async function startMedia(){
    if (hasUserMedia()) {
        var stream = await navigator.mediaDevices.getUserMedia(constraints)
        window.stream = stream;
        yourVideo.srcObject = stream;

        if (hasRTCPeerConnection()) {
            startPeerConnection(stream);
        } else {
            alert("Sorry, your browser does not support WebRTC.");
        }
    } else {
        alert("Sorry, your browser does not support WebRTC.");
    }
}

startMedia()
  
function startPeerConnection(stream) {
    var configuration = {
        // "iceServers": [{ "url": "stun:127.0.0.1:9876" }]
    };
    yourConnection = new webkitRTCPeerConnection(configuration);
    theirConnection = new webkitRTCPeerConnection(configuration);

    // Setup stream listening
    yourConnection.addTrack(stream);
    theirConnection.onaddstream = function (e) {
        theirVideo.src = window.URL.createObjectURL(e.stream);
    };
  
    // Setup ice handling
    yourConnection.onicecandidate = function (event) {
        if (event.candidate) {
            theirConnection.addIceCandidate(new RTCIceCandidate(event.candidate));
        }
    };
  
    theirConnection.onicecandidate = function (event) {
        if (event.candidate) {
            yourConnection.addIceCandidate(new RTCIceCandidate(event.candidate));
        }
    };
  
    // Begin the offer
    yourConnection.createOffer(function (offer) {
        yourConnection.setLocalDescription(offer);
        theirConnection.setRemoteDescription(offer);
  
        theirConnection.createAnswer(function (offer) {
            theirConnection.setLocalDescription(offer);
            yourConnection.setRemoteDescription(offer);
        });
    });

    alert("inside hausermedia rtc")
};
  