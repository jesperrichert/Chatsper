type FrontendConfig = {
    version:string,
    auth: {
        validate_url: string
    },
    twitch: {
        auth_url: string
    }
}

export type {FrontendConfig}