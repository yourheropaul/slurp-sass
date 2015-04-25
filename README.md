## Sass -> CSS compilation for Slurp, powered by libsass

Sample `slurp.go` code:

    slurp.Task{
        Name:  "scss",
        Usage: "Build scss files into app.css",
        Action: func(c *slurp.C) error {
            return fs.Src(c,
                "./input/scss/*.scss",
            ).Then(
                Compile(c),
                util.Concat(c, "app.css"),
                fs.Dest(c, "./output/css/"),
            )
        },
    },

Note: To conform with the standards of the Ruby Sass/Compass tools, .scss files prefixed with an underscore will not be compiled unless directly @imported.
