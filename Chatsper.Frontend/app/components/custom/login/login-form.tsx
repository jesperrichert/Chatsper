import {Field, FieldDescription, FieldGroup, FieldLabel, FieldSeparator} from "../../ui/field";
import {cn} from "../../../lib/utils";
import {Card, CardContent} from "../../ui/card";
import {Input} from "../../ui/input";
import {Button} from "../../ui/button";
import {useEffect, useState} from "react";
import type {FrontendConfig} from "../../../types/frontend";
import {fetchConfig} from "../../../lib/config/api";
import { Session } from "../../../lib/auth/session";
import Cookies from 'js-cookie'

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {

  const [config, setConfig] = useState<FrontendConfig>()

  useEffect(() => {
    if (config != null) return
    async function data() {
      const conf = await fetchConfig()
      setConfig(conf)
      
      if (await Session.validate(
        conf.auth.validate_url,
        Cookies.get("session") as string
      )) {
        window.open("/dashboard/overview", "_self")
      } else return
    }
    data()
  }, [config])

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card className="overflow-hidden p-0 rounded-4xl bg-bg-gray">
        <CardContent className="grid p-0 md:grid-cols-2">
          <form className="p-6 md:p-8">
            <FieldGroup>
              <div className="flex flex-col items-center gap-2 text-center">
                <h1 className="text-2xl font-bold text-white">Welcome to Chatsper</h1>
                <p className="text-balance text-gray-500">
                  Login to the Chatsper Platform with your Social Media account.
                </p>
              </div>
              <Field className="grid grid-cols-3 gap-4 ">
                <Button
                    onClick={event => window.open(config?.twitch.auth_url, "_self")}
                    className={"bg-bg-accent rounded-2xl border-2 border-zinc-600 cursor-pointer"} type="button">
                  <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="100" height="100" viewBox="0 0 48 48">
                    <path fill="#7e57c2" d="M42,27.676c-3,3.441-6,6.882-9,10.324c-2.333,0-4.667,0-7,0c-2.333,2-4.667,4-7,6c-1,0-2,0-3,0	c0-2,0-4,0-6c-3.333,0-6.667,0-10,0c0-7.431,0-14.863,0-22.294C7.455,12.804,8.909,9.902,10.364,7C20.909,7,31.455,7,42,7	C42,13.892,42,20.784,42,27.676z"></path><path fill="#fafafa" d="M39,26.369c-1.667,1.877-3.333,3.754-5,5.631c-2.333,0-4.667,0-7,0c-2.333,2-4.667,4-7,6c0-2,0-4,0-6	c-2.667-0.008-5.333-0.016-8-0.024c0-7.326,0-14.651,0-21.976c9,0,18,0,27,0C39,15.456,39,20.912,39,26.369z"></path><rect width="3" height="10" x="21" y="16" fill="#7e57c2"></rect><rect width="3" height="10" x="30" y="16" fill="#7e57c2"></rect>
                  </svg>
                  <span className="sr-only">Login with Twitch</span>
                </Button>
                <Button className={"bg-bg-accent rounded-2xl border-2 border-zinc-600 "} type="button" disabled={true}>
                  <svg xmlns="http://www.w3.org/2000/svg" className="external-icon" viewBox="0 0 28.57  20"
                       focusable="false">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 28.57 20" preserveAspectRatio="xMidYMid meet">
                      <g>
                        <path
                            d="M27.9727 3.12324C27.6435 1.89323 26.6768 0.926623 25.4468 0.597366C23.2197 2.24288e-07 14.285 0 14.285 0C14.285 0 5.35042 2.24288e-07 3.12323 0.597366C1.89323 0.926623 0.926623 1.89323 0.597366 3.12324C2.24288e-07 5.35042 0 10 0 10C0 10 2.24288e-07 14.6496 0.597366 16.8768C0.926623 18.1068 1.89323 19.0734 3.12323 19.4026C5.35042 20 14.285 20 14.285 20C14.285 20 23.2197 20 25.4468 19.4026C26.6768 19.0734 27.6435 18.1068 27.9727 16.8768C28.5701 14.6496 28.5701 10 28.5701 10C28.5701 10 28.5677 5.35042 27.9727 3.12324Z"
                            fill="#FF0000"/>
                        <path d="M11.4253 14.2854L18.8477 10.0004L11.4253 5.71533V14.2854Z" fill="white"/>
                      </g>
                    </svg>
                  </svg>
                  <span className="sr-only">Login with Youtube</span>
                </Button>
                <Button className={"bg-bg-accent rounded-2xl border-2 border-zinc-600"} type="button" disabled={true}>
                  <svg xmlns="http://www.w3.org/2000/svg" width="96" height="96" viewBox="0 0 96 96" fill="none">
                    <rect width="96" height="96" rx="21" fill="black" />
                    <path d="M73.31 25.7456C72.785 25.4743 72.274 25.1769 71.7788 24.8545C70.3389 23.9025 69.0186 22.7808 67.8465 21.5135C64.9139 18.158 63.8186 14.7538 63.4151 12.3705H63.4313C63.0943 10.3921 63.2337 9.11214 63.2547 9.11214H49.8974V60.7624C49.8974 61.4558 49.8974 62.1412 49.8682 62.8185C49.8682 62.9027 49.8601 62.9805 49.8553 63.0712C49.8553 63.1085 49.8553 63.1474 49.8472 63.1863C49.8472 63.196 49.8472 63.2057 49.8472 63.2154C49.7064 65.0686 49.1123 66.8588 48.1173 68.4286C47.1222 69.9983 45.7566 71.2994 44.1407 72.2175C42.4565 73.1757 40.5517 73.6782 38.614 73.6757C32.3906 73.6757 27.3468 68.6011 27.3468 62.334C27.3468 56.0669 32.3906 50.9923 38.614 50.9923C39.7921 50.9912 40.9629 51.1766 42.083 51.5415L42.0992 37.9412C38.6989 37.502 35.2444 37.7722 31.9538 38.7348C28.6631 39.6975 25.6077 41.3317 22.9802 43.5343C20.678 45.5346 18.7425 47.9214 17.2608 50.5872C16.6969 51.5594 14.5695 55.4658 14.3119 61.8058C14.1499 65.4044 15.2306 69.1326 15.7458 70.6734V70.7058C16.0699 71.6132 17.3256 74.7094 19.372 77.3197C21.0221 79.4135 22.9716 81.2527 25.1579 82.7783V82.7459L25.1903 82.7783C31.6567 87.1724 38.8263 86.884 38.8263 86.884C40.0674 86.8338 44.2249 86.884 48.9463 84.6464C54.183 82.1658 57.1642 78.47 57.1642 78.47C59.0688 76.2618 60.5832 73.7452 61.6426 71.0282C62.8513 67.8509 63.2547 64.0401 63.2547 62.5171V35.1155C63.4168 35.2127 65.5749 36.6401 65.5749 36.6401C65.5749 36.6401 68.6842 38.633 73.5352 39.9309C77.0155 40.8544 81.7045 41.0488 81.7045 41.0488V27.7887C80.0615 27.9669 76.7255 27.4485 73.31 25.7456Z" fill="white"/>
                  </svg>
                  <span className="sr-only">Login with Tiktok</span>
                </Button>
              </Field>
              <span className={"text-zinc-400 relative flex-row inline-flex text-xs"}>Running version {config?.version ?? "??.??.??"}</span>
            </FieldGroup>
          </form>
          <div className="p-5 relative hidden md:block">
           <span className={"text-gray-300"}>
             Welcome to the Chatsper Platform.
             <br />
             <i>A small introduction to this service.</i>
            <br /><br />
             Let's start with you own Bot. Chatsper is a Open Source Streaming Bot.
             You can self host this Software or use the Hosted variant on this page.
             When you want to read more about the Features, Self Hosting and everything else then
             <a className={"text-blue-500 hover:text-blue-800"} href={"/about"}> Click Here</a>
           </span>
          </div>
        </CardContent>
      </Card>
      <FieldDescription className="px-6 text-center">
        By clicking continue, you agree to our <a href="">Terms of Service</a>{" "}
        and <a href="#">Privacy Policy</a>.
      </FieldDescription>
    </div>
  )
}
