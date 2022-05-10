import React, { Component } from 'react'
import Select from 'react-select'
import axios from "axios";
import { useEffect, useState } from "react";

function SelectFaculty() {
    const [object, setObject] = useState([])

    useEffect(() => {
        getAPI();
      }, []);
      const getAPI = async() => {
        await axios.get("/faculty/all").then((res) =>{
          setObject(res.data)
          console.log(res.data)
        })
      };

    return(
        <div>
            <h2>เลือกคณะ</h2>
            <select>
                {object.map((item, index) =>(
                    <option key={index}>{item.faculty_name}</option>
                ))}
            </select>
            <h2>เลือกรอบที่ต้องการสมัคร</h2>
            <select>
                {object.map((item, index) =>(
                    <option key={index}>{item.faculty_name}</option>
                ))}
            </select>
            
            
        </div>
       
    )
}

export default SelectFaculty