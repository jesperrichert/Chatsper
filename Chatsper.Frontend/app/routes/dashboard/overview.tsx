import { useEffect, useState } from "react";
import type { Route } from "./+types/home";
import { fetchConfig } from "../../lib/config/api";
import type { FrontendConfig } from "../../types/frontend";
import { Session } from "../../lib/auth/session";
import Cookies from 'js-cookie'

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Overview - Chatsper Platform" }, ,
    ];
}

export default function Home() {
   const [loggedIn, setLoggedIn] = useState<boolean>(false)

  useEffect(() => {
    if (loggedIn) return
    async function data() {
      const conf = await fetchConfig()
    
      if (await Session.validate(
        conf.auth.validate_url,
        Cookies.get("session") as string
      )) {
        setLoggedIn(true)
      } else window.open("/", "_self")
    }
    data()
  })

    return <div>
        LOGGED-IN
    </div>;
}
