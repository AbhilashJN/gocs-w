import React, { useEffect, useState } from 'react'
import StatTable from '../StatTable'
import './StatsPage.css'
import Loader from '../Loader'

const statTypes=["Damage","Deaths/Kills","Accuracy"]


const StatsPage = ({player,setLoading,loading})=>{

    const [statType,setStatType] = useState(statTypes[0])
    const [stats,setStats] = useState({})

    const selectStatType=(type)=>()=>{
        if(loading){
            return
        }
        setStatType(type)
    }

	useEffect(()=>{
        const parseDemo = async(player)=>{
            setLoading(true)
            setStats({})
            console.log('fetch')
            const resp = await window.backend.DemoFile.GetStatsForPlayerWrapper(player,statType);
            console.log('complete',resp)
            setStats(resp)
            setLoading(false)
        }

		if (player.length>0){
			parseDemo(player)
		}
	},[player,statType, setLoading])

    return (
        <div className='StatsPage-container'>
        <div className='StatsPage-navbar'>
            {
                statTypes.map((type)=>{
                    const isSelected = type === statType
                    const style = isSelected?`StatsPage-navbar-btn StatsPage-navbar-btn-selected`:'StatsPage-navbar-btn'
                    return (<div className={style} onClick={selectStatType(type)}>{type}</div>)
                })
            }
        </div>
        <div className='StatsPage-body'>
        {
            stats.data?
            stats.data.map((stat)=>{
                return <StatTable data={stat} fields={stats.item_fields} comparisonFields={stats.comparison_fields}/>
            })
            :
            <Loader/>
        }
        </div>
        </div>
    )
}

export default StatsPage