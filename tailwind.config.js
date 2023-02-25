/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  cssPath: "~/assets/css/tailwind.css",
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./assets/**/*.{jpg,png,svg,gif}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
