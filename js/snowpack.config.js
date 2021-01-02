/** @type {import("snowpack").SnowpackUserConfig } */

const rollupSvelte = require("rollup-plugin-svelte")({
  include: ["./node_modules"],
});

const rollupSass = require("rollup-plugin-postcss")({
  use: [
    [
      "sass",
      {
        // These are options for the sass package, from https://www.npmjs.com/package/sass
        includePaths: ["./node_modules", "."],
        data: `
          // Copied files manually into the pub/fonts/material-icons folder from node_modules/material-icons/iconfont
          $material-icons-font-path: 'fonts/material-icons/';
          @import '~material-icons/iconfont/material-icons.scss';
        `,
      },
    ],
  ],
  minimize: true,
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
    materialIcons: "material-icons/iconfont",
  },
};
