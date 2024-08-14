/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./components/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        basier: ["Basier", "sans-serif"],
      },
      colors: {
        primary: "#085CA7",
        "primary-content": "#FFFFFFD8",
        secondary: "#BFAE71",
        "secondary-content": "#0d0b04",
        accent: "#DC8850",
        "accent-content": "#110602",
        neutral: "#1C1C1C",
        "neutral-content": "#757575",
        "base-100": "#161616",
        "base-200": "#282828",
        "base-300": "#000000",
        "base-content": "#bebebe",
        info: "#2563EB",
        "info-content": "#d2e2ff",
        success: "#52B45C",
        "success-content": "#020c03",
        warning: "#FFCA11",
        "warning-content": "#160f00",
        error: "#F35248",
        "error-content": "#140202",
      },
    },
  },
  plugins: [require("daisyui")],
};
