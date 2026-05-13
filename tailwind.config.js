/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/templates/*.html"],
  theme: {
    extend: {
      typography: (theme) => ({
        DEFAULT: {
          css: {
            a: {
              color: theme("colors.zinc.700"),
            },
            pre: {
              overflowX: "auto",
              maxWidth: '100%',
            },
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography")],
};
