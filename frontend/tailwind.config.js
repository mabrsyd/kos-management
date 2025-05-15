/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    safelist: [
        // Menambahkan warna untuk DashboardCard dan elemen lainnya
        {
            pattern: /(bg|text|border|from|to)-(blue|green|yellow|red|gray|purple|indigo|pink)-(100|200|300|400|500|600|700|800|900)/,
        },
    ],
    theme: {
        screens: {
            'xs': '480px',
            'sm': '640px',
            'md': '768px',
            'lg': '1024px',
            'xl': '1280px',
            '2xl': '1536px',
        },
        extend: {
            fontFamily: {
                inter: ['Inter', 'sans-serif'],
            },
            spacing: {
                '72': '18rem',
                '80': '20rem',
                '96': '24rem',
            },
            maxWidth: {
                '8xl': '88rem',
                '9xl': '96rem',
            },
            backgroundColor: {
                'black-opacity-50': 'rgba(0, 0, 0, 0.5)',
            },
            minHeight: {
                'screen-75': '75vh',
            },
            transitionProperty: {
                'height': 'height',
                'spacing': 'margin, padding',
            },
            zIndex: {
                '60': '60',
                '70': '70',
                '80': '80',
                '90': '90',
                '100': '100',
            },
        },
    },
    plugins: [],
    darkMode: 'class',
};
