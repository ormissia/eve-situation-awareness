import axios from "axios";
import URL_BACKEND from "./Url";

const API_GET_SOLAR_SYSTEM_KILL =()=>{
    return axios.create({
        baseURL:URL_BACKEND.ESA_WEB
    })
}

export default API_GET_SOLAR_SYSTEM_KILL
