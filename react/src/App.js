import React, { useEffect, useState } from 'react';

function App() {
        let [name, setName] = useState('');
        let [face, setFace] = useState('');

        useEffect(() => {
                refDice().then((dice) => {
                        setName(dice.name);
                        setFace(dice.faces[0]);
                });
        }, []);

        function refDice() {
                return fetch('/v1/dice').then((res) => {
                        return res.json();
                });
        }

        function rollDice() {
                fetch('/v1/dice', { method: 'POST' }).then((res) => {
                        if (res.status === 204) {
                                refDice().then(({ faces }) => {
                                        let i = faces.length;
                                        let roller = setInterval(() => {
                                                i--;
                                                setFace(faces[i]);
                                                if (i <= 0) {
                                                        clearInterval(roller);
                                                }
                                        }, 50);
                                });
                        }
                });
        }

        return (
                <div className="app">
                        <style jsx>{`
                                .app {
                                        text-align: center;
                                        font-family: serif;
                                }

                                h1 {
                                        text-align: center;
                                        font-size: 34pt;
                                }

                                .face {
                                        width: 8em;
                                        line-height: 8em;
                                        margin: 3ex auto;
                                        border: solid 1pt #000;
                                        font-size: 21pt;
                                }

                                .roll-button {
                                        background: #039;
                                        color: #fff;
                                        border: 0;
                                        width: 130pt;
                                        font-size: 21pt;
                                }

                                .roll-button:hover {
                                        background: #69f;
                                }
                        `}</style>
                        <h1>{name}</h1>
                        <p className="face">
                                {face}
                        </p>
                        <button className="roll-button" onClick={rollDice}>Roll</button>
                </div>
        );
}

export default App;
