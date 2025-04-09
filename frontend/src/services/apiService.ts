import {MainPageData} from "../types/MainPage.ts";
import axios from "axios";

const API_BASE_PATH = '/client';


export const fetchMockUsers = async (): Promise<MainPageData> => {
    try {
        const response = await axios.get(API_BASE_PATH + '/');
        return response.data
    } catch (error: any) {
        console.error(error);
        throw new Error(error.message);
    }
}