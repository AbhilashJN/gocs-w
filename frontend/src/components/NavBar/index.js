import React from 'react'
import "./NavBar.css"

const NavBar = ({selectedPage,selectPage,pages})=>{

    return (
        <div className="navbar-container">
            {
                pages.map((page)=>{
                    const isSelected = selectedPage === page
                    return (
                        <div 
                            className={`navbar-btn${isSelected?" navbar-btn-selected":""}`}
                            onClick={selectPage(page)}
                        >
                            {page}
                        </div>
                    )
                })
            }
        </div>
    )
}

export default NavBar