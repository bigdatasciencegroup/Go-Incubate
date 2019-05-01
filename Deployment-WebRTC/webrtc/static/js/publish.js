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
// pc.onicecandidate = handleICECandidate;
pc.onnegotiationneeded = handleNegotiationNeeded;
// pc.ontrack = handleTrack;


async function startMedia(){
  try {
    const stream = await navigator.mediaDevices.getUserMedia(mediaConstraints);
    document.getElementById("video1").srcObject = stream;
    stream.getTracks().forEach(track => pc.addTrack(track, stream));
  }
  catch (e) {
    return handleGetUserMediaError(e);
  }
}

// Set the handler for ICE connection state
// This will notify you when the peer has connected/disconnected
function handleICEConnectionStateChange(event){
  log("ICEConnectionStateChange: "+pc.iceConnectionState)
};

function handleICEGatheringStateChange(event){
  log("ICEGatheringStateChange: "+pc.iceGatheringState)
};

function handleSignalingStateChange(event){
  log("SignalingStateChange: "+pc.signalingState)
};

// function handleICECandidate(event) {
//   if (event.candidate === null) {
//     document.getElementById('localSessionDescription').value = btoa(JSON.stringify(pc.localDescription))
//   }
// };

function createOffer(){
  return pc.createOffer()
  .then(offer => pc.setLocalDescription(offer))
  .then(() => {
    let offerJSON = JSON.stringify(pc.localDescription));
    document.getElementById('localSessionDescription').value = pc.localDescription;
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
    document.getElementById('remoteSessionDescription').value = json.SDP
    return json.SDP
  })
}

// Start acquiation of media
startMedia()

async function handleNegotiationNeeded(){
  let p = createOffer()
  .then(offerJSON => {
    let msg = {
      name: myUsername,
      target: targetUsername,
      type: "video-offer",
      sdp: pc.localDescription
    }
    return sendToServer("/sdp", msg)
  })
  .then(sdp => {
    pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sdp))))
  })
  // .catch(log)
  // .finally(()=> "3")
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
// function handleTrackEvent(event) {
//   document.getElementById("video1").srcObject = event.streams[0];
// }

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