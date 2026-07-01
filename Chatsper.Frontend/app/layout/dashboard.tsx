import { Outlet } from "react-router";
import type { Route } from "./+types/home";
import React, { useEffect } from "react";
import { DashboardSidebar } from "../components/custom/dashboard/sidebar";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Chatsper Platform" },
        { name: "description", content: "Chatsper Platform" },
    ];
}

export default function Dashboard({ children }: { children: React.ReactNode }) {
    return <div>
        <SidebarProvider>
        <DashboardSidebar />
        <SidebarTrigger size={"lg"} />
        <main className="p-2">
            <Outlet />
        </main>
        </SidebarProvider>
    </div>;
}
