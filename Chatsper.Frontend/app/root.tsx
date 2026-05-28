import {
  isRouteErrorResponse,
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "react-router";

import type { Route } from "./+types/root";
import "./app.css";

export const links: Route.LinksFunction = () => [
  { rel: "preconnect", href: "https://fonts.googleapis.com" },
  {
    rel: "preconnect",
    href: "https://fonts.gstatic.com",
    crossOrigin: "anonymous",
  },
  {
    rel: "stylesheet",
    href: "https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,100..900;1,14..32,100..900&display=swap",
  },
];

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel={"icon"} href={"https://cdn.jespersen.zip/u/BU1OtH.png"} />
        <Meta />
        <Links />
        <title>Chatsper Platform</title>
      </head>
      <body style={{
        backgroundColor: "#0C0D0C"
      }} className={"h-screen w-screen"}>
      {children}
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return <Outlet />;
}

export function ErrorBoundary({ error }: Route.ErrorBoundaryProps) {
  return (
    <div className="min-h-screen flex justify-center items-center bg-[#0C0D0C]">
      <img
          className={"z-1 relative rounded-full shadow border-20"}
          width={1225}
          height={784}
          alt={"404"}
          src={"https://cdn.jespersen.zip/u/yWVc4p.png"}
      />
    </div>
  );
}
