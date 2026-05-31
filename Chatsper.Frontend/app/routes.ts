import {type RouteConfig, index, route, layout} from "@react-router/dev/routes";

export default [
    index("routes/home.tsx"),
    route("/about", "routes/about.tsx"),
    layout("layout/dashboard.tsx", [
        route("/dashboard/overview", "routes/dashboard/overview.tsx")
    ]),
] satisfies RouteConfig;
