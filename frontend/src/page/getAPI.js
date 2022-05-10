import axios from "axios";
import React, { useEffect, useState } from "react";

function GetAPI() {
  const [msg, setMsg] = useState("");
  const [object, setObject] = useState("")

  useEffect(() => {
    testgetAPI();
  }, []);
  const testgetAPI = async() => {
    await axios.get("/health").then((res) => {
      setMsg(res.data);
    });
    await axios.get("/test").then((res) =>{
      setObject(res.data)
    })
  };

  return(
    <div>
      <div>Get Message for API : {msg}</div>
      <div>Get Person for Database : {object}</div>
    </div>
  ) 
}
export default GetAPI;
