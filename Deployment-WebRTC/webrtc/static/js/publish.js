// this code works the modern way
"use strict";

/* eslint-env browser */
var log = msg => {
  document.getElementById('logs').innerHTML += msg + '<br>'
}

let pc = new RTCPeerConnection({
  iceServers: [
    {
      urls: 'stun:stun.l.google.com:19302'
    }
  ]
});

pc.oniceconnectionstatechange = function(e){
  log(pc.iceConnectionState)
};
  
pc.onicecandidate = function(event) {
  if (event.candidate === null) {
    document.getElementById('localSessionDescription').value = btoa(JSON.stringify(pc.localDescription))
  }
};

navigator.mediaDevices.getUserMedia({ video: true, audio: false })
  .then(stream => {
    pc.addStream(document.getElementById('video1').srcObject = stream);
    pc.createOffer()
      .then(d => pc.setLocalDescription(d))
      .catch(log);
  }).catch(log);


function startSession() {

  // let localSDP = document.getElementById('localSessionDescription').value
  // alert("inside startSession")
  let p = fetch("/sdp")
    p.then(response => response.json())
    .then(text => { 
      alert(text.Result);
      alert(text.ServerSDP);
    })
    .then(function(){ 
        alert("done")
      })

  // let sd = document.getElementById('remoteSessionDescription').value
  // if (sd === '') {
  //   return alert('Session Description must not be empty')
  // }

  // try {
  //   pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd))))
  // } catch (e) {
  //   alert(e)
  // }
}