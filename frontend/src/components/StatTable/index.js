import React from 'react';
import './StatTable.css'

const generateStatBar=(item,comparisonFields)=>{
    const val1 = item[comparisonFields[0]]
    const val2 = item[comparisonFields[1]]
    if (val1===0 && val2===0){
        return <div style={{backgroundColor:"#888888", width:'100%'}}/>
    }
    else if (val1===0){
        return <div style={{backgroundColor:"#bd4261",width:'100%'}}/>
    }
    else if (val2===0){
        return <div style={{backgroundColor:"#42bd77",width:'100%'}}/>
    }
    const fraction = Math.round(98 * val1/(val1+val2))
    const bar1 = <div style={{backgroundColor:"#42bd77",width:`${fraction}%`}}/>
    const separator = <div style={{width:`2%`}}/>
    const bar2 = <div style={{backgroundColor:"#bd4261",width:`${98-fraction}%`}}/>
    return (
        <>
        {bar1}{separator}{bar2}
        </>
    )
}

const StatTable =({data,fields,comparisonFields}) => {

	return (
		<div className='StatTable-container'>
			<table className='StatTable-table'>
                <thead>
                    <tr className='StatTable-table-headerrow'>
                        <td className='StatTable-table-name' colSpan={fields.length}>{data.category}</td>
                    </tr>
                    <tr className='StatTable-table-headerrow'>
                        {
                            fields.map((field)=>(
                                <td className='StatTable-table-cell StatTable-table-header-cell'>{field}</td>
                            ))
                        }
                    </tr>
                </thead>
                <tbody>
                {
                    data.items.map((item)=>
                        (<>
                        <tr className='StatTable-table-row'>
                            {fields.map((field)=>(<td className='StatTable-table-cell'>{item[field]}</td>))}
                        </tr>
                        {comparisonFields &&
                            <tr className='StatTable-table-bar-row'>
                                <td> 
                                <div className="StatTable-bar-container">
                                    {generateStatBar(item,comparisonFields)}
                                </div>
                                </td>  
                            </tr>
                        }
                        <tr className='StatTable-separator'/>
                        </>)
                    )
                }
                </tbody>
            </table>
		</div>
	);
}

export default StatTable;
