import DoughnutChart from "../components/Doughnut"






function Homepage() {
    return (
        <div className="homepage">
            <header className="homepage-header">
                <h1>Bienvenido a SanatorioApp</h1>
                <p>Visualiza el resumen semanal de las citas médicas.</p>
            </header>
            <section className="dashboard">
                <h2>Citas Semanales</h2>
                <DoughnutChart />
            </section>
        </div>
    );
}

export default Homepage
