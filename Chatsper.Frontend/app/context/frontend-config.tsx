import { fetchConfig } from "@/lib/config/api";
import type { FrontendConfig } from "@/types/frontend";
import { Edit } from "lucide-react";
import React, { useEffect, useState } from "react";

const FrontendConfigContext = React.createContext<FrontendConfig|null>(null)

export function useConfig() {
    return React.useContext(FrontendConfigContext)
}

export function FrontendConfigProvider({children}: { children: React.ReactNode }) {
    const [configData, setConfigData] = useState<FrontendConfig>()

  useEffect(() => {
    if (configData != null) return
    async function data() {
      const conf = await fetchConfig()
      setConfigData(conf)
    }
    data()
  }, [configData])

  return <FrontendConfigContext.Provider 
  value={configData as FrontendConfig}>
    {children}
  </FrontendConfigContext.Provider>
}