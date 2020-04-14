import React, { useEffect, useState } from 'react';

function App() {
        let [face, setFace] = useState('');

        useEffect(() => {
                fetch('/v1/dice')
                        .then((res) => {
                                return res.json();
                        })
                        .then((faces) => {
                                setFace(faces[0]);
                        });
        }, []);

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
                                        margin: 3ex 0;
                                        font-size: 34pt;
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
                        <h1>Reactive Dice</h1>
                        <p className="face">
                                {face}
                        </p>
                        <button className="roll-button">Roll</button>
                </div>
        );
}

export default App;
