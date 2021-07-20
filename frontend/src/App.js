import React from 'react';
import './App.css';
import {Home} from 'gocs-ui-core'
import FilePicker from './components/FilePicker'

function App() {
  return (
    <div className="App">
        <Home 
          getPlayersListApi={window.backend.DemoFile.GetPlayersList} 
          getStatsForPlayerApi={window.backend.DemoFile.GetStatsForPlayerWrapper} 
          getDeathPositionsForPlayerApi={window.backend.DemoFile.GetDeathsPositionForPlayer}
          FilePicker={FilePicker}
        />
    </div>
  );
}

export default App;
