import React from 'react'
import './PlayersList.css'

const PlayersList=({playersList,selectPlayer,selectedPlayer})=>{


    return (
        <div className="PlayersList-container">
            <div className="PlayersList-header">Select Player</div>
            {   
                playersList.length>0
                ?
                playersList.map((p)=>{
                    const isSelected = selectedPlayer===p
                    const className = "PlayersList-player" + (isSelected?" PlayersList-player-selected":"")
                   return (
                        <div className={className} onClick={selectPlayer(p)}>
                            {p}
                        </div>
                    )
                }
                )
                :
                "Loading players list"
            }
        </div>
    )
}

export default PlayersList