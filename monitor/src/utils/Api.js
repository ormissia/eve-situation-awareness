import axios from "axios";
import {URL_BACKEND} from "./Url";

export function API_GET_SOLAR_SYSTEM_KILL() {
    return axios.create({
        baseURL:URL_BACKEND.ESA_WEB+"/analysis"
    })
}

export  function API_GET_SOLAR_SYSTEM_INFO_FUZZY() {
    return axios.create({
        baseURL:URL_BACKEND.ESA_WEB+"/basic"
    })
}
