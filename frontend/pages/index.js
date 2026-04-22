import React, { useState, useEffect } from 'react';
import { Chessboard } from 'react-chessboard';
let Chess;
if (typeof window !== 'undefined') {
  Chess = require('chess.js').Chess;
}

export default function Game() {
  const [game, setGame] = useState(null);
  const [subTurn, setSubTurn] = useState(1);

  useEffect(() => {
    setGame(new Chess());
  }, []);

  function onDrop(source, target) {
    const moveStr = source + target;
    
    console.log("Click detected:", moveStr);

    fetch('http://localhost:8080/move', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ move: moveStr }),
    })
    .then(res => res.json())
    .then(data => {
      console.log("Backend responded:", data);
      if (data.error) {
        console.warn("Engine error:", data.error);
      } else {
        const newGame = new Chess(data.fen);
        setGame(newGame);
        setSubTurn(data.subTurn);
      }
    })
    .catch(err => console.error("Network error:", err));

    return true; 
  }

  if (!game) return <h1 style={{color:'white'}}>Loading...</h1>;

  return (
    <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#1a1a1a', height: '100vh', color: 'white' }}>
      <h1>MarsGo Marseille Chess</h1>
      <div style={{ width: '500px', touchAction: 'none' }}> 
        <Chessboard 
          position={game.fen()} 
          onPieceDrop={onDrop}
          getPositionObject={(pos) => console.log("Current position:", pos)}
        />
      </div>
      <p>SubTurn: {subTurn} | Turn: {game.turn()}</p>
    </div>
  );
}