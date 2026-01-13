/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "brand-orange": "#CF7E13",
        "brand-blue": "#227082",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
