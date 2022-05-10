import React, { Component } from 'react'
import axios from "axios";
import { useEffect, useState } from "react";

function Card() {
    const [object, setObject] = useState([])

    useEffect(() => {
        getAPI();
      }, []);
      const getAPI = async() => {
        await axios.get("/data/all").then((res) =>{
          setObject(res.data)
        })
      };
  return (
    <div>
        <h2>คณะที่เปิดรับสมัครอยู่</h2>
        {object.map((item, index) =>(
            <div>
                <h2 key={index}>{item.major_name}</h2>
                <p key={index}>{item.major}</p>
                <p key={index}>{item.major_des}</p>
                <p key={index}>{item.open_date}</p>
                <button>สมัคร</button>
                <button>ประกาศสมัคร</button>
                <button>ข้อมูลเพิ่มเติม</button>
            </div>
                    
                ))}
    </div>
  )
}
export default Card
