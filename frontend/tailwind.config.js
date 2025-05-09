/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    safelist: [
        // Menambahkan warna untuk DashboardCard dan elemen lainnya
        {
            pattern: /(bg|text|border|from|to)-(blue|green|yellow|red|gray)-(100|200|300|400|500|600|700|800|900)/,
        },
    ],
    theme: {
        extend: {
            fontFamily: {
                inter: ['Inter', 'sans-serif'],
            },
        },
    },
    plugins: [],
    darkMode: 'class',
};
