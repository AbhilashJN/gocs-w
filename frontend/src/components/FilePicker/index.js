import React from 'react';
import './FilePicker.css'

const FilePicker=({setFile})=>{

const selectFile=async ()=>{
    const mapName = await window.backend.DemoFile.SelFile()
    console.log(mapName,'***')
    setFile(mapName)
}

return (
    <div className="FilePicker-container">
        <div onClick={selectFile} className="FilePicker-picker">
            <span className="FilePicker-header">No demo file selected.</span>            
            Select a file to analyze.
        </div>
    </div>
)

}

export default FilePicker