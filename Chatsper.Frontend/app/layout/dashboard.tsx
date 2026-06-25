import type { Route } from "./+types/home";
import React, { useEffect } from "react";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Chatsper Platform" },
        { name: "description", content: "Chatsper Platform" },
    ];
}

export function Layout({ children }: { children: React.ReactNode }) {


    return <div>
        <main>
            {children}
        </main>
    </div>;
}
