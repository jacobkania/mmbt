import postcss from "rollup-plugin-postcss";

module.exports = {
    rollupWrapper: rollup => {
    rollup.plugins = [
      ...rollup.plugins, 
      postcss({
        // extract: true,
        minimize: true,
        use: [
          ['sass', {
            includePaths: [
              './src/theme',
              './node_modules'
            ]
          }]
        ]
      })
    ]
  }
}