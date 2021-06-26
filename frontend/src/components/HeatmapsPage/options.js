import React from 'react'
import './options.css'




const HeatmapOptions =({eventType,options,selectEventType,colorMap})=>{

    
return (
    <div className="heatmap-options">
        {
            options.map((opt)=>{
            const isSelected = opt===eventType
            const style= isSelected?"heatmap-options-opt heatmap-options-opt-selected":"heatmap-options-opt"
            return (
                <div className={style} onClick={selectEventType(opt)}>
                    {opt}
                    <div className={"heatmap-options-opt-index"} style={{backgroundColor:colorMap[opt]}}/>
                </div>
                )
            })
        }
    </div>
)
}

export default HeatmapOptions