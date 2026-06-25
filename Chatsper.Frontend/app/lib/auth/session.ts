import type { SessionResponse } from "../../types/api/auth";

class Session {

    public static async validate(path:string, sessionId:string): Promise<boolean> {
 
        console.log("sdfsd");
        
 
        const data = await fetch(path, {
            method: "POST",
            headers: {
                "Authorization": sessionId
            }
        })
        const json = await data.json() as SessionResponse
        return json.authenticated
    }

}

export {
    Session
}