import React, { useEffect,useState } from 'react'
import './HeatmapsPage.css'
import Loader from '../Loader'
import HeatmapOptions from './options' 
import mapNameToFile from '../../utils/mapFileUtils'

const heatmapsEventColorMapping={
    "death": "red",
    "kill": "#00ff00",
    "bomb_plant": "#0000ff"
}

const eventTypes = [
    "kill",
    "death",
    "bomb_plant"
]

const HeatmapsPage=({player,setLoading, mapName})=>{

    const [stats,setStats] = useState(null)
    const [eventType,setEventType] = useState("kill")

    const selectEventType=(type)=>()=>{
        setEventType(type)
    }

    const drawCircle=(ctx,point)=>{
        ctx.beginPath();
        ctx.arc(point.X, point.Y, 5,0, 2 * Math.PI);
        ctx.stroke();  
        ctx.fill()
    }

    const clearPlot=()=>{
        const c = document.getElementById("myCanvas");
        const ctx = c.getContext("2d");
        ctx.clearRect(0, 0, ctx.canvas.width, ctx.canvas.height);
        ctx.beginPath();
    }

	useEffect(()=>{
        const parseDemo = async(player)=>{
            setLoading(true)
            setStats(null)
            clearPlot()
            console.log('fetch')
            const resp = await window.backend.DemoFile.GetDeathsPositionForPlayer(player)
            console.log('complete',resp)
            setStats(resp)
            setLoading(false)
        }
		if (player.length>0){
			parseDemo(player)
		}
	},[player, setLoading])

    useEffect(()=>{
        if(!stats){
            return
        }
        console.log("redraw")
        clearPlot()
        const c = document.getElementById("myCanvas");
        const ctx = c.getContext("2d");
        ctx.fillStyle = heatmapsEventColorMapping[eventType];

        stats[eventType].forEach((pt)=>{
            drawCircle(ctx,pt)
        })
    },[eventType, stats])

    const mapFile = mapNameToFile[mapName]
    return (
        <div className="heatmaps-container">
            <div className="heatmaps-map-container">
                <div className="heatmaps-map-bg" style={{backgroundImage:`url(${mapFile})`,backgroundSize:"contain"}}>
                    <canvas 
                    id="myCanvas" 
                    height={512} 
                    width={512} 
                    />
                    {
                        !stats && <Loader/>
                    }
                </div>
            </div>
            <div className={"heatmaps-options-container"}>
                <HeatmapOptions eventType={eventType} selectEventType={selectEventType} options={eventTypes} colorMap={heatmapsEventColorMapping}/>
            </div>
        </div>
    )
}

export default HeatmapsPage