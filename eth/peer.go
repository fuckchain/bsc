// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"math/big"
	"net"

	"github.com/ethereum/go-ethereum/eth/protocols/bsc"
	"github.com/ethereum/go-ethereum/eth/protocols/trust"

	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
)

// ethPeerInfo represents a short summary of the `eth` sub-protocol metadata known
// about a connected peer.
type ethPeerInfo struct {
	Version     uint     `json:"version"`     // Ethereum protocol version negotiated
	Difficulty  *big.Int `json:"difficulty"`  // Total difficulty of the peer's blockchain
	Head        string   `json:"head"`        // Hex hash of the peer's best owned block
	BlockNumber uint64   `json:"blockNumber"` // Block number of the peer's best owned block
	BoundAt     int64    `json:"boundAt"`     // Time when the peer was bound
	TxSum       uint64   `json:"txSum"`       // Total number of transactions
	SendTxSum   uint64   `json:"sendTxSum"`   // Total number of transactions sent
}

// ethPeer is a wrapper around eth.Peer to maintain a few extra metadata.
type ethPeer struct {
	*eth.Peer
	snapExt  *snapPeer // Satellite `snap` connection
	trustExt *trustPeer
	bscExt   *bscPeer // Satellite `bsc` connection
}

// info gathers and returns some `eth` protocol metadata known about a peer.
func (p *ethPeer) info() *ethPeerInfo {
	hash, td := p.Head()
	blockNumber, txSum, sendTxSum, boundAt := p.Stat()
	return &ethPeerInfo{
		Version:     p.Version(),
		Difficulty:  td,
		Head:        hash.String(),
		BlockNumber: blockNumber,
		TxSum:       txSum,
		SendTxSum:   sendTxSum,
		BoundAt:     boundAt.UnixNano() / 1e6,
	}
}

func (p *ethPeer) remoteAddr() net.Addr {
	if p.Peer != nil && p.Peer.Peer != nil {
		return p.Peer.Peer.RemoteAddr()
	}
	return nil
}

// snapPeerInfo represents a short summary of the `snap` sub-protocol metadata known
// about a connected peer.
type snapPeerInfo struct {
	Version uint `json:"version"` // Snapshot protocol version negotiated
}

// trustPeerInfo represents a short summary of the `trust` sub-protocol metadata known
// about a connected peer.
type trustPeerInfo struct {
	Version uint `json:"version"` // Trust protocol version negotiated
}

// bscPeerInfo represents a short summary of the `bsc` sub-protocol metadata known
// about a connected peer.
type bscPeerInfo struct {
	Version uint `json:"version"` // bsc protocol version negotiated
}

// snapPeer is a wrapper around snap.Peer to maintain a few extra metadata.
type snapPeer struct {
	*snap.Peer
}

// trustPeer is a wrapper around trust.Peer to maintain a few extra metadata.
type trustPeer struct {
	*trust.Peer
}

// bscPeer is a wrapper around bsc.Peer to maintain a few extra metadata.
type bscPeer struct {
	*bsc.Peer
}

// info gathers and returns some `snap` protocol metadata known about a peer.
func (p *snapPeer) info() *snapPeerInfo {
	return &snapPeerInfo{
		Version: p.Version(),
	}
}

// info gathers and returns some `trust` protocol metadata known about a peer.
func (p *trustPeer) info() *trustPeerInfo {
	return &trustPeerInfo{
		Version: p.Version(),
	}
}

// info gathers and returns some `bsc` protocol metadata known about a peer.
func (p *bscPeer) info() *bscPeerInfo {
	return &bscPeerInfo{
		Version: p.Version(),
	}
}
