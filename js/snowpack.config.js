/** @type {import("snowpack").SnowpackUserConfig } */
module.exports = {
  mount: {
	  pub: '/',
	  src: '/dist',
  },
  plugins: [
    '@snowpack/plugin-svelte',
  ],
  install: [
    /* ... */
  ],
  installOptions: {
    rollup: {
      plugins: [
        require("rollup-plugin-svelte")({
          include: ["./node_modules"],
        }),
        require("rollup-plugin-postcss")({
          use: [
            [
              "sass",
              {
                includePaths: [".", "./node_modules"],
              },
            ],
          ],
        }),
      ],
    },
  },
  devOptions: {
    port: 3000,
    output: "stream",
    open: "none"
  },
  buildOptions: {
    out: "../bin/js"
  },
  proxy: {
    /* ... */
  },
  alias: {
    /* ... */
  },
};
