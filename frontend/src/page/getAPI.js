import axios from "axios";
import React, { useEffect, useState } from "react";

function GetAPI() {
  const [msg, setMsg] = useState("");

  useEffect(() => {
    testgetAPI();
  }, []);
  const testgetAPI = async() => {
    await axios.get("/health").then((res) => {
      setMsg(res.data);
    });
  };

  return <div>Get Message for API {msg}</div>;
}
export default GetAPI;
