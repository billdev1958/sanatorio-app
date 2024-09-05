import { Doughnut } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import ChartDataLabels from 'chartjs-plugin-datalabels';

ChartJS.register(ArcElement, Tooltip, Legend, ChartDataLabels);

const DoughnutChart = () => {
    const data = {
        labels: ['Cardiología', 'Dermatología', 'Pediatría', 'Ginecología', 'Canceladas'],
        datasets: [
            {
                label: 'Citas Médicas',
                data: [50, 30, 40, 20, 10], // Datos de ejemplo, reemplázalos con los reales
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(255, 206, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                ],
                borderWidth: 1,
            },
        ],
    };

    const options = {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
            legend: {
                position: 'top' as const,
            },
            tooltip: {
                callbacks: {
                    label: function (tooltipItem: any) {
                        return `${tooltipItem.label}: ${tooltipItem.raw} citas`;
                    },
                },
            },
            datalabels: {
                color: '#000',
                anchor: 'end' as const,
                align: 'start' as const, // Coloca la etiqueta hacia el exterior del gráfico
                clamp: true, // Asegura que las etiquetas no se salgan del canvas
                font: {
                    weight: 'bold' as const,
                    size: 14,
                },
                formatter: (value: number) => value,
            },
        },
    };

    return (
        <>
        <div className="chart-container">
            <Doughnut data={data} options={options} />
        </div>
        </>
    );
};

export default DoughnutChart;
