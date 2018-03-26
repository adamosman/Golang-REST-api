import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import ReactFileUpload from './ReactFileUpload';

class App extends Component {
  state = {
    wordCount: null,
    wordCountPairs: []
  }

  onSuccess = data => {
    console.log('GOT THE DATA!!', data)
    this.setState({
      wordCount: data.wordCount,
      wordCountPairs: data.wordCountPairs
    })
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
        <ReactFileUpload onSuccess={this.onSuccess}/>

        {this.state.wordCount && this.state.wordCountPairs.length ? (
          <div>
            <h1>There are {this.state.wordCount} words</h1>
            <table>
              <thead>
                <tr>
                  <th>Word</th>
                  <th>Count</th>
                </tr>
              </thead>
              <tbody>
                {this.state.wordCountPairs.map(pair => {
                  return (
                    <tr>
                      <td>{pair.word}</td>
                      <td>{pair.count}</td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        ) : null}
      </div>
    );
  }
}

export default App;
