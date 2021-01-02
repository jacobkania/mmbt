/** @type {import("snowpack").SnowpackUserConfig } */

const rollupSvelte = require("rollup-plugin-svelte")({
  include: ["./node_modules"],
});

const rollupSass = require("rollup-plugin-postcss")({
  use: [
    [
      "sass",
      {
        includePaths: ["./node_modules", "."],
      },
    ],
  ],
});

module.exports = {
  mount: {
    pub: "/",
    src: "/dist",
  },
  plugins: ["@snowpack/plugin-svelte"],
  install: [
    /* ... */
  ],
  installOptions: {
    rollup: {
      plugins: [rollupSvelte, rollupSass],
    },
  },
  devOptions: {
    port: 3000,
    output: "stream",
    open: "none",
  },
  buildOptions: {
    out: "../bin/js",
  },
  proxy: {
    /* ... */
  },
  alias: {
    components: "./src/components",
    pages: "./src/pages",
    utils: "./src/utils",
    app: "./src",
  },
};
