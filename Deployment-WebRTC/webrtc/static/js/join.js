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
pc.ontrack = handleTrack;

// Set the handler for ICE connection state
// This will notify you when the peer has connected/disconnected
function handleICEConnectionStateChange(event){
  log("ICEConnectionStateChange: "+pc.iceConnectionState)
}

function handleICEGatheringStateChange(event){
  log("ICEGatheringStateChange: "+pc.iceGatheringState)
}

function handleSignalingStateChange(event){
  log("SignalingStateChange: "+pc.signalingState)
}

function handleICECandidate(event) {
  log("ICECandidate: "+event.candidate)
  if (event.candidate === null) {
    document.getElementById('finalLocalSessionDescription').value = JSON.stringify(pc.localDescription)
  }
}

function handleTrack(event){
  var el = document.getElementById('video1')
  el.srcObject = event.streams[0]
  el.autoplay = true
  el.controls = true
}

function createOffer(){
  return pc.createOffer()
  .then(offer => pc.setLocalDescription(offer))
  .then(() => {
    document.getElementById('localSessionDescription').value = JSON.stringify(pc.localDescription);
  })
}

function sendToServer(url, msg){
  return fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'text/plain; charset=utf-8'
    },
    body: msg
  })
  .then(response => {
    // Verify HTTP-status is 200-299
    if (response.ok){ 
      if (response.headers.get('Content-Type') == "application/json; charset=utf-8") {
        return response.json();
      } else {
        throw new Error("Content-Type expected `application/json; charset=utf-8` but got "+response.headers.get('Content-Type'))
      }
    } else {
      throw new HttpError(response);
    }
  })
  .then(json => { 
    document.getElementById('remoteSessionDescription').value = JSON.stringify(json.SD)
    return json.SD
  })
}

pc.addTransceiver('video', {'direction': 'recvonly'})

createOffer()
  .then(() => {
    let myUsername = "Client";
    let msg = {
      Name: myUsername,
      SD: pc.localDescription
    };
    return sendToServer("/sdp", JSON.stringify(msg))
  })
  .then(sdp => {
    pc.setRemoteDescription(sdp)
  })
  .catch(log)
