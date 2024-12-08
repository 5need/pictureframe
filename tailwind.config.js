/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        mono: ["Ubuntu Mono", "monospace"],
      },
      keyframes: {
        diceRoll: {
          "0%": {
            transform: "scale(100%) rotate(0deg)",
          },
          "20%": {
            transform: "scale(100%) rotate(30deg)",
          },
          "40%": {
            transform: "scale(100%) rotate(-30deg)",
          },
          "60%": {
            transform: "scale(100%) rotate(30deg)",
          },
          "100%": {
            transform: "scale(100%) rotate(0deg)",
          },
        },
        diceRollOld: {
          "0%": {
            transform:
              "perspective(400px) rotateY(0deg) translate(0,0) scale(100%) rotate(0deg)",
            transformOrigin: "center left",
          },
          "40%": {
            transform:
              "perspective(400px) rotateY(-30deg) translate(0,0) scale(120%) rotate(0deg)",
            transformOrigin: "center left",
          },
          "80%": {
            transform:
              "perspective(400px) rotateY(-30deg) translate(0,0) scale(100%) rotate(0deg)",
            transformOrigin: "center left",
          },
          "100%": {
            transform:
              "perspective(400px) rotateY(0deg) translate(0%,0) scale(100%) rotate(0deg)",
            transformOrigin: "center left",
          },
        },
        diceRollOldOld: {
          "0%": {
            transform:
              "perspective(400px) rotateY(0deg) translate(0,0) scale(100%) rotate(0deg)",
          },
          "20%": {
            transform:
              "perspective(400px) rotateY(-35deg) translate(-35%,0) scale(180%) rotate(0deg)",
          },
          "40%": {
            transform:
              "perspective(400px) rotateY(-35deg) translate(-15%,0) scale(100%) rotate(0deg)",
          },
          "60%": {
            transform:
              "perspective(400px) rotateY(0deg) translate(0%,0) scale(100%) rotate(0deg)",
          },
          "80%": {
            transform:
              "perspective(400px) rotateY(0deg) translate(0%,0) scale(100%) rotate(0deg)",
          },
          "100%": {
            transform:
              "perspective(400px) rotateY(0deg) translate(0%,0) scale(100%) rotate(0deg)",
          },
        },

        diceShine: {
          "0%": { transform: "translate(-100%, 0)", opacity: "0" },
          "50%": { transform: "translate(0%, 0)", opacity: "1" },
          "100%": { transform: "translate(100%, 0)", opacity: "0" },
        },
      },
      animation: {
        // diceRoll: "diceRoll 0.8s cubic-bezier(0.1, 0.7, 1, 0.1)",
        diceRoll: "diceRoll 0.2s linear",
        diceShine: "diceShine 0.4s 1s linear forwards",
      },
    },
  },
  plugins: [
    require("@catppuccin/tailwindcss")({
      prefix: "ctp",
      defaultFlavour: "mocha",
    }),
    require("@tailwindcss/typography"),
    require("@tailwindcss/forms"),
    require("@tailwindcss/aspect-ratio"),
    require("@tailwindcss/container-queries"),
  ],
};
