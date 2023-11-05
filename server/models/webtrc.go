package models

// IceCandidate struct
type ICECandidate struct {
	Address          string `json:"address"`
	Candidate        string `json:"candidate"`
	Component        string `json:"component"`
	Foundation       string `json:"foundation"`
	Port             int    `json:"port"`
	Priority         int    `json:"priority"`
	Protocol         string `json:"protocol"`
	RelatedAddress   string `json:"relatedAddress"`
	RelatedPort      int    `json:"relatedPort"`
	SDPMLineIndex    int    `json:"sdpMLineIndex"`
	SDPMid           string `json:"sdpMid"`
	TCPType          string `json:"tcpType"`
	Type             string `json:"type"`
	UsernameFragment string `json:"usernameFragment"`
}

// Media Description struct
type MediaDescription struct {
	Type     string   `json:"type"`
	Port     int      `json:"port"`
	Protocol string   `json:"protocol"`
	Formats  []string `json:"formats"`
}

// SDP struct
type SDP struct {
	Version           string             `json:"version"`
	Origin            string             `json:"origin"`
	SessionName       string             `json:"sessionName"`
	Timing            string             `json:"timing"`
	Connection        string             `json:"connection"`
	MediaDescriptions []MediaDescription `json:"mediaDescriptions"`
}

// Offer struct
type Offer struct {
	Type string `json:"type"`
	Sdp  SDP    `json:"sdp"`
}

// Answer struct
type Answer struct {
	Type string `json:"type"`
	Sdp  SDP    `json:"sdp"`
}
