// this code works the modern way
"use strict";

/* eslint-env browser */
var log = msg => {
  document.getElementById('logs').innerHTML += msg + '<br>'
}

var mediaConstraints = {
  audio: false, // We dont want an audio track
  video: true // ...and we want a video track
};

let pc = new RTCPeerConnection({
  iceServers: [
    {
      urls: ['stun:stun.l.google.com:19302','stun:stun.stunprotocol.org'],
    }
  ]
});

pc.oniceconnectionstatechange = handleICEConnectionStateChange;
pc.onicegatheringstatechange = handleICEGatheringStateChange;
pc.onsignalingstatechange = handleSignalingStateChange;
pc.onicecandidate = handleICECandidate;
// pc.onnegotiationneeded = handleNegotiationNeeded;
// pc.ontrack = handleTrack;


function startMedia(){
  return navigator.mediaDevices.getUserMedia(mediaConstraints)
    .then(function(stream){
      document.getElementById("video1").srcObject = stream;
      stream.getTracks().forEach(track => pc.addTrack(track, stream));
    })
    .catch(handleGetUserMediaError);
}

function handleICEConnectionStateChange(event){
  log(pc.iceConnectionState)
};

function handleICEGatheringStateChange(event){
  log(pc.iceGatheringState)
};

function handleSignalingStateChange(event){
  log(pc.signalingState)
};

function handleICECandidate(event) {
  if (event.candidate === null) {
    document.getElementById('localSessionDescription').value = btoa(JSON.stringify(pc.localDescription))
  }
};


function createOffer(){
  return pc.createOffer()
    .then(offer => {
      pc.setLocalDescription(offer)
      alert(offer)
    })
    .then(function(){
      let sdp = btoa(JSON.stringify(pc.localDescription));
      document.getElementById('localSessionDescription').value = localSDP
      return sdp
    })
    .catch(log)
}

function send(){
  let offer = "hi yes yes from browser"
  sendToServer("a")
  .then(
    alert("12")
  )
  .then(() => {
    sendToServer("b")
  })

function fetchSDP(){
  return fetch("/sdp", {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json; charset=utf-8'
    },
    body: JSON.stringify(offer)
  })
  .then(response => {
    if (response.ok){ // if HTTP-status is 200-299
      if (response.headers.get('Content-Type') == "application/json; charset=utf-8") {
        return response.json();
      } else {
        throw new Error("Content-Type expected `application/json; charset=utf-8` but got "+response.headers.get('Content-Type'))
      }
    } else {
      throw new HttpError(response);
    }
  }
  let text = await then(text => { 
    alert(text.Result);
    alert(text.ServerSDP);
  })
  .then(function(){ 
    let sd = document.getElementById('remoteSessionDescription').value
    if (sd === '') {
      return alert('Session Description must not be empty')
    }
  })
  .catch(log)
}

startMedia()

function handleNegotiationNeeded(){

}

  // let sd = document.getElementById('remoteSessionDescription').value
  // if (sd === '') {
  //   return alert('Session Description must not be empty')
  // }

  // try {
  //   pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd))))
  // } catch (e) {
  //   alert(e)
  // }
// }

// This handler for the track event is called by the local WebRTC layer when a 
// track is added to the connection. 
function handleTrackEvent(event) {
  document.getElementById("video1").srcObject = event.streams[0];
}

function handleGetUserMediaError(e) {
  switch(e.name) {
    case "NotFoundError":
      log("Unable to open your call because no camera and/or microphone" +
            "were found.");
      break;
    case "SecurityError":
    case "PermissionDeniedError":
      // Do nothing; this is the same as the user canceling the call.
      break;
    default:
      log("Error opening your camera and/or microphone: " + e.message);
      break;
  }
}

class HttpError extends Error { // (1)
  constructor(response) {
    super(`${response.status} for ${response.url}`);
    this.name = 'HttpError';
    this.response = response;
  }
}

window.addEventListener('unhandledrejection', function(event) {
  // the event object has two special properties:
  alert(event.promise); // [object Promise] - the promise that generated the error
  alert(event.reason); // Error: Whoops! - the unhandled error object
});
