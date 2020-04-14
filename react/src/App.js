import React from 'react';

function App() {
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
                                Face
                        </p>
                        <button className="roll-button">Roll</button>
                </div>
        );
}

export default App;
