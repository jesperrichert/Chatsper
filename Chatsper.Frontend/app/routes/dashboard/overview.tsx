import type { Route } from "./+types/home";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Overview - Chatsper Platform" }, ,
    ];
}

export default function Home() {
    return <div>
        Overview
    </div>;
}
