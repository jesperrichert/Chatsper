import type { Route } from "./+types/home";
import {Card, CardContent, CardFooter} from "../components/ui/card";
import {ArrowLeft} from "lucide-react";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "About Chatsper" },
        { name: "description", content: "What is Chatsper?" },
    ];
}

export default function Home() {
    return  <div className="flex min-h-svh flex-col items-center justify-center p-6 md:p-10">
        <div className="w-full max-w-sm md:max-w-4xl">
            <Card className="overflow-hidden p-0 rounded-4xl bg-bg-gray">
                <CardContent className="p-6 text-gray-200 text-sm inline-flex flex-row">
                    <div className={"w-96"}>
                        <h1>Introductions</h1>
                        <hr className={"p-2"} />
                        <span> Chatsper? Yeah right. Chatsper is an open-source, free to use Twitch Bot...</span>
                        <br/>
                        <span>
                        You can use it here at <i className={"text-zinc-400"}><code>chataper.jespersen.zip</code></i> or setup it on your own VPS/PC.<br/>
                        When you need help please contact us at <a className={"text-blue-500"} target={"_blank"}
                                                                   href={"https://discord.gg/ZXeQCpUYWZ"}>@Discord</a> or create an issue on <a
                            className={"text-blue-500"} target={"_blank"}
                            href={"https://github.com/jesperrichert/Chatsper/issues"}>@GitHub</a>.
                    </span>
                    </div>
                    <span className="border-t sm:border-t-0 sm:border-s border-stone-200 dark:border-neutral-700 mx-5"></span>
                   <div className={"w-96"}>
                       <h1>What & Why Chatsper?</h1>
                       <hr className={"p-2"} />
                        <span>
                            So the question is why Chatsper and what is it? So Chatsper is an Bot for Streamer. It can be used as Cloud Platform online like other Bot on Twitch, etc. or you can control it on your PC or VPS. 100% your Control! Data ist only stored locally...
                            <br />
                            <a target={"_blank"} href={"https://doc.jespersen.zip/s/chatspers"} className={"text-blue-500"}>Read more in the Docs.</a>
                        </span>
                   </div>
                </CardContent>
                <CardFooter className={"flex items-center min-h-min justify-center"}>
                    <button onClick={() => window.open("/")} className={"text-white text-2xl hover:scale-50 cursor-pointer"}><ArrowLeft className={"inline-flex"} /> Go Back</button>
                </CardFooter>
            </Card>
        </div>
    </div>;
}
