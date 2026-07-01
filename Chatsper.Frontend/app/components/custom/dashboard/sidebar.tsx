import {
  BarChart3,
  ClipboardList,
  HelpCircle,
  LayoutDashboard,
  Settings,
} from "lucide-react";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Separator } from "@/components/ui/separator";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarInset,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
  SidebarRail,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import { cn } from "@/lib/utils";
import { useEffect, useState } from "react";
import type { FrontendConfig } from "@/types/frontend";
import { fetchConfig } from "@/lib/config/api";
import { GearIcon } from "@phosphor-icons/react";
import { useConfig } from "@/context/frontend-config";

const sidebarData = [
    {
      title: "Overview",
      items: [
        {
          label: "Dashboard",
          href: "#",
          isActive: true,
        },
        { label: "Tasks", icon: ClipboardList, href: "#" },
        { label: "Roadmap", icon: BarChart3, href: "#" },
      ],
    }
]
 
export function DashboardSidebar() {
  const config = useConfig()

  return (
    <Sidebar>
      <SidebarHeader className="bg-bg-base text-text-gray">
        <SidebarMenu className="text-text-gray">
      <SidebarMenuItem>
        <SidebarMenuButton size="lg" className="">
          <div className="flex aspect-square size-8 items-center justify-center rounded-s">
            <img
              src={"https://github.com/jesperrichert.png"}
              alt={"404"}
              className="size-6 rounded-2xl"
            />
          </div>
          <div className="flex flex-col gap-0.5 leading-none">
            <span className="font-medium">Chatsper Platform</span>
            <span className="text-xs ext-muted-foreground">
              <GearIcon className="inline-flex"></GearIcon> {" "}
               v.{config?.version}
            </span>
          </div>
        </SidebarMenuButton>
      </SidebarMenuItem>
    </SidebarMenu>
      </SidebarHeader>
      <SidebarContent className="bg-bg-base">
        {sidebarData.map((group) => (
          <SidebarGroup key={group.title}>
            <SidebarGroupLabel className="text-md text-zinc-500"><b>{group.title}</b></SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {group.items.map((item) => (
                  <SidebarMenuItem key={item.label}>
                    <SidebarMenuButton asChild isActive={item.isActive}>
                      <a className="text-text-gray" href={item.href}>{item.label}</a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        ))}
      </SidebarContent>
     {/*
     <SidebarFooter>
          <SidebarGroup>
          <SidebarGroupLabel>{sidebarData.footerGroup.title}</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {sidebarData.footerGroup.items.map((item) => (
                <SidebarMenuItem key={item.label}>
                  <SidebarMenuButton asChild>
                    <a href={item.href}>{item.label}</a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarFooter> 
      */}
      <SidebarRail />
    </Sidebar>
  );
};