const path = require("path");

/** @type {import('@remix-run/dev').AppConfig} */
module.exports = {
  cacheDirectory: "./node_modules/.cache/remix",
  future: {
    v2_errorBoundary: true,
    v2_meta: true,
    v2_normalizeFormMethod: true,
    v2_routeConvention: true,
    v2_headers: true,
  },
  ignoredRouteFiles: ["**/.*", "**/*.test.{js,jsx,ts,tsx}"],
  publicPath: "/_static/build/",
  postcss: true,
  server: "./server.ts",
  serverBuildPath: "server/index.js",
  serverModuleFormat: "cjs",
  tailwind: true,
  routes(defineRoutes) {
    return defineRoutes((route) => {
      if (process.env.NODE_ENV === "production") return;

      return;
    });
  },
};
