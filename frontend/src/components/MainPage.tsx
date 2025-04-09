import {useEffect, useState} from "react";
import {MainPageData} from "../types/MainPage.ts";
import {fetchMockUsers} from "../services/apiService.ts";

const MainPage: React.FC = () => {
    const [data, setData] = useState<MainPageData | null>(null)

    useEffect(() => {
        const loadData = async() => {
            await fetchMockUsers().then(result => setData(result))
        }
        loadData()
    }, []);

    
    return (
        <div>
            <h1>Users</h1>
            <ul>
                {data?.users.map((user, index) => (
                    <li key={index}>{user.name}</li>
                ))}
            </ul>
        </div>
    )
}

export default MainPage