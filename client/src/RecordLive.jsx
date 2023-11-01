import React, { useState, useEffect } from "react";

const RecordLive = () => {
  const [localSessionDescription, setLocalSessionDescription] = useState("");
  const [remoteSessionDescription, setRemoteSessionDescription] = useState("");
  const [logs, setLogs] = useState([]);
  const [videoStream, setVideoStream] = useState(null);
  const [pc, setPC] = useState(null);

  const log = (msg) => {
    setLogs((prevLogs) => [...prevLogs, msg]);
  };

  const copySDP = () => {
    if (!localSessionDescription) {
      log("Session Description must not be empty");
      return;
    }

    try {
      const browserSDP = document.getElementById("localSessionDescription");
      browserSDP.focus();
      browserSDP.select();
      const successful = document.execCommand("copy");
      const msg = successful ? "successful" : "unsuccessful";
      log("Copying SDP was " + msg);
    } catch (err) {
      log("Unable to copy SDP " + err);
    }
  };

  const startSession = () => {
    if (!remoteSessionDescription) {
      log("Session Description must not be empty");
      return;
    }

    if (pc) {
      try {
        log("Starting the session");
        console.log("pc");
        console.log(pc);
        // pc.setRemoteDescription(
        //   new RTCSessionDescription({
        //     type: "offer",
        //     sdp: atob(remoteSessionDescription),
        //   })
        // );
        pc.setRemoteDescription(JSON.parse(atob(remoteSessionDescription)));
        log("Starting the session");
      } catch (e) {
        alert(e);
      }
    } else {
      log("Peer connection not initialized. Please try again.");
    }
  };

  const initPeerConnection = async () => {
    const peerConnection = new RTCPeerConnection({
      iceServers: [
        {
          urls: "stun:stun.l.google.com:19302",
        },
      ],
    });

    peerConnection.oniceconnectionstatechange = (e) =>
      log(peerConnection.iceConnectionState);
    peerConnection.onicecandidate = (event) => {
      if (event.candidate === null) {
        setLocalSessionDescription(
          btoa(JSON.stringify(peerConnection.localDescription))
        );
      }
    };

    try {
      const stream = await navigator.mediaDevices.getUserMedia({
        video: true,
        audio: true,
      });
      setVideoStream(stream);
      console.log("stream", stream);
      document.getElementById("video1").srcObject = stream;
      stream
        .getTracks()
        .forEach((track) => peerConnection.addTrack(track, stream));

      const offer = await peerConnection.createOffer();
      await peerConnection.setLocalDescription(offer);

      setPC(() => peerConnection); // Set the peer connection in the component state.
    } catch (err) {
      log(err);
    }
  };

  useEffect(() => {
    initPeerConnection();
  }, []);

  return (
    <div>
      <h1>RecordLive</h1>
      <div>
        <textarea
          id="localSessionDescription"
          readOnly
          value={localSessionDescription}
          style={{ width: "500px", minHeight: "75px" }}
        ></textarea>
        <button onClick={copySDP}>Copy browser SDP to clipboard</button>
      </div>
      <div>
        <textarea
          id="remoteSessionDescription"
          value={remoteSessionDescription}
          onChange={(e) => setRemoteSessionDescription(() => e.target.value)}
          style={{ width: "500px", minHeight: "75px" }}
        ></textarea>
        <button onClick={startSession}>Start Session</button>
      </div>
      <div>
        <video
          id="video1"
          width="160"
          height="120"
          autoPlay
          muted
          srcObject={videoStream}
        ></video>
      </div>
      <div>
        <h2>Logs</h2>
        <div id="logs">
          {logs.map((log, index) => (
            <p key={index}>{log}</p>
          ))}
        </div>
      </div>
    </div>
  );
};

export default RecordLive;
