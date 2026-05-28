import type {FrontendConfig} from "../../types/frontend";

const fetchConfig = async () : Promise<FrontendConfig> => {
    const data = await fetch("/api/frontend")
    return await data.json()
}

export {
    fetchConfig
}